import { Link } from 'react-router-dom';

import DefaultIcon from '@mui/icons-material/Deblur';
import {
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Toolbar,
} from '@mui/material';
import BottomNavigation from '@mui/material/BottomNavigation';
import BottomNavigationAction from '@mui/material/BottomNavigationAction';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';

import isMobile from 'is-mobile';

import { useNavigation } from '@/hooks/useNavigation';
import routes from '@/routes';

function Navigation() {
  const [navigationValue, navigationActions] = useNavigation();
  const mobile = isMobile();

  return (
    <Box>
      {mobile ? (
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
      ) : (
        <Drawer
          variant="permanent"
          sx={{
            width: 240,
          }}
        >
          <Toolbar />
          <Box sx={{ overflow: 'auto' }}>
            <List>
              {Object.values(routes)
                .filter((route) => route.title)
                .filter((route) => route.navigation)
                .map(({ path, title, icon: Icon }) => (
                  <ListItem disablePadding key={path}>
                    <ListItemButton component={Link} to={path as string}>
                      <ListItemIcon>{Icon ? <Icon /> : <DefaultIcon />}</ListItemIcon>
                      <ListItemText primary={title} />
                    </ListItemButton>
                  </ListItem>
                ))}
            </List>
          </Box>
        </Drawer>
      )}
    </Box>
  );
}

export default Navigation;
