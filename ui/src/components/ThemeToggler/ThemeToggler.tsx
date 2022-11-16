import IconButton from '@mui/material/IconButton';

import useTheme from '@/store/theme';
import Tooltip from '@mui/material/Tooltip';
import DarkModeIcon from '@mui/icons-material/DarkMode';
import LightModeIcon from '@mui/icons-material/LightMode';

function ThemeToggler() {
  const [theme, themeActions] = useTheme();

  return (
    <Tooltip title={`Toggle ${theme}`}>
      <IconButton color="inherit" onClick={themeActions.toggle}>
        {theme === 'dark' ? <DarkModeIcon /> : <LightModeIcon />}
      </IconButton>
    </Tooltip>

  );
}

export default ThemeToggler;
