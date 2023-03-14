import { Stack, Typography } from '@mui/material';

import GoalList from '@/components/GoalList';

function Goals() {
  return (
    <Stack
      direction="column"
      spacing={4}
      alignItems="center"
      justifyContent="center"
      width={{ xs: '60vw', md: '40vw' }}
    >
      <Typography textAlign="center" variant="h4">
        Carbon footprint goals
      </Typography>
      <GoalList />
    </Stack>
  );
}

export default Goals;
