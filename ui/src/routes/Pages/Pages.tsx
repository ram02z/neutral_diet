import { ErrorBoundary } from 'react-error-boundary';
import { Route, Routes } from 'react-router-dom';

import Box from '@mui/material/Box';

import PrivateRoute from '@/components/PrivateRoute';
import { NAVIGATION_DRAWER_WIDTH } from '@/config';
import useLayout from '@/hooks/useLayout';
import ErrorPage from '@/pages/Error';
import routes from '@/routes';

import { getPageHeight } from './utils';

function Pages() {
  const { desktopLayout } = useLayout();

  const marginLeft = desktopLayout ? NAVIGATION_DRAWER_WIDTH : 0;

  return (
    <Box sx={{ height: (theme) => getPageHeight(theme), p: 3, ml: marginLeft }}>
      <ErrorBoundary FallbackComponent={ErrorPage}>
        <Routes>
          {Object.values(routes).map(
            ({ path, component: Component, subComponents, requireAuth }) => {
              return (
                <Route
                  key={path}
                  path={path}
                  element={
                    requireAuth ? (
                      <PrivateRoute>
                        <Component />
                      </PrivateRoute>
                    ) : (
                      <Component />
                    )
                  }
                >
                  {subComponents.map(({ path, component: SubComponent }) => {
                    return (
                      <Route
                        key={path}
                        path={path}
                        element={
                          requireAuth ? (
                            <PrivateRoute>
                              <SubComponent />
                            </PrivateRoute>
                          ) : (
                            <SubComponent />
                          )
                        }
                      />
                    );
                  })}
                </Route>
              );
            },
          )}
        </Routes>
      </ErrorBoundary>
    </Box>
  );
}

export default Pages;
