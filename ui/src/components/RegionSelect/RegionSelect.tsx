import { useEffect, useState } from 'react';

import {
  Button,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  SelectChangeEvent,
} from '@mui/material';
import { Box } from '@mui/system';

import { User } from 'firebase/auth';

import * as food_service from '@/api/food_service';
import * as user_service from '@/api/user_service';
import { Region } from '@/api/gen/neutral_diet/food/v1/region_pb';
import { ID_TOKEN_HEADER } from '@/api/transport';
import { useCurrentUserSettings } from '@/hooks/useCurrentUserSettings';

import Loading from '../Loading';

type RegionSelectProps = {
  user: User;
};

function updateUserRegion(user: User, regionName: string) {
  user.getIdToken().then((idToken) => {
    const headers = new Headers();
    headers.set(ID_TOKEN_HEADER, idToken);
    user_service.default.updateUserRegion(
      { region: new Region({ name: regionName }) },
      { headers: headers },
    );
  });
}

function RegionSelect({ user }: RegionSelectProps) {
  const userSettings = useCurrentUserSettings();
  const [regions, setRegions] = useState<Region[]>([]);
  const [regionName, setRegionName] = useState('');

  const handleChange = (event: SelectChangeEvent) => {
    setRegionName(event.target.value as string);
  };
  useEffect(() => {
    food_service.default
      .listRegions({})
      .then((response) => setRegions(response.regions))
      .catch((e) => console.error(e.message));
  }, []);

  useEffect(() => {
    if (userSettings && userSettings.region) setRegionName(userSettings.region?.name);
  }, [userSettings]);

  if (userSettings === undefined || userSettings === null) {
    return (
      <>
        <Loading />
      </>
    );
  } else {
    return (
      <>
        <Box sx={{ minWidth: 120 }}>
          <FormControl fullWidth>
            <InputLabel id="region-select-label">Region</InputLabel>
            <Select
              labelId="region-select-label"
              id="region-select"
              label="Region"
              defaultValue={regionName}
              value={regionName}
              onChange={handleChange}
            >
              {regions.map((region, idx) => (
                <MenuItem key={idx} value={region.name}>
                  {region.name}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <Button onClick={() => updateUserRegion(user, regionName)}>Save</Button>
        </Box>
      </>
    );
  }
}

export default RegionSelect;
