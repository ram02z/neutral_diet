import { Actions } from "@/store/theme/types";
import ThemeModeState from "@/store/theme";
import { Themes } from "@/theme/types";
import { useRecoilState } from "recoil";

export function useTheme(): [Themes, Actions] {
  const [themeMode, setThemeMode] = useRecoilState(ThemeModeState);

  function toggle() {
    setThemeMode((mode: Themes) => (mode === Themes.DARK ? Themes.LIGHT : Themes.DARK));
  }

  return [themeMode, { toggle }];
}
