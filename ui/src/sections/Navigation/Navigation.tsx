import { Link } from 'react-router-dom';

import DefaultIcon from '@mui/icons-material/Deblur';
import BottomNavigation from '@mui/material/BottomNavigation';
import BottomNavigationAction from '@mui/material/BottomNavigationAction';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';

import { useNavigation } from '@/hooks/useNavigation';
import routes from '@/routes';

function Navigation() {
  const [navigationValue, navigationActions] = useNavigation();

  return (
    <Box>
      <Paper sx={{ position: 'fixed', bottom: 0, left: 0, right: 0 }} elevation={3}>
        <BottomNavigation value={navigationValue} onChange={navigationActions.change}>
          {Object.values(routes)
            .filter((route) => route.title)
            .filter((route) => route.navigation)
            .map(({ path, title, icon: Icon }) => (
              <BottomNavigationAction
                key={path}
                component={Link}
                to={path as string}
                value={title}
                icon={Icon ? <Icon /> : <DefaultIcon />}
              />
            ))}
        </BottomNavigation>
      </Paper>
    </Box>
  );
}

export default Navigation;
