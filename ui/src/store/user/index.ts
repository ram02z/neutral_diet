import { atom, atomFamily, selector, selectorFamily } from 'recoil';

import dayjs from 'dayjs';
import { User, updateProfile } from 'firebase/auth';

import { Region } from '@/api/gen/neutral_diet/food/v1/region_pb';
import { Meal as MealProto, Unit } from '@/api/gen/neutral_diet/user/v1/food_item_log_pb';
import {
  UserSettings,
  UserSettings_DietaryRequirement,
} from '@/api/gen/neutral_diet/user/v1/user_pb';
import client from '@/api/user_service';
import { MIN_CF_LIMIT } from '@/config';
import DietaryRequirement from '@/core/dietary_requirements';
import { auth } from '@/core/firebase';
import { FoodUnit } from '@/core/food_unit';
import { Meal } from '@/core/meal';
import NotificationService from '@/core/notifications';
import { toDayJsDate, toSerializableDate } from '@/utils/date';

import {
  FoodLogStats,
  LocalFoodLogItem,
  LocalUserGoal,
  LocalUserGoals,
  LocalUserSettings,
  SerializableDate,
  UserInsights,
  UserProgress,
} from './types';

export const CurrentUserState = atom<User | null>({
  key: 'CurrentUserState',
  dangerouslyAllowMutability: true,
  effects: [
    ({ trigger, setSelf }) => {
      if (trigger === 'get') {
        return auth.onIdTokenChanged((user) => {
          setSelf(user);
        });
      }
    },
  ],
});

export const CurrentUserDisplayName = atom<string | null>({
  key: 'CurrentUserDisplayName',
  default: selector({
    key: 'CurrentUserDisplayName/Default',
    get: ({ get }) => {
      const user = get(CurrentUserState);
      return user?.displayName ?? null;
    },
  }),
  effects: [
    ({ onSet, getPromise }) => {
      onSet((newValue, oldValue) => {
        if (newValue === oldValue) return;
        if (newValue === null) return;
        getPromise(CurrentUserState).then((user) => {
          if (!user) return;
          updateProfile(user, { displayName: newValue });
        });
      });
    },
  ],
});

