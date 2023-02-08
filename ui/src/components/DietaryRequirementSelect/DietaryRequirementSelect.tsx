import { useRecoilValue, useSetRecoilState } from 'recoil';

import { FormControl, InputLabel, MenuItem, Select, SelectChangeEvent } from '@mui/material';
import { Box } from '@mui/system';

import { DietaryRequirementsState, LocalUserSettingsState } from '@/store/user';
import { MIN_WIDTH } from '@/config';
import DietaryRequirement from '@/core/dietary_requirements';

function DietaryRequirementSelect() {
  const localUserSettings = useRecoilValue(LocalUserSettingsState);
  const setLocalUserSettings = useSetRecoilState(LocalUserSettingsState);
  const dietaryRequirements = useRecoilValue(DietaryRequirementsState);

  const handleChange = (event: SelectChangeEvent) => {
    setLocalUserSettings((old) => {
      return {
        ...old,
        dietaryRequirement: new DietaryRequirement(parseInt(event.target.value)),
        dirty: true,
      };
    });
  };

  return (
    <Box sx={{ minWidth: MIN_WIDTH }}>
      <FormControl fullWidth>
        <InputLabel id="dietary-requirement-select-label">Dietary Requirement</InputLabel>
        <Select
          labelId="dietary-requirement-select-label"
          id="dietary-requirement-select"
          label="dietary-requirement"
          value={localUserSettings.dietaryRequirement.value.toString()}
          onChange={handleChange}
        >
          {dietaryRequirements.map((dr, idx) => (
            <MenuItem key={idx} value={dr.value}>
              {dr.getSettingName()}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    </Box>
  );
}

export default DietaryRequirementSelect;
