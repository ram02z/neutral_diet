import { BaseSyntheticEvent } from 'react';
import { Controller, useForm } from 'react-hook-form';

import { Button, Dialog, DialogActions, DialogTitle, TextField } from '@mui/material';
import { Stack } from '@mui/system';

import { FormValues } from './types';

type DisplayNameDialogProps = {
  onSubmit: (data: FormValues, event?: BaseSyntheticEvent<object, void, void | undefined>) => void;
  openDialog: boolean;
  handleClose: () => void;
  currentDisplayName: string;
};

/**
 * Dialog to allow users to update their display name.
 */
function DisplayNameDialog({
  onSubmit,
  openDialog,
  handleClose,
  currentDisplayName,
}: DisplayNameDialogProps) {
  const { handleSubmit, control } = useForm<FormValues>();

  return (
    <Dialog open={openDialog} onClose={handleClose}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <DialogTitle textAlign="center">Profile</DialogTitle>
        <Stack sx={{ px: 5 }} spacing={3}>
          <Controller
            control={control}
            name="displayName"
            defaultValue={currentDisplayName}
            rules={{ required: true }}
            render={({ field: { onChange, value }, fieldState: { error } }) => (
              <TextField
                label="Display Name"
                error={!!error}
                onChange={onChange}
                value={value}
              ></TextField>
            )}
          />
        </Stack>
        <DialogActions sx={{ mt: '24px' }}>
          <Button color="secondary" onClick={handleClose}>
            Cancel
          </Button>
          <Button type="submit">Save</Button>
        </DialogActions>
      </form>
    </Dialog>
  );
}

export default DisplayNameDialog;
