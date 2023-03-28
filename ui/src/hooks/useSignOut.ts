import { useCallback } from 'react';
import { useSetRecoilState } from 'recoil';

import { auth } from '@/core/firebase';
import { CurrentUserDisplayName } from '@/store/user';

export function useSignOut() {
  const setUserDisplayName = useSetRecoilState(CurrentUserDisplayName);
  return useCallback(() => {
    auth.signOut();
    setUserDisplayName(null);
  }, [setUserDisplayName]);
}
