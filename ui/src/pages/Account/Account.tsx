import { useEffect } from 'react';
import { Link } from 'react-router-dom';
import { useRecoilState, useRecoilValue } from 'recoil';

import { Avatar, Button, Chip, Divider, Stack } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import { useSnackbar } from 'notistack';

import client from '@/api/user_service';
import DeleteAccount from '@/components/DeleteAccount';
import DietaryRequirementSelect from '@/components/DietaryRequirementSelect';
import Loading from '@/components/Loading';
import RegionSelect from '@/components/RegionSelect';
import { CarbonFootprintSlider } from '@/components/StyledSlider';
import { useCurrentUser } from '@/hooks/useCurrentUser';
import { useSignOut } from '@/hooks/useSignOut';
import {
  CurrentUserHeadersState,
  LocalUserSettingsState,
  RemoteUserSettingsState,
} from '@/store/user';

function Account() {
  const user = useCurrentUser();
  const userHeaders = useRecoilValue(CurrentUserHeadersState);
  const remoteUserSettings = useRecoilValue(RemoteUserSettingsState);
  const [localUserSettings, setLocalUserSettings] = useRecoilState(LocalUserSettingsState);
  const signOut = useSignOut();
  const { enqueueSnackbar } = useSnackbar();
  const saveSettings = () => {
    client
      .updateUserSettings({ userSettings: remoteUserSettings }, { headers: userHeaders })
      .then(() =>
        setLocalUserSettings((old) => {
          return { ...old, dirty: false };
        }),
      )
      .then(() => enqueueSnackbar('Updated account settings.', { variant: 'success' }))
      .catch((err) => {
        enqueueSnackbar('Could not save account settings.', { variant: 'error' });
        console.error(err);
      });
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
      <Stack direction="row" justifyContent="center" alignItems="center" height="100%" spacing={2}>
        <Button component={Link} to="/login" variant="contained">
          Log in
        </Button>
        <Button component={Link} to="/signup" variant="contained">
          Sign up
        </Button>
      </Stack>
    );
  } else {
    return (
      <Grid>
        <Grid
          container
          columns={10}
          spacing={1}
          justifyContent="center"
          alignItems="center"
          pt="4vh"
          disableEqualOverflow
          direction="column"
        >
          <Grid>
            <Avatar sx={{ width: 80, height: 80, fontSize: 40 }}>{user.displayName?.at(0)}</Avatar>
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
          <Grid xs={8} lg={7} xl={6}>
            <RegionSelect />
          </Grid>
          <Grid xs={8} lg={7} xl={6}>
            <DietaryRequirementSelect />
          </Grid>
          <Grid xs={8} lg={7} xl={6}>
            <CarbonFootprintSlider />
          </Grid>
          <Grid textAlign="center" xs={8} lg={7} xl={6}>
            <Button variant="contained" disabled={!localUserSettings.dirty} onClick={saveSettings}>
              Save
            </Button>
          </Grid>
          <Grid textAlign="center" xs={8} lg={7} xl={6}>
            <Button variant="contained" onClick={signOut}>
              Log out
            </Button>
          </Grid>
          <Grid textAlign="center" xs={8} lg={7} xl={6}>
            <DeleteAccount user={user} />
          </Grid>
        </Grid>
      </Grid>
    );
  }
}

export default Account;
