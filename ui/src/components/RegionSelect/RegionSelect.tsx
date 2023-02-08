import { useRecoilValue, useSetRecoilState } from 'recoil';

import { FormControl, InputLabel, MenuItem, Select, SelectChangeEvent } from '@mui/material';
import { Box } from '@mui/system';

import { RegionsState } from '@/store/food';
import { LocalUserSettingsState } from '@/store/user';
import { MIN_WIDTH } from '@/config';

function RegionSelect() {
  const localUserSettings = useRecoilValue(LocalUserSettingsState);
  const setLocalUserSettings = useSetRecoilState(LocalUserSettingsState);
  // TODO: handle errors
  const regions = useRecoilValue(RegionsState);

  const handleChange = (event: SelectChangeEvent) => {
    setLocalUserSettings((old) => {
      return {
        ...old,
        region: event.target.value as string,
        dirty: true,
      };
    });
  };

  return (
    <Box sx={{ minWidth: MIN_WIDTH }}>
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
  );
}

export default RegionSelect;
