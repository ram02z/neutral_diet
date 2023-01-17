import { Navigate } from 'react-router-dom';

import { Box } from '@mui/system';

import Loading from '@/components/Loading';
import LogIn from '@/components/LogIn';
import { FullSizeCenteredFlexBox } from '@/components/styled';
import { auth } from '@/core/firebase';
import useIdToken from '@/hooks/useIdToken';
import routes from '@/routes';
import { Pages } from '@/routes/types';

function LogInPage() {
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
          <LogIn />
        </Box>
      </FullSizeCenteredFlexBox>
    </>
  );
}

export default LogInPage;
