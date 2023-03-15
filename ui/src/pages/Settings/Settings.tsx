import { useEffect } from 'react';
import { useRecoilState, useRecoilValue } from 'recoil';

import { Button, Stack } from '@mui/material';

import { useSnackbar } from 'notistack';

import client from '@/api/user_service';
import DietaryRequirementSelect from '@/components/DietaryRequirementSelect';
import RegionSelect from '@/components/RegionSelect';
import { CarbonFootprintSlider } from '@/components/StyledSlider';
import {
  CurrentUserHeadersState,
  LocalUserSettingsState,
  RemoteUserSettingsState,
} from '@/store/user';

function Settings() {
  const userHeaders = useRecoilValue(CurrentUserHeadersState);
  const remoteUserSettings = useRecoilValue(RemoteUserSettingsState);
  const [localUserSettings, setLocalUserSettings] = useRecoilState(LocalUserSettingsState);
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

  return (
    <Stack
      direction="column"
      spacing={4}
      alignItems="center"
      justifyContent="center"
      width={{ xs: '60vw', md: '40vw' }}
    >
      <RegionSelect />
      <DietaryRequirementSelect />
      <CarbonFootprintSlider />
      <Button variant="contained" disabled={!localUserSettings.dirty} onClick={saveSettings}>
        Save
      </Button>
    </Stack>
  );
}

export default Settings;
