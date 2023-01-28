import { atom } from 'recoil';

import { Themes } from '@/theme/types';
import { recoilPersist } from 'recoil-persist';

const isDarkMode = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;

const { persistAtom } = recoilPersist({ key: 'theme-mode' });

const ThemeModeState = atom({
  key: 'ThemeModeState',
  default: isDarkMode ? ('dark' as Themes) : ('light' as Themes),
  effects: [persistAtom],
});

export default ThemeModeState;
