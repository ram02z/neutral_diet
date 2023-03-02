import { BaseSyntheticEvent } from 'react';
import { Controller, useForm } from 'react-hook-form';

import { Button, Dialog, DialogActions, DialogTitle, MenuItem, TextField } from '@mui/material';
import { Stack } from '@mui/system';

import { FormValues } from '@/components/FoodItemCard/types';
import { Weight } from '@/core/weight';
import { useRecoilValue } from 'recoil';
import { WeightUnitsState } from '@/store/user';

type EditFoodItemDialogProps = {
  onSubmit: (data: FormValues, event?: BaseSyntheticEvent<object, void, void | undefined>) => void;
  openDialog: boolean;
  handleClose: () => void;
  currentWeight: Weight;
};

function EditFoodItemDialog({
  openDialog,
  currentWeight,
  handleClose,
  onSubmit,
}: EditFoodItemDialogProps) {
  const { handleSubmit, control } = useForm<FormValues>();
  const weightUnits = useRecoilValue(WeightUnitsState);
  return (
    <Dialog open={openDialog} onClose={handleClose}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <DialogTitle textAlign="center">Edit food</DialogTitle>
        <Stack sx={{ px: 5 }} spacing={3}>
          <Controller
            control={control}
            name="weight"
            defaultValue={currentWeight.value.toString()}
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
            defaultValue={currentWeight.unit.value}
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
            Update
          </Button>
        </DialogActions>
      </form>
    </Dialog>
  );
}

export default EditFoodItemDialog;
