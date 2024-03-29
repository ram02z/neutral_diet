import { Link } from 'react-router-dom';
import { useLocation } from 'react-router-dom';

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

import { NAVIGATION_DRAWER_WIDTH } from '@/config';
import useLayout from '@/hooks/useLayout';
import routes from '@/routes';

function Navigation() {
  const location = useLocation();
  const { desktopLayout } = useLayout();

  return (
    <>
      {desktopLayout ? (
        <Drawer
          variant="permanent"
          PaperProps={{
            elevation: 0,
            sx: {
              width: NAVIGATION_DRAWER_WIDTH,
            },
          }}
          sx={{
            width: NAVIGATION_DRAWER_WIDTH,
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
                    <ListItemButton
                      selected={path == location.pathname}
                      component={Link}
                      to={path as string}
                    >
                      <ListItemIcon>{Icon ? <Icon /> : <DefaultIcon />}</ListItemIcon>
                      <ListItemText primary={title} />
                    </ListItemButton>
                  </ListItem>
                ))}
            </List>
          </Box>
        </Drawer>
      ) : (
        <Paper sx={{ position: 'fixed', bottom: 0, left: 0, right: 0 }} elevation={3}>
          <BottomNavigation value={location.pathname}>
            {Object.values(routes)
              .filter((route) => route.title)
              .filter((route) => route.navigation)
              .map(({ path, icon: Icon }) => (
                <BottomNavigationAction
                  key={path}
                  component={Link}
                  to={path as string}
                  value={path}
                  icon={Icon ? <Icon /> : <DefaultIcon />}
                />
              ))}
          </BottomNavigation>
        </Paper>
      )}
    </>
  );
}

export default Navigation;
