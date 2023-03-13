import { BaseSyntheticEvent } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { useRecoilValue } from 'recoil';

import { Button, Dialog, DialogActions, DialogTitle, MenuItem, TextField } from '@mui/material';
import { Stack } from '@mui/system';

import { FormValues } from '@/components/FoodItemLogCard/types';
import { QuantityUnit } from '@/core/food_unit';
import { FoodUnitsState, MealsState } from '@/store/user';

type EditFoodItemDialogProps = {
  onSubmit: (data: FormValues, event?: BaseSyntheticEvent<object, void, void | undefined>) => void;
  openDialog: boolean;
  handleClose: () => void;
  currentQuantity: QuantityUnit;
  currentMeal: number;
};

function EditFoodItemDialog({
  openDialog,
  currentQuantity,
  currentMeal,
  handleClose,
  onSubmit,
}: EditFoodItemDialogProps) {
  const { handleSubmit, control } = useForm<FormValues>();
  const foodUnits = useRecoilValue(FoodUnitsState);
  const meals = useRecoilValue(MealsState);

  return (
    <Dialog open={openDialog} onClose={handleClose}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <DialogTitle textAlign="center">Edit food</DialogTitle>
        <Stack sx={{ px: 5 }} spacing={3}>
          <Controller
            control={control}
            name="meal"
            defaultValue={currentMeal}
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
            defaultValue={currentQuantity.value.toString()}
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
            defaultValue={currentQuantity.unit.value}
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
          <Button type="submit">Update</Button>
        </DialogActions>
      </form>
    </Dialog>
  );
}

export default EditFoodItemDialog;
