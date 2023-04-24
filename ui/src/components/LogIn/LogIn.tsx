import { useState } from 'react';
import { Controller, SubmitHandler, useForm } from 'react-hook-form';

import LoadingButton from '@mui/lab/LoadingButton';
import { Alert, Collapse, TextField } from '@mui/material';
import { Stack } from '@mui/system';

import PasswordTextField from '@/components/PasswordTextField';
import { auth } from '@/core/firebase';
import useDefaultSignIn from '@/hooks/useDefaultSignIn';

import { FormValues } from './types';

/**
 * Login form with input validation.
 */
function LogIn() {
  const { handleSubmit, control } = useForm<FormValues>();
  const [signIn, , loading, error] = useDefaultSignIn(auth);
  const [open, setOpen] = useState(false);

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    signIn(data.email, data.password);
    setOpen(true);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Stack direction="column" spacing={2}>
        <Collapse in={error && open}>
          <Alert
            icon={false}
            severity="error"
            onClose={() => {
              setOpen(false);
            }}
          >
            Incorrect username or password.
          </Alert>
        </Collapse>
        <Controller
          control={control}
          name="email"
          rules={{ required: true }}
          render={({ field: { onChange, value }, fieldState: { error } }) => (
            <TextField
              error={!!error}
              variant="filled"
              margin="dense"
              label="Email"
              placeholder="Enter email"
              type="email"
              value={value}
              onChange={onChange}
              fullWidth
              required
            />
          )}
        />
        <Controller
          control={control}
          name="password"
          rules={{ required: true }}
          render={({ field: { onChange, value }, fieldState: { error } }) => (
            <PasswordTextField password={value} onChangeHandler={onChange} error={!!error} />
          )}
        />
        <LoadingButton loading={loading} variant="contained" type="submit" fullWidth>
          Continue
        </LoadingButton>
      </Stack>
    </form>
  );
}

export default LogIn;
