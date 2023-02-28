import { useState } from 'react';

import { Google } from '@mui/icons-material';
import { LoadingButton } from '@mui/lab';
import { Alert, Collapse } from '@mui/material';
import { Stack } from '@mui/system';

import { auth } from '@/core/firebase';
import { useSignInWithGoogle } from '@/hooks/useSignInWithPopup';

function GoogleAuthButton() {
  const [signIn, , loading, error] = useSignInWithGoogle(auth);
  const [open, setOpen] = useState(false);

  const onClickHandler = () => {
    signIn();
    setOpen(true);
  };

  return (
    <Stack direction="column" spacing={1}>
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
      <LoadingButton
        loading={loading}
        onClick={onClickHandler}
        variant="contained"
        color="secondary"
        startIcon={<Google />}
        fullWidth
      >
        Continue with Google
      </LoadingButton>
    </Stack>
  );
}

export default GoogleAuthButton;
