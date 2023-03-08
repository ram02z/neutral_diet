import { useEffect } from 'react';
import { useRecoilRefresher_UNSTABLE, useRecoilValue } from 'recoil';

import { Paper, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import dayjs from 'dayjs';

import Carousel from '@/components/Carousel';
import CircularProgressWithLabel from '@/components/CircularProgressWithLabel';
import { LocalFoodItemLogStats, UserInsightsState } from '@/store/user';
import { toSerializableDate } from '@/utils/date';
import TrendCard from '@/components/TrendCard';

function Home() {
  const todayStats = useRecoilValue(LocalFoodItemLogStats(toSerializableDate(dayjs())));
  const userInsights = useRecoilValue(UserInsightsState);
  const refreshUserInsights = useRecoilRefresher_UNSTABLE(UserInsightsState);

  // eslint-disable-next-line react-hooks/exhaustive-deps
  useEffect(() => refreshUserInsights(), []);

  return (
    <Grid container direction="column" alignItems="center" spacing={2} sx={{ mt: 4, mb: 8 }}>
      <Grid>
        <CircularProgressWithLabel
          value={todayStats.carbonFootprintGoalPercent}
          size={200}
          remaining={todayStats.carbonFootprintRemaining}
        />
      </Grid>
      <Grid>
        <Typography variant="subtitle2" color="secondary">
          Your longest streak is <b>{userInsights.streakLength}</b> days,
          {userInsights.isStreakActive ? ' and is still active!' : ' but is no longer active!'}
        </Typography>
      </Grid>
      <Grid>
        <Typography variant="h4">Trends</Typography>
      </Grid>
      <Carousel>
        <TrendCard
          title="Your daily average"
          stat={userInsights.overallUserAverage}
          today={todayStats.totalCarbonFootprint}
        />
        <TrendCard
          title="Global daily average"
          stat={userInsights.dailyGlobalAverage}
          today={todayStats.totalCarbonFootprint}
        />
        <TrendCard
          title="Global daily diet average"
          stat={userInsights.dailyGlobalAverageUserDietaryRequirement}
          today={todayStats.totalCarbonFootprint}
        />
      </Carousel>
    </Grid>
  );
}

export default Home;
