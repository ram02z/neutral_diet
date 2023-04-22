import { BaseSyntheticEvent, useMemo, useState } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { useRecoilValue } from 'recoil';

import { Tune } from '@mui/icons-material';
import {
  Badge,
  Button,
  Chip,
  Collapse,
  FormControl,
  InputLabel,
  Link,
  MenuItem,
  OutlinedInput,
  Select,
  Stack,
  TextField,
} from '@mui/material';
import { Box } from '@mui/system';

import { SubTypologiesState, TypologiesState } from '@/store/food';
import { SearchSortMethodsState } from '@/store/search';
import { SearchFilters, SearchSortMethod } from '@/store/search/types';

import { FormValues } from './types';

const ITEM_HEIGHT = 48;
const ITEM_PADDING_TOP = 8;
const MENU_WIDTH = 330;
const MenuProps = {
  PaperProps: {
    style: {
      maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
      width: MENU_WIDTH,
    },
  },
};

type SortFilterMenuProps = {
  onSubmit: (data: FormValues, event?: BaseSyntheticEvent<object, void, void | undefined>) => void;
  currentSearchFilters: SearchFilters;
  currentSortingMethod: SearchSortMethod;
};

/**
 * Menu with sort and filter settings. Multi-select is enabled for filters.
 */
function SortFilterMenu({
  onSubmit,
  currentSearchFilters,
  currentSortingMethod,
}: SortFilterMenuProps) {
  const { handleSubmit, control, reset } = useForm<FormValues>();
  const [expanded, setExpanded] = useState(false);
  const searchSortMethods = useRecoilValue(SearchSortMethodsState);
  const typologyNames = useRecoilValue(TypologiesState);
  const subTypologyNames = useRecoilValue(SubTypologiesState);
  const noSearchFilters = useMemo(() => {
    return currentSearchFilters.typologies.length + currentSearchFilters.subTypologies.length;
  }, [currentSearchFilters]);

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  const resetForm = () => {
    reset({ typologyNames: [], subTypologyNames: [] });
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Stack direction="column" alignItems="center" spacing={2}>
        <Badge color="info" overlap="circular" badgeContent={noSearchFilters}>
          <Button variant="contained" startIcon={<Tune />} onClick={handleExpandClick}>
            Sort and filter
          </Button>
        </Badge>
        <Collapse in={expanded} timeout={0} unmountOnExit>
          <Stack direction="column" alignItems="center" spacing={2} sx={{ py: 2 }}>
            <Link sx={{ marginLeft: 'auto' }} href="#" onClick={resetForm}>
              Clear All
            </Link>
            <Stack spacing={2} direction={{ xs: 'column', xl: 'row' }}>
              <Controller
                control={control}
                name="sortingMethod"
                defaultValue={currentSortingMethod}
                render={({ field: { onChange, value } }) => (
                  <TextField
                    select
                    label="Sort by"
                    onChange={onChange}
                    value={value}
                    sx={{ minWidth: MENU_WIDTH }}
                  >
                    {searchSortMethods.map((method, key) => (
                      <MenuItem key={key} value={method.value}>
                        {method.getName()}
                      </MenuItem>
                    ))}
                  </TextField>
                )}
              />
              <FormControl sx={{ minWidth: MENU_WIDTH }}>
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
              <FormControl sx={{ minWidth: MENU_WIDTH }}>
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
            </Stack>
            <Button variant="contained" color="secondary" type="submit">
              Apply
            </Button>
          </Stack>
        </Collapse>
      </Stack>
    </form>
  );
}

export default SortFilterMenu;
