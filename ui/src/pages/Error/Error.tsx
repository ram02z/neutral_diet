import { useEffect } from 'react';

import { Code, ConnectError } from '@bufbuild/connect-web';
import { useSnackbar } from 'notistack';

import { useSignOut } from '@/hooks/useSignOut';

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

  return <></>;
}

export default ErrorPage;
