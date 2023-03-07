import { useEffect } from 'react';
import { useRecoilRefresher_UNSTABLE, useRecoilValue } from 'recoil';

import Grid from '@mui/material/Unstable_Grid2';

import dayjs from 'dayjs';

import UserStatCard from '@/components/UserStatCard';
import DietaryRequirement from '@/core/dietary_requirements';
import { LocalFoodItemLogStats, LocalUserSettingsState, UserInsightsState } from '@/store/user';
import { toSerializableDate } from '@/utils/date';
import { Typography } from '@mui/material';

function Home() {
  const userSettings = useRecoilValue(LocalUserSettingsState);
  const todayStats = useRecoilValue(LocalFoodItemLogStats(toSerializableDate(dayjs())));
  const userInsights = useRecoilValue(UserInsightsState);
  const refreshUserInsights = useRecoilRefresher_UNSTABLE(UserInsightsState);
  const dietaryRequirement = new DietaryRequirement(userSettings.dietaryRequirement);

  // eslint-disable-next-line react-hooks/exhaustive-deps
  useEffect(() => refreshUserInsights(), []);

  return (
    <Grid container direction="column" alignItems="center" spacing={2} sx={{ mt: 4 }}>
      <Grid>
        <Typography>Active streak: {userInsights.activeStreak}</Typography>
      </Grid>
      <Grid>
        <UserStatCard
          title="Today carbon footprint"
          carbonFootprint={todayStats.totalCarbonFootprint}
        />
      </Grid>
      <Grid>
        <UserStatCard
          title="Overall user average"
          carbonFootprint={userInsights.overallUserAverage}
        />
      </Grid>
      <Grid>
        <UserStatCard
          title="Daily global average"
          carbonFootprint={userInsights.dailyGlobalAverage}
        />
      </Grid>
      <Grid>
        <UserStatCard
          title={`Daily global ${dietaryRequirement.getSettingName()} average`}
          carbonFootprint={userInsights.dailyGlobalAverageUserDietaryRequirement}
        />
      </Grid>
    </Grid>
  );
}

export default Home;
