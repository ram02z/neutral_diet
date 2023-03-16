import { Link, Navigate, useOutlet } from 'react-router-dom';

import { Button, Stack, Typography } from '@mui/material';

import { useCurrentUser } from '@/hooks/useCurrentUser';
import routes from '@/routes';
import { Pages } from '@/routes/types';

function Auth() {
  const outlet = useOutlet();
  const user = useCurrentUser();
  if (user !== null) return <Navigate to={routes[Pages.Account].path} />;
  if (outlet) return outlet;

  return (
    <Stack direction="column" justifyContent="center" alignItems="center" height="100%" spacing={2}>
      <Typography>Welcome to Neutral Diet</Typography>
      <Typography>Log in with your account to continue</Typography>
      <Stack direction="row" justifyContent="center" alignItems="center" spacing={2}>
        <Button component={Link} to="login" variant="contained">
          Log in
        </Button>
        <Button component={Link} to="signup" variant="contained">
          Sign up
        </Button>
      </Stack>
    </Stack>
  );
}

export default Auth;
