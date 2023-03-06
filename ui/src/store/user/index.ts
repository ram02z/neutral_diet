import { atom, atomFamily, selector, selectorFamily } from 'recoil';

import dayjs from 'dayjs';
import { User } from 'firebase/auth';

import { Region } from '@/api/gen/neutral_diet/food/v1/region_pb';
import { WeightUnit as WeightUnitProto } from '@/api/gen/neutral_diet/user/v1/food_item_log_pb';
import {
  UserSettings,
  UserSettings_DietaryRequirement,
} from '@/api/gen/neutral_diet/user/v1/user_pb';
import { ID_TOKEN_HEADER } from '@/api/transport';
import client from '@/api/user_service';
import { MIN_CF_LIMIT } from '@/config';
import DietaryRequirement from '@/core/dietary_requirements';
import { auth } from '@/core/firebase';
import { WeightUnit } from '@/core/weight';

import {
  FoodLogStats,
  LocalFoodLogItem,
  LocalUserSettings,
  SerializableDate,
  UserInsights,
} from './types';

export const CurrentUserState = atom<User | null>({
  key: 'CurrentUserState',
  dangerouslyAllowMutability: true,
  effects: [
    (ctx) => {
      if (ctx.trigger === 'get') {
        return auth.onAuthStateChanged((user) => {
          ctx.setSelf(user);
        });
      }
    },
  ],
});

export const CurrentUserTokenIDState = selector({
  key: 'CurrentUserTokenIDState',
  get: async ({ get }) => {
    const idToken = await get(CurrentUserState)?.getIdToken();
    return idToken;
  },
});

export const CurrentUserHeadersState = selector({
  key: 'CurrentUserHeadersState',
  get: async ({ get }) => {
    const headers = new Headers();
    const idToken = get(CurrentUserTokenIDState);
    if (idToken) {
      headers.set(ID_TOKEN_HEADER, idToken);
    }
    return headers;
  },
});

export const LocalUserSettingsState = atom<LocalUserSettings>({
  key: 'LocalUserSettingsState',
  default: selector({
    key: 'LocalUserSettingsState/Default',
    get: async ({ get }) => {
      const defaults: LocalUserSettings = {
        region: Region.UNSPECIFIED,
        cfLimit: MIN_CF_LIMIT,
        dirty: false,
        dietaryRequirement: UserSettings_DietaryRequirement.UNSPECIFIED,
      };

      const userHeaders = get(CurrentUserHeadersState);
      try {
        const response = await client.getUserSettings({}, { headers: userHeaders });
        defaults.cfLimit = response.userSettings?.cfLimit ?? defaults.cfLimit;
        defaults.region = response.userSettings?.region ?? defaults.region;
        defaults.dietaryRequirement =
          response.userSettings?.dietaryRequirement ?? defaults.dietaryRequirement;
      } catch (err) {
        console.error(err);
      }
      return defaults;
    },
  }),
});

export const RemoteUserSettingsState = selector({
  key: 'RemoteUserSettingsState',
  get: ({ get }) => {
    const localUserSettings = get(LocalUserSettingsState);
    return new UserSettings({
      region: localUserSettings.region,
      cfLimit: localUserSettings.cfLimit,
      dietaryRequirement: localUserSettings.dietaryRequirement,
    });
  },
});

export const DietaryRequirementsState = selector({
  key: 'DietaryRequirementsState',
  get: () => {
    return Object.values(UserSettings_DietaryRequirement)
      .filter((x) => typeof x === 'number')
      .map((dr) => new DietaryRequirement(dr as UserSettings_DietaryRequirement));
  },
});

export const WeightUnitsState = selector({
  key: 'WeightUnitsState',
  get: () => {
    return Object.values(WeightUnitProto)
      .filter((x) => typeof x === 'number')
      .map((w) => new WeightUnit(w as WeightUnitProto));
  },
});

export const FoodItemLogDateState = atom<dayjs.Dayjs>({
  key: 'FoodItemLogDateState',
  default: dayjs(),
});

export const LocalFoodItemLogState = atomFamily<LocalFoodLogItem[], SerializableDate>({
  key: 'LocalFoodItemLogState',
  default: selectorFamily({
    key: 'LocalFoodItemLogState/Default',
    get:
      (date) =>
      async ({ get }) => {
        const userHeaders = get(CurrentUserHeadersState);
        try {
          const response = await client.getFoodItemLog(
            {
              date: { year: date.year, month: date.month, day: date.day },
            },
            { headers: userHeaders },
          );

          return response.foodItemLog.map((foodLogItem) => {
            return {
              dbId: foodLogItem.id,
              foodItemId: foodLogItem.foodItemId,
              name: foodLogItem.name,
              weight: {
                value: foodLogItem.weight,
                unit: new WeightUnit(foodLogItem.weightUnit),
              },
              carbonFootprint: foodLogItem.carbonFootprint,
              region: foodLogItem.region,
            };
          });
        } catch (err) {
          console.error(err);
          return [];
        }
      },
  }),
});

export const LocalFoodItemLogStats = selectorFamily<FoodLogStats, SerializableDate>({
  key: 'LocalFoodItemLogStats',
  get:
    (date) =>
    async ({ get }) => {
      const foodItemLog = get(LocalFoodItemLogState(date));
      const userSettings = get(LocalUserSettingsState);
      const stats = {
        totalCarbonFootprint: 0.0,
        carbonFootprintGoalPercent: 0,
        carbonFootprintRemaining: userSettings.cfLimit,
      };
      foodItemLog.forEach((item) => {
        stats.totalCarbonFootprint += item.carbonFootprint;
      });
      stats.carbonFootprintGoalPercent =
        userSettings.cfLimit != 0 ? (stats.totalCarbonFootprint / userSettings.cfLimit) * 100 : 0;
      stats.carbonFootprintRemaining -= stats.totalCarbonFootprint;
      return stats;
    },
});

export const UserInsightsState = selector<UserInsights>({
  key: 'UserInsightsState',
  get: async ({ get }) => {
    const userHeaders = get(CurrentUserHeadersState);
    const response = await client.getUserInsights({}, { headers: userHeaders });

    return {
      overallUser: response.overallCarbonFootprint,
      noUserEntries: response.noEntries,
      overallUserAverage: response.overallCarbonFootprint / response.noEntries || 0,
      dailyGlobalAverage: response.dailyAverageCarbonFootprintOverall,
      dailyGlobalAverageUserDietaryRequirement:
        response.dailyAverageCarbonFootprintDietaryRequirement,
      activeStreak: response.activeStreak,
    };
  },
});
