import { useState } from 'react';

import { LoadingButton } from '@mui/lab';
import { Alert, Collapse, TextField } from '@mui/material';

import PasswordTextField from '@/components/PasswordTextField';
import { auth } from '@/core/firebase';
import useDefaultSignUp from '@/hooks/useDefaultSignUp';

function SignUp() {
  const [displayName, setDisplayName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [signUp, , loading, error] = useDefaultSignUp(auth);
  const [open, setOpen] = useState(false);

  const handleFormSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    signUp(displayName, email, password);
    setDisplayName('');
    setEmail('');
    setPassword('');
    setOpen(true);
  };

  const handleSetPassword = (newPassword: string) => {
    setPassword(newPassword);
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
            Error occurred. Try again!
          </Alert>
        </Collapse>
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
    </>
  );
}

export default SignUp;
