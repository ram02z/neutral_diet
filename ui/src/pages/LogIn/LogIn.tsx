import { Link, Navigate } from 'react-router-dom';

import { Divider, Link as MuiLink, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import Loading from '@/components/Loading';
import LogIn from '@/components/LogIn';
import { useCurrentUser } from '@/hooks/useCurrentUser';
import routes from '@/routes';
import { Pages } from '@/routes/types';
import GoogleAuthButton from '@/components/GoogleAuthButton';

function LogInPage() {
  const user = useCurrentUser();

  if (user === undefined) {
    return (
      <>
        <Loading />
      </>
    );
  } else if (user === null) {
    return (
      <Grid
        container
        direction="column"
        columns={10}
        justifyContent="center"
        alignItems="center"
        height="100%"
        spacing={2}
      >
        <Typography variant="h4"> Welcome back</Typography>
        <Grid xs={7} sm={6} md={5} lg={4} xl={3}>
          <LogIn />
        </Grid>
        <Grid>
          <Typography>
            {"Don't have an account? "}
            <MuiLink component={Link} to="/signup">
              Sign up
            </MuiLink>
          </Typography>
        </Grid>
        <Grid xs={7} sm={6} md={5} lg={4} xl={3}>
          <Divider flexItem>
          <Typography variant="overline">
            OR
            </Typography>
          </Divider>
        </Grid>
        <Grid xs={7} sm={6} md={5} lg={4} xl={3}>
          <GoogleAuthButton/>
        </Grid>
      </Grid>
    );
  } else {
    return <Navigate to={routes[Pages.Account].path} />;
  }
}

export default LogInPage;
