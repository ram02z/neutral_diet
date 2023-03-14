import { useRecoilValue, useSetRecoilState } from 'recoil';

import { FormControl, InputLabel, MenuItem, Select, SelectChangeEvent } from '@mui/material';
import { Box } from '@mui/system';

import { MIN_CARD_WIDTH } from '@/config';
import { DietaryRequirementsState, LocalUserSettingsState } from '@/store/user';

function DietaryRequirementSelect() {
  const localUserSettings = useRecoilValue(LocalUserSettingsState);
  const setLocalUserSettings = useSetRecoilState(LocalUserSettingsState);
  const dietaryRequirements = useRecoilValue(DietaryRequirementsState);

  const handleChange = (event: SelectChangeEvent) => {
    setLocalUserSettings((old) => {
      return {
        ...old,
        dietaryRequirement: parseInt(event.target.value),
        dirty: true,
      };
    });
  };

  return (
    <FormControl sx={{ minWidth: MIN_CARD_WIDTH }} fullWidth>
      <InputLabel id="dietary-requirement-select-label">Dietary Requirement</InputLabel>
      <Select
        labelId="dietary-requirement-select-label"
        id="dietary-requirement-select"
        label="dietary-requirement"
        value={localUserSettings.dietaryRequirement.toString()}
        onChange={handleChange}
      >
        {dietaryRequirements.map((dr, idx) => (
          <MenuItem key={idx} value={dr.value}>
            {dr.getSettingName()}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
}

export default DietaryRequirementSelect;
