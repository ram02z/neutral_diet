import { BaseSyntheticEvent } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { useRecoilValue } from 'recoil';

import { Button, Dialog, DialogActions, DialogTitle, MenuItem, TextField } from '@mui/material';
import { Stack } from '@mui/system';
import { DatePicker } from '@mui/x-date-pickers';

import dayjs from 'dayjs';

import { FormValues } from '@/components/FoodItemCard/types';
import { WeightUnitsState } from '@/store/user';

type AddFoodItemDialogProps = {
  onSubmit: (data: FormValues, event?: BaseSyntheticEvent<object, void, void | undefined>) => void;
  openDialog: boolean;
  handleClose: () => void;
};

function AddFoodItemDialog({ openDialog, handleClose, onSubmit }: AddFoodItemDialogProps) {
  const { handleSubmit, control } = useForm<FormValues>();
  const weightUnits = useRecoilValue(WeightUnitsState);
  return (
    <Dialog open={openDialog} onClose={handleClose}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <DialogTitle textAlign="center">Add food</DialogTitle>
        <Stack sx={{ px: 5 }} spacing={3}>
          <Controller
            control={control}
            name="date"
            defaultValue={dayjs()}
            rules={{
              required: true,
            }}
            render={({ field: { ref, onChange, value }, fieldState: { error } }) => (
              <DatePicker
                value={value}
                onChange={onChange}
                inputRef={ref}
                label="Date"
                inputFormat="YYYY-MM-DD"
                maxDate={dayjs()}
                renderInput={(inputProps) => (
                  <TextField {...inputProps} error={!!error} helperText={error?.message} />
                )}
              />
            )}
          />
          <Controller
            control={control}
            name="weight"
            defaultValue="1.0"
            rules={{ required: true, min: 0.001 }}
            render={({ field: { onChange, value }, fieldState: { error } }) => (
              <TextField
                error={!!error}
                helperText={error?.message}
                onChange={onChange}
                value={value}
                type="number"
                label="Weight"
                inputProps={{ step: 0.001 }}
              />
            )}
          />
          <Controller
            control={control}
            name="weightUnit"
            rules={{ required: true }}
            render={({ field: { onChange, value }, fieldState: { error } }) => (
              <TextField select label="Unit" error={!!error} onChange={onChange} value={value}>
                {weightUnits.map((weight, key) => (
                  <MenuItem key={key} value={weight.value}>
                    {weight.getName()}
                  </MenuItem>
                ))}
              </TextField>
            )}
          />
        </Stack>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button variant="contained" type="submit">
            Add
          </Button>
        </DialogActions>
      </form>
    </Dialog>
  );
}

export default AddFoodItemDialog;
