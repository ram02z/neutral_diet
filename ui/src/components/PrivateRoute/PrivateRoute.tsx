import { FC } from 'react';
import { ErrorBoundary } from 'react-error-boundary';
import { Navigate } from 'react-router';
import { useSetRecoilState } from 'recoil';

import { Code, ConnectError } from '@bufbuild/connect-web';
import { useSnackbar } from 'notistack';

import { useCurrentUser } from '@/hooks/useCurrentUser';
import { useSignOut } from '@/hooks/useSignOut';
import routes from '@/routes';
import { Pages } from '@/routes/types';
import { PrivateRoutePathState } from '@/store/location';

const PrivateRoute: FC<{ children: JSX.Element }> = ({ children }) => {
  const user = useCurrentUser();
  const signOut = useSignOut();
  const { enqueueSnackbar } = useSnackbar();
  const setPath = useSetRecoilState(PrivateRoutePathState);
  if (user === null) {
    setPath(location.pathname);
    return <Navigate to={routes[Pages.Auth].path} replace />;
  }

  return (
    <ErrorBoundary
      onError={(error) => {
        signOut();
        if (error instanceof ConnectError && error.code == Code.Unauthenticated) {
          enqueueSnackbar('Session expired. Log in again to continue.', { variant: 'warning' });
        } else {
          enqueueSnackbar('Something went wrong. Please try again later.', { variant: 'error' });
        }
      }}
      fallbackRender={() => <></>}
    >
      {children}
    </ErrorBoundary>
  );
};

export default PrivateRoute;
