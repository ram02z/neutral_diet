import { useEffect } from 'react';
import { useRecoilRefresher_UNSTABLE, useRecoilValue } from 'recoil';

import { Typography } from '@mui/material';
import { Stack } from '@mui/system';

import { UserInsightsState } from '@/store/user';

function Home() {
  const userInsights = useRecoilValue(UserInsightsState);
  const refreshUserInsights = useRecoilRefresher_UNSTABLE(UserInsightsState);

  // eslint-disable-next-line react-hooks/exhaustive-deps
  useEffect(() => refreshUserInsights(), []);

  return (
    <Stack>
      <Typography>{userInsights.overallAverageCarbonFootprint}</Typography>
      <Typography>{userInsights.noEntries}</Typography>
      <Typography>{userInsights.overallCarbonFootprint}</Typography>
    </Stack>
  );
}

export default Home;
