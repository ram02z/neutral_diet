import { Link, Navigate } from 'react-router-dom';

import { Divider, Link as MuiLink, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import GoogleAuthButton from '@/components/GoogleAuthButton';
import Loading from '@/components/Loading';
import SignUp from '@/components/SignUp';
import { useCurrentUser } from '@/hooks/useCurrentUser';
import routes from '@/routes';
import { Pages } from '@/routes/types';
import { useRecoilValue } from 'recoil';
import { PrivateRoutePathState } from '@/store/location';

function SignUpPage() {
  const user = useCurrentUser();
  const path = useRecoilValue(PrivateRoutePathState);

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
          spacing={2}
        >
          <Typography variant="h4">Create your account</Typography>
          <Grid xs={7} sm={6} md={5} lg={4} xl={3}>
            <SignUp />
          </Grid>
          <Grid>
            <Typography>
              {'Already have an account? '}
              <MuiLink component={Link} to="../login">
                Sign in
              </MuiLink>
            </Typography>
          </Grid>
          <Grid xs={7} sm={6} md={5} lg={4} xl={3}>
            <Divider flexItem>
              <Typography variant="overline">OR</Typography>
            </Divider>
          </Grid>
          <Grid xs={7} sm={6} md={5} lg={4} xl={3}>
            <GoogleAuthButton />
          </Grid>
        </Grid>
      </>
    );
  } else {
    return <Navigate to={path ?? routes[Pages.Account].path} />;
  }
}

export default SignUpPage;
