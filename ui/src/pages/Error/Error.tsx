import { Navigate, useRouteError } from 'react-router';
import { useSetRecoilState } from 'recoil';

import { Code, ConnectError } from '@bufbuild/connect-web';
import { useSnackbar } from 'notistack';

import { useSignOut } from '@/hooks/useSignOut';
import routes from '@/routes';
import { Pages } from '@/routes/types';
import { PrivateRoutePathState } from '@/store/location';

function ErrorPage() {
  const error = useRouteError();
  const signOut = useSignOut();
  const { enqueueSnackbar } = useSnackbar();
  const setPath = useSetRecoilState(PrivateRoutePathState);
  signOut();
  setPath(location.pathname);

  if (error instanceof ConnectError && error.code == Code.Unauthenticated) {
    enqueueSnackbar('Session expired. Log in again to continue.', { variant: 'warning' });
  } else {
    enqueueSnackbar('Something went wrong. Please try again later.', { variant: 'error' });
  }
  return <Navigate to={routes[Pages.Auth].path} replace />;
}

export default ErrorPage;
