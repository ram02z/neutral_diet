import { useState } from 'react';

import LoadingButton from '@mui/lab/LoadingButton';
import { TextField, Typography } from '@mui/material';

import PasswordTextField from '@/components/PasswordTextField';
import { auth } from '@/core/firebase';
import useDefaultSignIn from '@/hooks/useDefaultSignIn';

function LogIn() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [signIn, , loading, error] = useDefaultSignIn(auth);

  const handleSetPassword = (newPassword: string) => {
    setPassword(newPassword);
  };

  const handleFormSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    signIn(email, password);
    setEmail('');
    setPassword('');
  };

  return (
    <>
      <form onSubmit={handleFormSubmit}>
        <TextField
          variant="filled"
          margin="dense"
          label="Email"
          placeholder="Enter email"
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value as string)}
          fullWidth
          required
        />
        <PasswordTextField password={password} onChangeHandler={handleSetPassword} />
        <LoadingButton loading={loading} variant="contained" type="submit" fullWidth>
          Continue
        </LoadingButton>
      </form>
      {error !== undefined && <Typography>Error occurred. Try again! </Typography>}
    </>
  );
}

export default LogIn;
