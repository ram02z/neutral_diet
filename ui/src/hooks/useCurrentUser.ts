import { useRecoilValueLoadable } from 'recoil';

import { CurrentUserState } from '@/store/user';

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
  const value = useRecoilValueLoadable(CurrentUserState);
  return value.state === 'loading' ? undefined : value.valueOrThrow();
}
