import { Link } from 'react-router-dom';

import { Button, Typography } from '@mui/material';
import { Box } from '@mui/system';

import DeleteAccount from '@/components/DeleteAccount';
import Loading from '@/components/Loading';
import RegionSelect from '@/components/RegionSelect';
import { FullSizeCenteredFlexBox } from '@/components/styled';
import { useCurrentUser } from '@/hooks/useCurrentUser';
import { useSignOut } from '@/hooks/useSignOut';

function Account() {
  const user = useCurrentUser();
  const signOut = useSignOut();

  if (user === undefined) {
    return (
      <>
        <Loading />
      </>
    );
  } else if (user === null) {
    return (
      <FullSizeCenteredFlexBox>
        <Box m="auto" sx={{ textAlign: 'center' }}>
          <Button component={Link} to="/login" variant="contained">
            Log in
          </Button>
          <Button component={Link} to="/signup" variant="contained">
            Sign up
          </Button>
        </Box>
      </FullSizeCenteredFlexBox>
    );
  } else {
    return (
      <div>
        <Typography>Current User: {user.displayName}</Typography>
        <RegionSelect user={user}></RegionSelect>
        <br />
        <Button onClick={signOut}>Log out</Button>
        <DeleteAccount user={user}></DeleteAccount>
      </div>
    );
  }
}

export default Account;
