import { BaseSyntheticEvent } from 'react';
import { Controller, useForm } from 'react-hook-form';

import { Button, Dialog, DialogActions, DialogTitle, TextField } from '@mui/material';
import { Stack } from '@mui/system';
import { DatePicker } from '@mui/x-date-pickers';

import dayjs from 'dayjs';

import { MIN_CF_LIMIT } from '@/config';

import { FormValues } from './types';

type AddGoalDialogProps = {
  onSubmit: (data: FormValues, event?: BaseSyntheticEvent<object, void, void | undefined>) => void;
  openDialog: boolean;
  handleClose: () => void;
  startCarbonFootprint: number;
};

function AddGoalDialog({
  onSubmit,
  openDialog,
  handleClose,
  startCarbonFootprint,
}: AddGoalDialogProps) {
  const { handleSubmit, control } = useForm<FormValues>();

  // TODO: use CarbonFootprintSlider
  return (
    <Dialog fullWidth open={openDialog} onClose={handleClose}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <DialogTitle textAlign="center">Add Carbon footprint Goal</DialogTitle>
        <Stack sx={{ px: 5 }} spacing={3}>
          <Controller
            control={control}
            name="description"
            rules={{ required: true }}
            render={({ field: { onChange, value }, fieldState: { error } }) => (
              <TextField
                label="Description"
                multiline
                error={!!error}
                onChange={onChange}
                value={value}
              ></TextField>
            )}
          />
          <Controller
            control={control}
            name="targetCarbonFootprint"
            rules={{ required: true, min: MIN_CF_LIMIT, max: startCarbonFootprint }}
            render={({ field: { onChange, value }, fieldState: { error } }) => (
              <TextField
                error={!!error}
                helperText={error?.message}
                onChange={onChange}
                value={value}
                type="number"
                label="Target"
                inputProps={{ step: 0.1 }}
              />
            )}
          />
          <Controller
            control={control}
            name="endDate"
            defaultValue={dayjs().add(1, 'day')}
            rules={{
              required: true,
              validate: (value) => value.isAfter(dayjs()),
            }}
            render={({ field: { ref, onChange, value }, fieldState: { error } }) => (
              <DatePicker
                value={value}
                onChange={onChange}
                inputRef={ref}
                label="End"
                inputFormat="YYYY-MM-DD"
                minDate={dayjs().add(1, 'day')}
                renderInput={(inputProps) => (
                  <TextField {...inputProps} error={!!error} helperText={error?.message} />
                )}
              />
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

export default AddGoalDialog;
