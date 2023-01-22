import { Navigate } from 'react-router-dom';

import { Box } from '@mui/system';

import Loading from '@/components/Loading';
import LogIn from '@/components/LogIn';
import { FullSizeCenteredFlexBox } from '@/components/styled';
import { useCurrentUser } from '@/hooks/useCurrentUser';
import routes from '@/routes';
import { Pages } from '@/routes/types';

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
      <>
        <FullSizeCenteredFlexBox>
          <Box>
            <LogIn />
          </Box>
        </FullSizeCenteredFlexBox>
      </>
    );
  } else {
    return <Navigate to={routes[Pages.Account].path} />;
  }
}

export default LogInPage;
