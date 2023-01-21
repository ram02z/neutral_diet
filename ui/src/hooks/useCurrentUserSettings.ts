import { atom, useRecoilValueLoadable } from 'recoil';

import { UserSettings } from '@/api/gen/neutral_diet/user/v1/user_pb';
import { ID_TOKEN_HEADER } from '@/api/transport';
import client from '@/api/user_service';

import { useCurrentUser } from './useCurrentUser';

export const CurrentUserSettings = atom<UserSettings | null>({
  key: 'CurrentUserSettings',
  dangerouslyAllowMutability: true,
  effects: [
    (ctx) => {
      if (ctx.trigger === 'get') {
        const user = useCurrentUser();
        user?.getIdToken().then((idToken) => {
          const headers = new Headers();
          headers.set(ID_TOKEN_HEADER, idToken);
          client.getUserSettings({}, { headers: headers }).then((response) => {
            if (response.userSettings) ctx.setSelf(response.userSettings);
          });
        });
      }
    },
  ],
});

export function useCurrentUserSettings() {
  const value = useRecoilValueLoadable(CurrentUserSettings);
  return value.state === 'loading' ? undefined : value.valueOrThrow();
}
