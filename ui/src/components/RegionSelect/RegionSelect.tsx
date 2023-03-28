import { useRecoilValue, useSetRecoilState } from 'recoil';

import { FormControl, InputLabel, MenuItem, Select, SelectChangeEvent } from '@mui/material';

import { MIN_CARD_WIDTH } from '@/config';
import { RegionsState } from '@/store/food';
import { LocalUserSettingsState } from '@/store/user';

function RegionSelect() {
  const localUserSettings = useRecoilValue(LocalUserSettingsState);
  const setLocalUserSettings = useSetRecoilState(LocalUserSettingsState);
  const regions = useRecoilValue(RegionsState);

  const handleChange = (event: SelectChangeEvent) => {
    setLocalUserSettings((old) => {
      return {
        ...old,
        region: parseInt(event.target.value),
        dirty: true,
      };
    });
  };

  return (
    <FormControl sx={{ minWidth: MIN_CARD_WIDTH }} fullWidth>
      <InputLabel id="region-select-label">Region</InputLabel>
      <Select
        labelId="region-select-label"
        id="region-select"
        label="Region"
        value={localUserSettings.region.toString()}
        onChange={handleChange}
      >
        {regions.map((region, idx) => (
          <MenuItem key={idx} value={region.value}>
            {region.getSettingName()}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
}

export default RegionSelect;
