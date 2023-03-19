import { Navigate } from 'react-router';
import { useSetRecoilState } from 'recoil';

import { Typography } from '@mui/material';

import { Code, ConnectError } from '@bufbuild/connect-web';

import { useSignOut } from '@/hooks/useSignOut';
import routes from '@/routes';
import { Pages } from '@/routes/types';
import { PrivateRoutePathState } from '@/store/location';

type ErrorProps = {
  error: Error;
};

function ErrorPage({ error }: ErrorProps) {
  const signOut = useSignOut();
  const setPath = useSetRecoilState(PrivateRoutePathState);

  if (error instanceof ConnectError && error.code == Code.Unauthenticated) {
    signOut();
    setPath(location.pathname);
    return <Navigate to={routes[Pages.Auth].path} replace />;
  }
  return (
    <>
      <Typography variant="h3">Error</Typography>
    </>
  );
}

export default ErrorPage;
