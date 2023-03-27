import { Link } from 'react-router-dom';

import { Button, Stack, Typography } from '@mui/material';

function NotFound() {
  return (
    <Stack spacing={2} justifyContent="center" alignItems="center" height="100%">
      <Typography variant="h1">404</Typography>
      <Typography variant="h5" color="text.secondary" gutterBottom>
        Page not found
      </Typography>
      <Button component={Link} to="/" variant="contained">
        Back to home
      </Button>
    </Stack>
  );
}

export default NotFound;
