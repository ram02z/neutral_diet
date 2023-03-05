import { useEffect } from 'react';
import { useRecoilRefresher_UNSTABLE, useRecoilValue } from 'recoil';

import { Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import UserStatCard from '@/components/UserStatCard';
import DietaryRequirement from '@/core/dietary_requirements';
import { LocalUserSettingsState, UserInsightsState } from '@/store/user';

function Home() {
  const userSettings = useRecoilValue(LocalUserSettingsState);
  const userInsights = useRecoilValue(UserInsightsState);
  const refreshUserInsights = useRecoilRefresher_UNSTABLE(UserInsightsState);
  const dietaryRequirement = new DietaryRequirement(userSettings.dietaryRequirement);

  // eslint-disable-next-line react-hooks/exhaustive-deps
  useEffect(() => refreshUserInsights(), []);

  return (
    <Grid container direction="column" alignItems="center" spacing={2} sx={{ mt: 4 }}>
      <Grid>
        <UserStatCard title="Daily average" carbonFootprint={userInsights.overallUserAverage} />
      </Grid>
      <Grid>
        <UserStatCard title="Number of entries" carbonFootprint={userInsights.noUserEntries} />
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
