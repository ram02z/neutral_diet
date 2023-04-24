import { useCallback } from 'react';
import { useSetRecoilState } from 'recoil';

import { auth } from '@/core/firebase';
import { CurrentUserDisplayName } from '@/store/user';

/**
 * Sign out current user.
 * Wraps the underlying auth.signOut method and provides additional loading and error information.
 */
export function useSignOut() {
  const setUserDisplayName = useSetRecoilState(CurrentUserDisplayName);
  return useCallback(() => {
    auth.signOut();
    setUserDisplayName(null);
  }, [setUserDisplayName]);
}
