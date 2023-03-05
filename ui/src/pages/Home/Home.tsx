import { useEffect } from 'react';
import { useRecoilRefresher_UNSTABLE, useRecoilValue } from 'recoil';

import { Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import { UserInsightsState } from '@/store/user';

function Home() {
  const userInsights = useRecoilValue(UserInsightsState);
  const refreshUserInsights = useRecoilRefresher_UNSTABLE(UserInsightsState);

  // eslint-disable-next-line react-hooks/exhaustive-deps
  useEffect(() => refreshUserInsights(), []);

  return (
    <Grid container direction="column" alignItems="center" spacing={2} sx={{ mt: 4 }}>
      <Grid>
        <Typography>{userInsights.overallUserAverage.toFixed(3)}</Typography>
      </Grid>
      <Grid>
        <Typography>{userInsights.overallUser.toFixed(3)}</Typography>
      </Grid>
      <Grid>
        <Typography>{userInsights.noUserEntries}</Typography>
      </Grid>
      <Grid>
        <Typography>{userInsights.dailyGlobalAverage.toFixed(3)}</Typography>
      </Grid>
      <Grid>
        <Typography>{userInsights.dailyGlobalAverageUserDietaryRequirement.toFixed(3)}</Typography>
      </Grid>
    </Grid>
  );
}

export default Home;
