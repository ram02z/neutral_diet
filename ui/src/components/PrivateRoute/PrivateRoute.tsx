import { FC } from 'react';
import { ErrorBoundary } from 'react-error-boundary';
import { Navigate } from 'react-router';
import { useSetRecoilState } from 'recoil';

import { useCurrentUser } from '@/hooks/useCurrentUser';
import ErrorPage from '@/pages/Error';
import routes from '@/routes';
import { Pages } from '@/routes/types';
import { PrivateRoutePathState } from '@/store/location';

const PrivateRoute: FC<{ children: JSX.Element }> = ({ children }) => {
  const user = useCurrentUser();
  const setPath = useSetRecoilState(PrivateRoutePathState);
  if (user === null) {
    setPath(location.pathname);
    return <Navigate to={routes[Pages.Auth].path} replace />;
  }

  return <ErrorBoundary FallbackComponent={ErrorPage}>{children}</ErrorBoundary>;
};

export default PrivateRoute;
