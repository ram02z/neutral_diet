import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import ThemeToggler from '@/components/ThemeToggler';

function Header() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar color="transparent" elevation={1} position="static"><ThemeToggler /></AppBar>
    </Box>
  );
}

export default Header;
