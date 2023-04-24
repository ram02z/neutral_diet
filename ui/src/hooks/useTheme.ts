import { useRecoilState } from 'recoil';

import ThemeModeState from '@/store/theme';
import { Actions } from '@/store/theme/types';
import { Themes } from '@/theme/types';

/**
 * Provides access to current theme mode state and an action to toggle it.
 */
export function useTheme(): [Themes, Actions] {
  const [themeMode, setThemeMode] = useRecoilState(ThemeModeState);

  function toggle() {
    setThemeMode((mode: Themes) => (mode === Themes.DARK ? Themes.LIGHT : Themes.DARK));
  }

  return [themeMode, { toggle }];
}
