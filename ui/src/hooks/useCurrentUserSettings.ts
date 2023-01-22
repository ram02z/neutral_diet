import { useRecoilValueLoadable } from 'recoil';

import { CurrentUserSettingsState } from '@/store/user';

export function useCurrentUserSettings() {
  const value = useRecoilValueLoadable(CurrentUserSettingsState);
  return value.state === 'loading' ? undefined : value.valueOrThrow();
}
