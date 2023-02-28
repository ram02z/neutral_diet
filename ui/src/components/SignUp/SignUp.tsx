import { useState } from 'react';
import { Controller, SubmitHandler, useForm } from 'react-hook-form';

import { LoadingButton } from '@mui/lab';
import { Alert, Collapse, Stack, TextField } from '@mui/material';

import PasswordTextField from '@/components/PasswordTextField';
import { auth } from '@/core/firebase';
import useDefaultSignUp from '@/hooks/useDefaultSignUp';

import { FormValues } from './types';

function SignUp() {
  const { handleSubmit, control } = useForm<FormValues>();
  const [signUp, , loading, error] = useDefaultSignUp(auth);
  const [open, setOpen] = useState(false);

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    signUp(data.displayName, data.email, data.password);
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
            Error occurred. Try again later!
          </Alert>
        </Collapse>
        <Controller
          control={control}
          name="displayName"
          rules={{ required: true }}
          render={({ field: { onChange, value }, fieldState: { error } }) => (
            <TextField
              error={!!error}
              variant="filled"
              margin="dense"
              label="Name"
              placeholder="Enter name"
              type="text"
              value={value}
              onChange={onChange}
              fullWidth
              required
            />
          )}
        />
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

export default SignUp;
