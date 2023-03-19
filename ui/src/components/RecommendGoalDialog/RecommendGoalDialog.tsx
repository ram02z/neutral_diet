import { BaseSyntheticEvent, useEffect, useMemo, useState } from 'react';
import { Controller, useForm } from 'react-hook-form';

import {
  Button,
  Dialog,
  DialogActions,
  DialogTitle,
  MenuItem,
  Stack,
  TextField,
} from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers';

import dayjs from 'dayjs';

import { Insight } from '@/core/insights';

import { FormValues } from './types';

type RecommendGoalDialogProps = {
  onSubmit: (data: FormValues, event?: BaseSyntheticEvent<object, void, void | undefined>) => void;
  openDialog: boolean;
  handleClose: () => void;
  startCarbonFootprint: number;
  goals: Insight[];
};

function RecommendGoalDialog({
  onSubmit,
  openDialog,
  handleClose,
  startCarbonFootprint,
  goals,
}: RecommendGoalDialogProps) {
  const { handleSubmit, control, setValue } = useForm<FormValues>();
  const [goalIdx, setGoalIdx] = useState(0);
  const description = useMemo(() => `Meet ${goals[goalIdx].title} emissions`, [goalIdx, goals]);
  const targetCarbonFootprint = useMemo(
    () => goals[goalIdx].emissions.toString(),
    [goalIdx, goals],
  );

  useEffect(() => {
    setValue('description', description);
  }, [description, setValue]);

  useEffect(() => {
    setValue('targetCarbonFootprint', targetCarbonFootprint);
  }, [targetCarbonFootprint, setValue]);

  const handleChange = (event: { target: { value: string } }) => {
    setGoalIdx(parseInt(event.target.value));
  };

  return (
    <Dialog fullWidth open={openDialog} onClose={handleClose}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <DialogTitle textAlign="center">Recommended goals</DialogTitle>
        <Stack sx={{ px: 5 }} spacing={3}>
          <TextField select label="Goal" value={goalIdx} onChange={handleChange}>
            {goals.map((goal, idx) => (
              <MenuItem key={idx} value={idx}>
                {goal.title}
              </MenuItem>
            ))}
          </TextField>
          <TextField disabled label="Description" value={description} />
          <TextField
            aria-readonly
            InputProps={{
              readOnly: true,
            }}
            disabled
            value={startCarbonFootprint.toFixed(3)}
            type="number"
            label="Current daily average"
          />
          <TextField disabled type="number" label="Target" value={targetCarbonFootprint} />
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

export default RecommendGoalDialog;
