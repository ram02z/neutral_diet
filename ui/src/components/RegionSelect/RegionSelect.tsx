import { useEffect, useState } from 'react';

import { FormControl, InputLabel, MenuItem, Select, SelectChangeEvent } from '@mui/material';
import { Box } from '@mui/system';

import client from '@/api/food_service';
import { Region } from '@/api/gen/neutral_diet/food/v1/region_pb';

function RegionSelect() {
  const [regions, setRegions] = useState<Region[]>([]);
  const [regionName, setRegionName] = useState('');

  const handleChange = (event: SelectChangeEvent) => {
    setRegionName(event.target.value as string);
  };
  useEffect(() => {
    client
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
      </Box>
    </>
  );
}

export default RegionSelect;
