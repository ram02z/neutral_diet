import { useEffect } from 'react';
import { Link } from 'react-router-dom';
import { useRecoilState, useRecoilValue } from 'recoil';

import AccountCircle from '@mui/icons-material/AccountCircle';
import { Button, Chip, Divider, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { Box } from '@mui/system';

import { useSnackbar } from 'notistack';

import { ID_TOKEN_HEADER } from '@/api/transport';
import client from '@/api/user_service';
import DeleteAccount from '@/components/DeleteAccount';
import Loading from '@/components/Loading';
import RegionSelect from '@/components/RegionSelect';
import { CarbonFootprintSlider } from '@/components/StyledSlider';
import { useCurrentUser } from '@/hooks/useCurrentUser';
import { useSignOut } from '@/hooks/useSignOut';
import {
  CurrentUserTokenIDState,
  LocalUserSettingsState,
  RemoteUserSettingsState,
} from '@/store/user';
import DietaryRequirementSelect from '@/components/DietaryRequirementSelect';

function Account() {
  const user = useCurrentUser();
  const idToken = useRecoilValue(CurrentUserTokenIDState);
  const remoteUserSettings = useRecoilValue(RemoteUserSettingsState);
  const [localUserSettings, setLocalUserSettings] = useRecoilState(LocalUserSettingsState);
  const signOut = useSignOut();
  const { enqueueSnackbar } = useSnackbar();
  const saveSettings = () => {
    setLocalUserSettings((old) => {
      return { ...old, dirty: false };
    });
    if (idToken) {
      const headers = new Headers();
      headers.set(ID_TOKEN_HEADER, idToken);
      client
        .updateUserSettings({ userSettings: remoteUserSettings }, { headers: headers })
        .then(() => enqueueSnackbar('Updated account settings.', { variant: 'success' }))
        .catch((err) => {
          enqueueSnackbar('Could not save account settings.', { variant: 'error' });
          console.error(err);
        });
    }
  };

  useEffect(() => {
    if (localUserSettings.dirty === true) {
      enqueueSnackbar('Unsaved changes to settings.', { variant: 'warning' });
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

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
            <DietaryRequirementSelect />
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
