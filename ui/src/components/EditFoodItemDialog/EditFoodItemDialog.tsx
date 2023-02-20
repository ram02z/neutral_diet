import { BaseSyntheticEvent } from 'react';
import { Controller, useForm } from 'react-hook-form';

import { Button, Dialog, DialogActions, DialogTitle, MenuItem, Select, TextField } from '@mui/material';
import { Stack } from '@mui/system';

import { FormValues } from '@/components/FoodItemCard/types';
import { Weight, WeightUnitNameMap } from '@/core/weight';

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
            defaultValue={currentWeight.getWeightUnitName()}
            rules={{ required: true }}
            render={({ field: { onChange, value }, fieldState: { error } }) => (
              <TextField
                select
                label="Unit"
                error={!!error}
                onChange={onChange}
                value={value}
              >
                {[...WeightUnitNameMap.values()].map((value, key) => (
                  <MenuItem key={key} value={value}>
                    {value}
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
