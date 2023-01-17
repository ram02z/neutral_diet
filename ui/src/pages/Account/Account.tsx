import { Button } from '@mui/material';
import { Box } from '@mui/system';

import { signOut } from 'firebase/auth';

import Loading from '@/components/Loading';
import { FullSizeCenteredFlexBox } from '@/components/styled';
import { auth } from '@/core/firebase';
import useIdToken from '@/hooks/useIdToken';
import { Link } from 'react-router-dom';

const logout = () => {
  signOut(auth);
};

function Account() {
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
    return (
      <div>
        <p>Current User: {user.displayName}</p>
        <button onClick={logout}>Log out</button>
      </div>
    );
  }

  return (
    <FullSizeCenteredFlexBox>
      <Box m="auto" sx={{ textAlign: 'center' }}>
        <Button component={Link} to="/login" variant="contained">Log in</Button>
        <Button component={Link} to="/signup" variant="contained">Sign up</Button>
      </Box>
    </FullSizeCenteredFlexBox>
  );
}

export default Account;
