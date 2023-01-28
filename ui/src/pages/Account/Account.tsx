import { Link } from 'react-router-dom';
import { useRecoilValue, useSetRecoilState } from 'recoil';

import AccountCircle from '@mui/icons-material/AccountCircle';
import { Button, Chip, Divider, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { Box } from '@mui/system';

import DeleteAccount from '@/components/DeleteAccount';
import Loading from '@/components/Loading';
import RegionSelect from '@/components/RegionSelect';
import { CarbonFootprintSlider } from '@/components/StyledSlider';
import { useCurrentUser } from '@/hooks/useCurrentUser';
import { useSignOut } from '@/hooks/useSignOut';
import { LocalUserSettingsState } from '@/store/user';

function Account() {
  const user = useCurrentUser();
  const localUserSettings = useRecoilValue(LocalUserSettingsState);
  const setLocalUserSettings = useSetRecoilState(LocalUserSettingsState);
  const signOut = useSignOut();
  const saveSettings = () => {
    // TODO: sync userSettings
    setLocalUserSettings((old) => {
      return { ...old, dirty: false };
    });
  };

  if (user === undefined) {
    return (
      <>
        <Grid container justifyContent="center" alignItems="center" height="100%">
          <Loading />
        </Grid>
      </>
    );
  } else if (user === null) {
    return (
      <Grid container columns={16} justifyContent="center" alignItems="center" height="100%">
        <Grid xs={2} display="flex" justifyContent="center" alignItems="center">
          <Button component={Link} to="/login" variant="contained">
            Log in
          </Button>
        </Grid>
        <Grid xs={2} display="flex" justifyContent="center" alignItems="center">
          <Button component={Link} to="/signup" variant="contained">
            Sign up
          </Button>
        </Grid>
      </Grid>
    );
  } else {
    return (
      <Box>
        <Grid
          container
          columns={10}
          spacing={1}
          justifyContent="center"
          alignItems="center"
          pt="4vh"
          disableEqualOverflow
        >
          <Grid>
            <AccountCircle fontSize="large" />
          </Grid>
          <Grid>
            <Typography variant="h4">{user.displayName}</Typography>
          </Grid>
        </Grid>
        <Divider sx={{ py: '4vh' }}>
          <Chip label="Settings" />
        </Divider>
        <Grid
          container
          direction="column"
          columns={10}
          spacing={4}
          alignItems="center"
          disableEqualOverflow
        >
          <Grid xs={8} sm={7} md={6} lg={5} xl={4}>
            <RegionSelect />
          </Grid>
          <Grid xs={8} sm={7} md={6} lg={5} xl={4}>
            <CarbonFootprintSlider />
          </Grid>
          <Grid textAlign="center" xs={8} sm={7} md={6} lg={5} xl={4}>
            <Button variant="contained" disabled={!localUserSettings.dirty} onClick={saveSettings}>
              Save
            </Button>
          </Grid>
          <Grid textAlign="center" xs={8} sm={7} md={6} lg={5} xl={4}>
            <Button variant="contained" onClick={signOut}>
              Log out
            </Button>
          </Grid>
          <Grid textAlign="center" xs={8} sm={7} md={6} lg={5} xl={4}>
            <DeleteAccount user={user} />
          </Grid>
        </Grid>
      </Box>
    );
  }
}

export default Account;
