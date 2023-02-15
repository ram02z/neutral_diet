import { atom, selector } from 'recoil';

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

import { LocalUserSettings } from './types';

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

      const idToken = get(CurrentUserTokenIDState);
      if (idToken) {
        const headers = new Headers();
        headers.set(ID_TOKEN_HEADER, idToken);
        const response = await client.getUserSettings({}, { headers: headers });
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
})

export const FoodItemLogQuery = selector({
  key: 'FoodItemLogQuery',
  get:
    async ({ get }) => {
      const idToken = get(CurrentUserTokenIDState);
      const date = get(FoodItemLogDateState)
      if (idToken) {
        const headers = new Headers();
        headers.set(ID_TOKEN_HEADER, idToken);
        const response = await client.getFoodItemLog(
          {
            date: { year: date.year(), month: date.month() + 1, day: date.date() },
          },
          { headers: headers },
        );
        return response.foodItemLog;
      }

      return [];
    },
});
