import { useState } from 'react';

import { Visibility, VisibilityOff } from '@mui/icons-material';
import { FilledInput, FormControl, IconButton, InputAdornment, InputLabel } from '@mui/material';

type PasswordTextFieldProps = {
  password: string;
  onChangeHandler: (newPassword: string) => void;
  error: boolean;
};

function PasswordTextField({ password, onChangeHandler, error }: PasswordTextFieldProps) {
  const [showPassword, setShowPassword] = useState(false);
  const handleClickShowPassword = () => setShowPassword((show) => !show);

  const handleMouseDownPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault();
  };
  return (
    <>
      <FormControl variant="filled" margin="dense" fullWidth required>
        <InputLabel htmlFor="filled-adornment-password">Password</InputLabel>
        <FilledInput
          margin="dense"
          id="filled-adornment-password"
          type={showPassword ? 'text' : 'password'}
          value={password}
          onChange={(e) => onChangeHandler(e.target.value as string)}
          error={error}
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
    </>
  );
}

export default PasswordTextField;
