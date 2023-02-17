import { atom, atomFamily, selector, selectorFamily } from 'recoil';

import dayjs from 'dayjs';
import { User } from 'firebase/auth';

import {
  UserSettings,
  UserSettings_DietaryRequirement,
} from '@/api/gen/neutral_diet/user/v1/user_pb';
import { ID_TOKEN_HEADER } from '@/api/transport';
import client from '@/api/user_service';
import DietaryRequirement from '@/core/dietary_requirements';
import { auth } from '@/core/firebase';

import { LocalFoodLogItem, LocalUserSettings } from './types';

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
    const idToken = get(CurrentUserTokenIDState);
    if (idToken) {
      const headers = new Headers();
      headers.set(ID_TOKEN_HEADER, idToken);
      return headers;
    }
  },
});

export const LocalUserSettingsState = atom<LocalUserSettings>({
  key: 'LocalUserSettingsState',
  default: selector({
    key: 'LocalUserSettingsState/Default',
    get: async ({ get }) => {
      const defaults: LocalUserSettings = {
        region: 'World',
        cfLimit: 0.0,
        dirty: false,
        dietaryRequirement: UserSettings_DietaryRequirement.UNSPECIFIED,
      };

      const userHeaders = get(CurrentUserHeadersState);
      if (userHeaders) {
        const response = await client.getUserSettings({}, { headers: userHeaders });
        defaults.cfLimit = response.userSettings?.cfLimit ?? defaults.cfLimit;
        defaults.region = response.userSettings?.region?.name ?? defaults.region;
        defaults.dietaryRequirement =
          response.userSettings?.dietaryRequirement ?? defaults.dietaryRequirement;
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
      region: { name: localUserSettings.region },
      cfLimit: localUserSettings.cfLimit,
      dietaryRequirement: localUserSettings.dietaryRequirement,
    });
  },
});

export const DietaryRequirementsState = atom<DietaryRequirement[]>({
  key: 'DietaryRequirementsState',
  default: selector({
    key: 'DietaryRequirementsState/Default',
    get: () => {
      return Object.values(UserSettings_DietaryRequirement)
        .filter((x) => typeof x === 'number')
        .map((dr) => new DietaryRequirement(dr as UserSettings_DietaryRequirement));
    },
  }),
});

export const FoodItemLogDateState = atom<dayjs.Dayjs>({
  key: 'FoodItemLogDateState',
  default: dayjs(),
});

export const LocalFoodItemLogState = atomFamily<LocalFoodLogItem[], dayjs.Dayjs>({
  key: 'LocalFoodItemLogState',
  default: selectorFamily({
    key: 'LocalFoodItemLogState/Default',
    get:
      (date) =>
      async ({ get }) => {
        const idToken = get(CurrentUserTokenIDState);
        if (idToken) {
          const headers = new Headers();
          headers.set(ID_TOKEN_HEADER, idToken);
          const response = await client.getFoodItemLog(
            {
              date: { year: date.year(), month: date.month() + 1, day: date.date() },
            },
            { headers: headers },
          );

          // TODO: error handling
          return response.foodItemLog.map((foodLogItem) => {
            return {
              dbId: foodLogItem.id,
              name: foodLogItem.name,
              weight: foodLogItem.weight,
              carbonFootprint: foodLogItem.carbonFootprint,
            };
          });
        }

        return [];
      },
  }),
});
