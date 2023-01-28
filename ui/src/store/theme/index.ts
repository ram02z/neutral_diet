import { atom } from 'recoil';

import { Themes } from '@/theme/types';
import { persistAtom } from '@/store';

const isDarkMode = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;

const ThemeModeState = atom({
  key: 'ThemeModeState',
  default: isDarkMode ? ('dark' as Themes) : ('light' as Themes),
  effects: [persistAtom],
});

export default ThemeModeState;
