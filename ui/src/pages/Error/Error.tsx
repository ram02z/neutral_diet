import { useEffect } from 'react';
import { Link } from 'react-router-dom';

import { Button, Stack, Typography } from '@mui/material';

import { Code, ConnectError } from '@bufbuild/connect-web';
import { useSnackbar } from 'notistack';

import { useSignOut } from '@/hooks/useSignOut';
import routes from '@/routes';
import { Pages } from '@/routes/types';

type ErrorProps = {
  error: Error;
};

function ErrorPage({ error }: ErrorProps) {
  const signOut = useSignOut();
  const { enqueueSnackbar } = useSnackbar();

  useEffect(() => signOut(), [signOut]);

  if (error instanceof ConnectError && error.code == Code.Unauthenticated) {
    enqueueSnackbar('Session expired. Log in again to continue.', { variant: 'warning' });
  } else {
    enqueueSnackbar('Something went wrong. Please try again later.', { variant: 'error' });
  }
  return (
    <Stack spacing={2} justifyContent="center" alignItems="center" height="100%">
      <Typography>An error occured</Typography>
      <Button component={Link} to={routes[Pages.Auth].path}>
        Reauthenticate
      </Button>
    </Stack>
  );
}

export default ErrorPage;
