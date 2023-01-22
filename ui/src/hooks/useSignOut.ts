import { useCallback } from 'react';

import { auth } from '@/core/firebase';

export function useSignOut() {
  return useCallback(() => auth.signOut(), []);
}
