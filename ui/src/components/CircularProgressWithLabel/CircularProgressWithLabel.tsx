import { Box, CircularProgress, Typography, Stack } from '@mui/material';

type CircularProgressWithLabelProps = {
  value: number;
  size: number;
  remaining: number;
};

function CircularProgressWithLabel({ value, size, remaining }: CircularProgressWithLabelProps) {
  return (
    <Box sx={{ position: 'relative', display: 'inline-flex' }}>
      <CircularProgress
        size={size}
        thickness={4}
        sx={{
          color: (theme) => theme.palette.grey[theme.palette.mode === 'light' ? 200 : 800],
        }}
        variant="determinate"
        value={100}
      />
      <CircularProgress
        variant="determinate"
        sx={{
          position: 'absolute',
          left: 0,
        }}
        size={size}
        thickness={4}
        color={value > 100 ? 'error' : 'primary'}
        value={value > 100 ? 100 : value}
      />
      <Stack
        sx={{
          top: 0,
          left: 0,
          bottom: 0,
          right: 0,
          position: 'absolute',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
        }}
      >
        <Typography variant="h4" color="text.primary">
          {Math.abs(remaining).toFixed(3)}
        </Typography>
        <Typography variant="button" color="text.primary">
          CO<sub>2</sub>/kg {remaining > 0 ? "left" : "over"}
        </Typography>
      </Stack>
    </Box>
  );
}

export default CircularProgressWithLabel;
