import { BaseSyntheticEvent } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { useRecoilValue } from 'recoil';

import { Button, Dialog, DialogActions, DialogTitle, MenuItem, TextField } from '@mui/material';
import { Stack } from '@mui/system';
import { DatePicker } from '@mui/x-date-pickers';

import dayjs from 'dayjs';

import { FormValues } from '@/components/FoodItemCard/types';
import { FoodUnitsState, MealsState } from '@/store/user';

type AddFoodItemDialogProps = {
  onSubmit: (data: FormValues, event?: BaseSyntheticEvent<object, void, void | undefined>) => void;
  openDialog: boolean;
  handleClose: () => void;
};

function AddFoodItemDialog({ openDialog, handleClose, onSubmit }: AddFoodItemDialogProps) {
  const { handleSubmit, control } = useForm<FormValues>();
  const foodUnits = useRecoilValue(FoodUnitsState);
  const meals = useRecoilValue(MealsState);

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
              validate: (value) => !value.isAfter(dayjs()),
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
            name="meal"
            rules={{ required: true }}
            render={({ field: { onChange, value }, fieldState: { error } }) => (
              <TextField select label="Meal" error={!!error} onChange={onChange} value={value}>
                {meals.map((meal, key) => (
                  <MenuItem key={key} value={meal.value}>
                    {meal.getName()}
                  </MenuItem>
                ))}
              </TextField>
            )}
          />
          <Controller
            control={control}
            name="quantity"
            defaultValue="1.0"
            rules={{
              required: true,
              min: { value: 0.001, message: 'Quantity needs to be greater than 0.001' },
            }}
            render={({ field: { onChange, value }, fieldState: { error } }) => (
              <TextField
                error={!!error}
                helperText={error?.message}
                onChange={onChange}
                value={value}
                type="number"
                label="Quantity"
                inputProps={{ step: 0.001 }}
              />
            )}
          />
          <Controller
            control={control}
            name="unit"
            rules={{ required: true }}
            render={({ field: { onChange, value }, fieldState: { error } }) => (
              <TextField select label="Unit" error={!!error} onChange={onChange} value={value}>
                {foodUnits.map((unit, key) => (
                  <MenuItem key={key} value={unit.value}>
                    {unit.getName()}
                  </MenuItem>
                ))}
              </TextField>
            )}
          />
        </Stack>
        <DialogActions sx={{ mt: '24px' }}>
          <Button color="secondary" onClick={handleClose}>
            Cancel
          </Button>
          <Button type="submit">Add</Button>
        </DialogActions>
      </form>
    </Dialog>
  );
}

export default AddFoodItemDialog;
