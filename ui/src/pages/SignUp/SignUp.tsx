import { Link, Navigate } from 'react-router-dom';

import { Link as MuiLink, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import Loading from '@/components/Loading';
import SignUp from '@/components/SignUp';
import { useCurrentUser } from '@/hooks/useCurrentUser';
import routes from '@/routes';
import { Pages } from '@/routes/types';

function SignUpPage() {
  const user = useCurrentUser();

  if (user === undefined) {
    return (
      <>
        <Loading />
      </>
    );
  } else if (user === null) {
    return (
      <>
        <Grid
          container
          direction="column"
          columns={10}
          justifyContent="center"
          alignItems="center"
          height="100%"
        >
          <Typography variant="h4">Create your account</Typography>
          <Grid xs={7} sm={6} md={5} lg={4} xl={3}>
            <SignUp />
          </Grid>
          <Grid>
            <MuiLink component={Link} to="/login">
              {"Already have an account? Sign in"}
            </MuiLink>
        </Grid>
        </Grid>
      </>
    );
  } else {
    return <Navigate to={routes[Pages.Account].path} />;
  }
}

export default SignUpPage;
