import { useEffect } from 'react';
import { useRecoilRefresher_UNSTABLE, useRecoilValue } from 'recoil';

import { Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import dayjs from 'dayjs';

import Carousel from '@/components/Carousel';
import CircularProgressWithLabel from '@/components/CircularProgressWithLabel';
import TrendCard from '@/components/TrendCard';
import DietaryRequirement from '@/core/dietary_requirements';
import { LocalFoodItemLogStats, LocalUserSettingsState, UserInsightsState } from '@/store/user';
import { toSerializableDate } from '@/utils/date';

function Home() {
  const todayStats = useRecoilValue(LocalFoodItemLogStats(toSerializableDate(dayjs())));
  const userInsights = useRecoilValue(UserInsightsState);
  const refreshUserInsights = useRecoilRefresher_UNSTABLE(UserInsightsState);
  const userSettings = useRecoilValue(LocalUserSettingsState);
  const dietaryRequirement = new DietaryRequirement(userSettings.dietaryRequirement);

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
          {userInsights.streakLength > 0 &&
            `Your longest streak ${userInsights.isStreakActive ? 'is' : 'was'} ${
              userInsights.streakLength
            } days`}
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
          title="UK daily diet average"
          stat={dietaryRequirement.getMeanEmissionsPerDay()}
          today={todayStats.totalCarbonFootprint}
        />
        <TrendCard
          title="User daily average"
          stat={userInsights.dailyGlobalAverage}
          today={todayStats.totalCarbonFootprint}
        />
        <TrendCard
          title="User daily diet average"
          stat={userInsights.dailyGlobalAverageUserDietaryRequirement}
          today={todayStats.totalCarbonFootprint}
        />
      </Carousel>
    </Grid>
  );
}

export default Home;
