import { UserSettings } from "@/api/gen/neutral_diet/user/v1/user_pb";
import { ID_TOKEN_HEADER } from "@/api/transport";
import client from "@/api/user_service";
import { auth } from "@/core/firebase";
import { useCurrentUser } from "@/hooks/useCurrentUser";
import { User } from "firebase/auth";
import { atom } from "recoil";

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

export const CurrentUserSettingsState = atom<UserSettings | null>({
  key: 'CurrentUserSettingsState',
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
