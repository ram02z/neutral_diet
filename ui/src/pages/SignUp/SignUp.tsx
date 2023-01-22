import { Navigate } from 'react-router-dom';

import { Box } from '@mui/material';

import Loading from '@/components/Loading';
import SignUp from '@/components/SignUp';
import { FullSizeCenteredFlexBox } from '@/components/styled';
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
        <FullSizeCenteredFlexBox>
          <Box>
            <SignUp />
          </Box>
        </FullSizeCenteredFlexBox>
      </>
    );
  } else {
    return <Navigate to={routes[Pages.Account].path} />;
  }
}

export default SignUpPage;
