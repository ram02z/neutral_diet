import { useState } from 'react';

import LoadingButton from '@mui/lab/LoadingButton';
import { Alert, Collapse, TextField } from '@mui/material';

import PasswordTextField from '@/components/PasswordTextField';
import { auth } from '@/core/firebase';
import useDefaultSignIn from '@/hooks/useDefaultSignIn';

function LogIn() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [signIn, , loading, error] = useDefaultSignIn(auth);
  const [open, setOpen] = useState(false);

  const handleSetPassword = (newPassword: string) => {
    setPassword(newPassword);
  };

  const handleFormSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    signIn(email, password);
    setEmail('');
    setPassword('');
    setOpen(true);
  };

  return (
    <>
      <form onSubmit={handleFormSubmit}>
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
    </>
  );
}

export default LogIn;
