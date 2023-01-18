import { Navigate } from 'react-router-dom';

import { Box } from '@mui/material';

import Loading from '@/components/Loading';
import SignUp from '@/components/SignUp';
import { FullSizeCenteredFlexBox } from '@/components/styled';
import { auth } from '@/core/firebase';
import useIdToken from '@/hooks/useIdToken';
import routes from '@/routes';
import { Pages } from '@/routes/types';

function SignUpPage() {
  const [user, loading, error] = useIdToken(auth);

  if (loading) {
    return (
      <>
        <Loading />
      </>
    );
  }

  if (error) {
    return (
      <div>
        <p>Error:</p>
        <>{error}</>
      </div>
    );
  }

  if (user) {
    return <Navigate to={routes[Pages.Account].path} />;
  }

  return (
    <>
      <FullSizeCenteredFlexBox>
        <Box>
          <SignUp />
        </Box>
      </FullSizeCenteredFlexBox>
    </>
  );
}

export default SignUpPage;
