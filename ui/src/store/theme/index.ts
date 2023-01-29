import { atom } from 'recoil';

import { persistAtom } from '@/store';
import { Themes } from '@/theme/types';

const isDarkMode = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;

const ThemeModeState = atom({
  key: 'ThemeModeState',
  default: isDarkMode ? ('dark' as Themes) : ('light' as Themes),
  effects: [persistAtom],
});

export default ThemeModeState;
