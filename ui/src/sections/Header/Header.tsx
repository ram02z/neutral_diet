import { Link } from 'react-router-dom';

import AccountCircle from '@mui/icons-material/AccountCircle';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import IconButton from '@mui/material/IconButton';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';

import ThemeToggler from '@/components/ThemeToggler';
import routes from '@/routes';
import { Pages } from '@/routes/types';

import type { HeaderProps } from './types';

function Header({ title }: HeaderProps) {
  return (
    <Box sx={{ flexGrow: 1, pb: 5 }}>
      <AppBar position="fixed" elevation={0} sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}>
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            {title}
          </Typography>
          <ThemeToggler />
          <IconButton
            size="large"
            aria-label="account of current user"
            color="inherit"
            component={Link}
            to={routes[Pages.Account].path as string}
          >
            <AccountCircle />
          </IconButton>
        </Toolbar>
      </AppBar>
    </Box>
  );
}

export default Header;
