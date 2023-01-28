import { useEffect, useState } from 'react';
import { useRecoilValue, useSetRecoilState } from 'recoil';

import { FormControl, InputLabel, MenuItem, Select, SelectChangeEvent } from '@mui/material';
import { Box } from '@mui/system';

import * as food_service from '@/api/food_service';
import { Region } from '@/api/gen/neutral_diet/food/v1/region_pb';
import { LocalUserSettingsState } from '@/store/user';

function RegionSelect() {
  const localUserSettings = useRecoilValue(LocalUserSettingsState);
  const setLocalUserSettings = useSetRecoilState(LocalUserSettingsState);
  const [regions, setRegions] = useState<Region[]>([]);

  const handleChange = (event: SelectChangeEvent) => {
    setLocalUserSettings((old) => {
      return {
        ...old,
        region: event.target.value as string,
        dirty: true,
      };
    });
  };
  useEffect(() => {
    food_service.default
      .listRegions({})
      .then((response) => setRegions(response.regions))
      .catch((e) => console.error(e.message));
  }, []);

  return (
    <>
      <Box sx={{ minWidth: 120 }}>
        <FormControl fullWidth>
          <InputLabel id="region-select-label">Region</InputLabel>
          <Select
            labelId="region-select-label"
            id="region-select"
            label="Region"
            value={localUserSettings.region}
            onChange={handleChange}
          >
            {regions.map((region, idx) => (
              <MenuItem key={idx} value={region.name}>
                {region.name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      </Box>
    </>
  );
}

export default RegionSelect;
