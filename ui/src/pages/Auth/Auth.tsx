import { Link, useOutlet } from 'react-router-dom';

import { Button, Stack } from '@mui/material';

function Auth() {
  const outlet = useOutlet();

  if (outlet) {
    return outlet;
  } else {
    return (
      <Stack direction="row" justifyContent="center" alignItems="center" height="100%" spacing={2}>
        <Button component={Link} to="login" variant="contained">
          Log in
        </Button>
        <Button component={Link} to="signup" variant="contained">
          Sign up
        </Button>
      </Stack>
    );
  }
}

export default Auth;