export const CurrentUserHeadersState = selector({
  key: 'CurrentUserHeadersState',
  get: async ({ get }) => {
    const headers = new Headers();
    const user = get(CurrentUserState);
    user?.getIdToken().then((idToken) => {
      headers.set('Authorization', `Bearer ${idToken}`);
    });
    try {
      NotificationService.getWebToken().then((token) => {
        client.addDevice({ fcmToken: token }, { headers: headers });
      });
    } catch {
      /* empty */
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
      const response = await client.getUserSettings({}, { headers: userHeaders });
      defaults.cfLimit = response.userSettings?.cfLimit ?? defaults.cfLimit;
      defaults.region = response.userSettings?.region ?? defaults.region;
      defaults.dietaryRequirement =
        response.userSettings?.dietaryRequirement ?? defaults.dietaryRequirement;
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

export const FoodUnitsState = selector({
  key: 'FoodUnitsState',
  get: () => {
    return Object.values(Unit)
      .filter((x) => typeof x === 'number')
      .map((u) => new FoodUnit(u as Unit));
  },
});

export const MealsState = selector({
  key: 'MealsState',
  get: () => {
    return Object.values(MealProto)
      .filter((x) => typeof x === 'number')
      .map((m) => new Meal(m as MealProto))
      .sort((m1, m2) => m1.getOrder() - m2.getOrder());
  },
});

export const FoodItemLogDateState = atom<dayjs.Dayjs>({
  key: 'FoodItemLogDateState',
  default: dayjs(),
});

export const FoodItemLogSerializableDateState = selector({
  key: 'FoodItemLogSerializableDateState',
  get: ({ get }) => {
    const date = get(FoodItemLogDateState);

    return toSerializableDate(date);
  },
});

export const LocalFoodItemLogState = atomFamily<LocalFoodLogItem[], SerializableDate>({
  key: 'LocalFoodItemLogState',
  default: selectorFamily({
    key: 'LocalFoodItemLogState/Default',
    get:
      (date) =>
      async ({ get }) => {
        const userHeaders = get(CurrentUserHeadersState);
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
            quantity: {
              value: foodLogItem.quantity,
              unit: new FoodUnit(foodLogItem.unit),
            },
            carbonFootprint: foodLogItem.carbonFootprint,
            region: foodLogItem.region,
            meal: foodLogItem.meal,
          };
        });
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

export const LocalFoodItemLogStatsByMeal = selectorFamily<Map<number, number>, SerializableDate>({
  key: 'LocalFoodItemLogStatsByMeal',
  get:
    (date) =>
    async ({ get }) => {
      const foodItemLog = get(LocalFoodItemLogState(date));
      const mealMap = new Map(
        Object.values(MealProto)
          .filter((x) => typeof x === 'number')
          .fill(0.0)
          .entries(),
      ) as Map<number, number>;
      foodItemLog.forEach((item) => {
        const value = mealMap.get(item.meal) ?? 0.0;
        mealMap.set(item.meal, value + item.carbonFootprint);
      });
      return mealMap;
    },
});

export const UserInsightsState = selector<UserInsights>({
  key: 'UserInsightsState',
  get: async ({ get }) => {
    const userHeaders = get(CurrentUserHeadersState);
    const response = await client.getUserInsights({}, { headers: userHeaders });

    return {
      userDailyAverage: response.dailyUserAverage,
      dailyGlobalAverage: response.dailyAverageCarbonFootprintOverall,
      dailyGlobalAverageUserDietaryRequirement:
        response.dailyAverageCarbonFootprintDietaryRequirement,
      streakLength: response.streakLen,
      isStreakActive: response.isStreakActive,
    };
  },
});

export const UserProgressState = selector<UserProgress>({
  key: 'UserProgressState',
  get: async ({ get }) => {
    const userHeaders = get(CurrentUserHeadersState);
    const response = await client.getUserProgress({}, { headers: userHeaders });
    const mealMap = new Map<number, Record<string, number>>();
    mealMap.set(MealProto.BREAKFAST, response.dailyProgressBreakfast);
    mealMap.set(MealProto.LUNCH, response.dailyProgressLunch);
    mealMap.set(MealProto.DINNER, response.dailyProgressDinner);
    mealMap.set(MealProto.UNSPECIFIED, response.dailyProgressSnacks);

    return { all: response.dailyProgressAll, meal: mealMap };
  },
});

export const UserGoalsState = atom<LocalUserGoals>({
  key: 'UserGoalsState',
  default: selector({
    key: 'UserGoalsState/Default',
    get: async ({ get }) => {
      const userHeaders = get(CurrentUserHeadersState);
      const response = await client.getCarbonFootprintGoals({}, { headers: userHeaders });

      return {
        active: response.active.map((goal) => {
          return {
            dbId: goal.id,
            description: goal.description,
            startDate: toDayJsDate(goal.startDate as SerializableDate),
            endDate: toDayJsDate(goal.endDate as SerializableDate),
            startCarbonFootprint: goal.startCarbonFootprint,
            targetCarbonFootprint: goal.targetCarbonFootprint,
          };
        }),
        completed: response.completed.map((goal) => {
          return {
            dbId: goal.id,
            description: goal.description,
            startDate: toDayJsDate(goal.startDate as SerializableDate),
            endDate: toDayJsDate(goal.endDate as SerializableDate),
            startCarbonFootprint: goal.startCarbonFootprint,
            targetCarbonFootprint: goal.targetCarbonFootprint,
          };
        }),
      };
    },
  }),
});

export const ActiveGoalState = selector<LocalUserGoal | null>({
  key: 'ActiveGoalState',
  get: ({ get }) => {
    const goals = get(UserGoalsState);
    if (goals.active.length === 0) return null;
    return goals.active.reduce((a, b) => (a.endDate.isBefore(b.endDate) ? a : b));
  },
});

export const SelectedUserGoalState = atom<LocalUserGoal | null>({
  key: 'SelectedUserGoalState',
  default: null,
});

export const UserGoalProgressState = selector<UserProgress | null>({
  key: 'UserGoalProgressState',
  get: async ({ get }) => {
    const goal = get(SelectedUserGoalState);
    if (!goal) return null;
    const userHeaders = get(CurrentUserHeadersState);
    const response = await client.getUserProgress(
      {
        dateRange: {
          start: toSerializableDate(goal.startDate),
          end: toSerializableDate(goal.endDate),
        },
      },
      { headers: userHeaders },
    );
    const mealMap = new Map<number, Record<string, number>>();
    mealMap.set(MealProto.BREAKFAST, response.dailyProgressBreakfast);
    mealMap.set(MealProto.LUNCH, response.dailyProgressLunch);
    mealMap.set(MealProto.DINNER, response.dailyProgressDinner);
    mealMap.set(MealProto.UNSPECIFIED, response.dailyProgressSnacks);

    return { all: response.dailyProgressAll, meal: mealMap };
  },
});
