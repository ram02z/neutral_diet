import { useRecoilValueLoadable } from 'recoil';

import { RemoteUserSettingsState } from '@/store/user';

export function useCurrentUserSettings() {
  const value = useRecoilValueLoadable(RemoteUserSettingsState);
  return value.state === 'loading' ? undefined : value.valueOrThrow();
}
