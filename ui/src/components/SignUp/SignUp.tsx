import { useState } from 'react';

import { Visibility, VisibilityOff } from '@mui/icons-material';
import { LoadingButton } from '@mui/lab';
import {
  FilledInput,
  FormControl,
  IconButton,
  InputAdornment,
  InputLabel,
  TextField,
  Typography,
} from '@mui/material';

import { auth } from '@/core/firebase';
import useDefaultSignUp from '@/hooks/useDefaultSignUp';

function SignUp() {
  const [displayName, setDisplayName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [signUp, , loading, error] = useDefaultSignUp(auth);

  const handleClickShowPassword = () => setShowPassword((show) => !show);

  const handleMouseDownPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault();
  };

  const handleFormSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    signUp(displayName, email, password);
    setDisplayName('');
    setEmail('');
    setPassword('');
  };

  return (
    <>
      <Typography>Create your account</Typography>
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
        <FormControl variant="filled" margin="dense" fullWidth required>
          <InputLabel htmlFor="filled-adornment-password">Password</InputLabel>
          <FilledInput
            margin="dense"
            id="filled-adornment-password"
            type={showPassword ? 'text' : 'password'}
            value={password}
            onChange={(e) => setPassword(e.target.value as string)}
            endAdornment={
              <InputAdornment position="end">
                <IconButton
                  aria-label="toggle password visibility"
                  onClick={handleClickShowPassword}
                  onMouseDown={handleMouseDownPassword}
                  edge="end"
                >
                  {showPassword ? <VisibilityOff /> : <Visibility />}
                </IconButton>
              </InputAdornment>
            }
          />
        </FormControl>
        <LoadingButton loading={loading} variant="contained" type="submit" fullWidth>
          Continue
        </LoadingButton>
      </form>
      {error !== undefined && <Typography>Error occurred. Try again! </Typography>}
    </>
  );
}

export default SignUp;
