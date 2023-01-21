import { Alert, Button, Typography } from '@mui/material';

import { User } from 'firebase/auth';
import { useConfirm } from 'material-ui-confirm';

import { ID_TOKEN_HEADER } from '@/api/transport';
import client from '@/api/user_service';
import { useSignOut } from '@/hooks/useSignOut';

type DeleteAccountProps = {
  user: User;
};

function DeleteAccount({ user }: DeleteAccountProps) {
  const confirm = useConfirm();
  const signOut = useSignOut();

  const handleClick = async () => {
    confirm({
      title: 'Delete account',
      content: (
        <div>
          <Alert severity="error">
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
      cancellationButtonProps: { color: 'info' },
      confirmationText: 'Delete',
      confirmationButtonProps: { color: 'error', variant: 'contained' },
    }).then(() => {
      user.getIdToken().then((idToken) => {
        const headers = new Headers();
        headers.set(ID_TOKEN_HEADER, idToken);
        client.deleteUser({}, { headers: headers });
        signOut();
      });
    });
  };

  return (
    <>
      <Button onClick={handleClick}>Delete account</Button>
    </>
  );
}

export default DeleteAccount;
