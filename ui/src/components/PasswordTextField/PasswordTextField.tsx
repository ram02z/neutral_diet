import { useState } from 'react';

import { Visibility, VisibilityOff } from '@mui/icons-material';
import {
  FilledInput,
  FormControl,
  FormHelperText,
  IconButton,
  InputAdornment,
  InputLabel,
} from '@mui/material';

type PasswordTextFieldProps = {
  password: string;
  onChangeHandler: (newPassword: string) => void;
  error: boolean;
  errorText?: string;
};

/**
 * Text field with a visibility toggle.
 */
function PasswordTextField({
  password,
  onChangeHandler,
  error,
  errorText,
}: PasswordTextFieldProps) {
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
          aria-describedby="password-error-text"
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
        {errorText && (
          <FormHelperText error id="password-error-text">
            {errorText}
          </FormHelperText>
        )}
      </FormControl>
    </>
  );
}

export default PasswordTextField;
