import { BaseSyntheticEvent, useState } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { useRecoilState, useRecoilValue } from 'recoil';

import { Tune } from '@mui/icons-material';
import {
  Button,
  Chip,
  Collapse,
  FormControl,
  InputLabel,
  MenuItem,
  OutlinedInput,
  Select,
  Stack,
} from '@mui/material';
import { Box } from '@mui/system';

import { SubTypologiesState, TypologiesState } from '@/store/food';

import { FormValues } from './types';
import { SearchFilters } from '@/store/search/types';

const ITEM_HEIGHT = 48;
const ITEM_PADDING_TOP = 8;
const MenuProps = {
  PaperProps: {
    style: {
      maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
      width: 250,
    },
  },
};

type SortFilterMenuProps = {
  onSubmit: (data: FormValues, event?: BaseSyntheticEvent<object, void, void | undefined>) => void;
  currentSearchFilters: SearchFilters;
}

function SortFilterMenu({ onSubmit, currentSearchFilters }: SortFilterMenuProps) {
  const { handleSubmit, control } = useForm<FormValues>();
  const [expanded, setExpanded] = useState(false);
  const typologyNames = useRecoilValue(TypologiesState);
  const subTypologyNames = useRecoilValue(SubTypologiesState);

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };


  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Stack direction="column" alignItems="center" spacing={2}>
        <Button variant="contained" startIcon={<Tune />} onClick={handleExpandClick}>
          Sort and filter
        </Button>
        <Collapse in={expanded} timeout={0} unmountOnExit>
          <Stack direction="column" alignItems="center" spacing={2} sx={{ py: 2 }}>
            <FormControl sx={{ m: 1, width: 300 }}>
              <InputLabel id="typology-select-label">Typology</InputLabel>
              <Controller
                control={control}
                name="typologyNames"
                defaultValue={currentSearchFilters.typologies}
                render={({ field: { ref, onChange, value } }) => (
                  <Select
                    multiple
                    fullWidth
                    labelId="typology-select-label"
                    inputRef={ref}
                    value={value}
                    onChange={onChange}
                    input={<OutlinedInput id="select-multiple-chip" label="Chip" />}
                    renderValue={(selected) => (
                      <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                        {selected.map((value) => (
                          <Chip key={value} label={value} />
                        ))}
                      </Box>
                    )}
                    MenuProps={MenuProps}
                  >
                    {typologyNames.map((name, idx) => (
                      <MenuItem key={idx} value={name}>
                        {name}
                      </MenuItem>
                    ))}
                  </Select>
                )}
              />
            </FormControl>
            <FormControl sx={{ m: 1, width: 300 }}>
              <InputLabel id="sub-typology-select-label">Sub-Typology</InputLabel>
              <Controller
                control={control}
                name="subTypologyNames"
                defaultValue={currentSearchFilters.subTypologies}
                render={({ field: { ref, onChange, value } }) => (
                  <Select
                    multiple
                    fullWidth
                    labelId="sub-typology-select-label"
                    inputRef={ref}
                    value={value}
                    onChange={onChange}
                    input={<OutlinedInput id="select-multiple-chip" label="Chip" />}
                    renderValue={(selected) => (
                      <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                        {selected.map((value) => (
                          <Chip key={value} label={value} />
                        ))}
                      </Box>
                    )}
                    MenuProps={MenuProps}
                  >
                    {subTypologyNames.map((name, idx) => (
                      <MenuItem key={idx} value={name}>
                        {name}
                      </MenuItem>
                    ))}
                  </Select>
                )}
              />
            </FormControl>
            <Button variant="contained" type="submit">
              Apply
            </Button>
          </Stack>
        </Collapse>
      </Stack>
    </form>
  );
}

export default SortFilterMenu;
