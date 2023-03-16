import { useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { Link, Navigate, Outlet, useOutlet } from 'react-router-dom';
import { useRecoilState } from 'recoil';

import CrisisAlertRoundedIcon from '@mui/icons-material/CrisisAlertRounded';
import SettingsRoundedIcon from '@mui/icons-material/SettingsRounded';
import { Button, Divider, IconButton, Stack, Typography, useTheme } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import DeleteAccount from '@/components/DeleteAccount';
import DisplayNameDialog from '@/components/DisplayNameDialog';
import { FormValues } from '@/components/DisplayNameDialog/types';
import Loading from '@/components/Loading';
import UserAvatar from '@/components/UserAvatar';
import { useCurrentUser } from '@/hooks/useCurrentUser';
import { useSignOut } from '@/hooks/useSignOut';
import routes from '@/routes';
import { Pages } from '@/routes/types';
import { CurrentUserDisplayName } from '@/store/user';

function Account() {
  const user = useCurrentUser();
  const theme = useTheme();
  const [displayName, setDisplayName] = useRecoilState(CurrentUserDisplayName);
  const outlet = useOutlet();
  const signOut = useSignOut();
  const [openNameDialog, setOpenNameDialog] = useState(false);

  const onDisplayNameSave: SubmitHandler<FormValues> = (data) => {
    setDisplayName(data.displayName);
    handleCloseNameDialog();
  };

  const handleOpenNameDialog = () => {
    setOpenNameDialog(true);
  };

  const handleCloseNameDialog = () => {
    setOpenNameDialog(false);
  };

  if (user === undefined) {
    return (
      <Grid container justifyContent="center" alignItems="center" height="100%">
        <Loading />
      </Grid>
    );
  } else if (user === null) {
    return <Navigate to={routes[Pages.Auth].path} />;
  } else {
    return (
      <Stack spacing={2} justifyContent="center" alignItems="center" pt="4vh" direction="column">
        <IconButton onClick={handleOpenNameDialog}>
          <UserAvatar
            sx={{ width: 80, height: 80, fontSize: 40, bgcolor: theme.palette.primary.main }}
            name={displayName ?? ''}
          />
        </IconButton>
        <Typography variant="h5" color="text.secondary">
          {user.email}
        </Typography>
        <Divider flexItem sx={{ pt: '2vh' }} />
        <Stack direction="column" spacing={4} alignItems="center" py="4vh">
          {outlet ? (
            <>
              <Outlet />
              <Button variant="contained" color="secondary" component={Link} to="/account">
                Back
              </Button>
            </>
          ) : (
            <>
              <Button
                fullWidth
                color="primary"
                variant="contained"
                component={Link}
                to="goals"
                startIcon={<CrisisAlertRoundedIcon />}
              >
                Goals
              </Button>
              <Button
                fullWidth
                color="primary"
                variant="contained"
                component={Link}
                to="settings"
                startIcon={<SettingsRoundedIcon />}
              >
                Settings
              </Button>
              <Button variant="contained" onClick={signOut}>
                Log out
              </Button>
              <DeleteAccount user={user} />
            </>
          )}
        </Stack>
        <DisplayNameDialog
          openDialog={openNameDialog}
          handleClose={handleCloseNameDialog}
          onSubmit={onDisplayNameSave}
          currentDisplayName={displayName ?? ''}
        />
      </Stack>
    );
  }
}

export default Account;
