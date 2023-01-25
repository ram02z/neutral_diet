import { useState } from 'react';

import { LoadingButton } from '@mui/lab';
import { TextField, Typography } from '@mui/material';

import { auth } from '@/core/firebase';
import useDefaultSignUp from '@/hooks/useDefaultSignUp';

import PasswordTextField from '../PasswordTextField';

function SignUp() {
  const [displayName, setDisplayName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [signUp, , loading, error] = useDefaultSignUp(auth);

  const handleFormSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    signUp(displayName, email, password);
    setDisplayName('');
    setEmail('');
    setPassword('');
  };

  const handleSetPassword = (newPassword: string) => {
    setPassword(newPassword);
  };

  return (
    <>
      <form onSubmit={handleFormSubmit}>
        <TextField
          variant="filled"
          margin="dense"
          label="Name"
          type="text"
          placeholder="Enter name"
          value={displayName}
          onChange={(e) => setDisplayName(e.target.value)}
          fullWidth
          required
        />
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

export default SignUp;
