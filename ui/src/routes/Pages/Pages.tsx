import { Route, Routes } from 'react-router-dom';

import Box from '@mui/material/Box';

import { NAVIGATION_DRAWER_WIDTH } from '@/config';
import useLayout from '@/hooks/useLayout';

import routes from '..';
import { getPageHeight } from './utils';

function Pages() {
  const { desktopLayout } = useLayout();

  const marginLeft = desktopLayout ? NAVIGATION_DRAWER_WIDTH : 0;

  return (
    <Box sx={{ height: (theme) => getPageHeight(theme), p: 3, ml: marginLeft }}>
      <Routes>
        {Object.values(routes).map(({ path, component: Component }) => {
          return <Route key={path} path={path} element={<Component />} />;
        })}
      </Routes>
    </Box>
  );
}

export default Pages;
