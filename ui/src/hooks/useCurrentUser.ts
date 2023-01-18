import { atom, useRecoilValueLoadable } from 'recoil';

import { User } from 'firebase/auth';

import { auth } from '@/core/firebase';

export const CurrentUser = atom<User | null>({
  key: 'CurrentUser',
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

/**
 * The currently logged-in (authenticated) user object.
 *
 * @example
 *   const { useCurrentUser } from "../core/auth.js";
 *
 *   function Example(): JSX.Element {
 *     const me = useCurrentUser();
 *     // => { uid: "xxx", email: "me@example.com", ... }
 *     // => Or, `null` when not authenticated
 *     // => Or, `undefined` when not initialized
 *   }
 */
export function useCurrentUser() {
  const value = useRecoilValueLoadable(CurrentUser);
  return value.state === 'loading' ? undefined : value.valueOrThrow();
}
