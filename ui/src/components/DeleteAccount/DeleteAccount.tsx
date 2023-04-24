import { useRecoilValue } from 'recoil';

import { Alert, Button, Typography } from '@mui/material';

import { User } from 'firebase/auth';
import { useConfirm } from 'material-ui-confirm';
import { useSnackbar } from 'notistack';

import client from '@/api/user_service';
import { useSignOut } from '@/hooks/useSignOut';
import { CurrentUserHeadersState } from '@/store/user';

type DeleteAccountProps = {
  user: User;
};

/**
 * Button to delete user account with confirmation.
 */
function DeleteAccount({ user }: DeleteAccountProps) {
  const userHeaders = useRecoilValue(CurrentUserHeadersState);
  const confirm = useConfirm();
  const signOut = useSignOut();
  const { enqueueSnackbar } = useSnackbar();

  const handleClick = async () => {
    confirm({
      title: 'Delete account',
      content: (
        <div>
          <Alert variant="filled" severity="error">
            After you have deleted an account, it will be permanently deleted. Accounts cannot be
            recovered.
          </Alert>
          <br />
          <Typography variant="caption" display="block">
            User account
          </Typography>
          {user.email}
        </div>
      ),
      cancellationButtonProps: { color: 'secondary' },
      confirmationText: 'Delete',
      confirmationButtonProps: { color: 'error' },
    }).then(() => {
      client
        .deleteUser({}, { headers: userHeaders })
        .then(() => {
          enqueueSnackbar('Deleted account successfully.', { variant: 'success' });
          signOut();
        })
        .catch((err) => {
          enqueueSnackbar('Could not delete account.', { variant: 'error' });
          console.error(err);
        });
    });
  };

  return (
    <>
      <Button variant="contained" color="error" onClick={handleClick}>
        Delete account
      </Button>
    </>
  );
}

export default DeleteAccount;
