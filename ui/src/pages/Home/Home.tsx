import { useEffect } from 'react';
import { useRecoilRefresher_UNSTABLE, useRecoilValue } from 'recoil';

import { Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import dayjs from 'dayjs';

import Carousel from '@/components/Carousel';
import CircularProgressWithLabel from '@/components/CircularProgressWithLabel';
import GoalCard from '@/components/GoalCard';
import GoalLinePlot from '@/components/GoalLinePlot';
import TrendCard from '@/components/TrendCard';
import Insights from '@/core/insights';
import { Meal } from '@/core/meal';
import {
  ActiveGoalState,
  LocalFoodItemLogStats,
  LocalUserSettingsState,
  UserInsightsState,
  UserProgressState,
} from '@/store/user';
import { toSerializableDate } from '@/utils/date';

function Home() {
  const todayStats = useRecoilValue(LocalFoodItemLogStats(toSerializableDate(dayjs())));
  const userInsights = useRecoilValue(UserInsightsState);
  const userProgress = useRecoilValue(UserProgressState);
  const refreshUserInsights = useRecoilRefresher_UNSTABLE(UserInsightsState);
  const userSettings = useRecoilValue(LocalUserSettingsState);
  const activeUserGoal = useRecoilValue(ActiveGoalState);
  const externalInsights = new Insights(userSettings.dietaryRequirement);

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
      <Grid container direction="column" justifyContent="center" alignItems="center" xs={11}>
        <Grid xs={10} sm={11} md={10} lg={9} xl={8}>
          <Typography textAlign="center" variant="h4">
            Goal
          </Typography>
          <Typography textAlign="center" color="text.secondary" variant="subtitle2">
            Active
          </Typography>
          {activeUserGoal ? (
            <GoalCard goal={activeUserGoal} active />
          ) : (
            <Typography>No active goals</Typography>
          )}
        </Grid>
        <Grid xs={10} sm={11} md={10} lg={9} xl={8}>
          <Typography textAlign="center" variant="h4">
            Trends
          </Typography>
          <Typography textAlign="center" color="text.secondary" variant="subtitle2">
            Daily
          </Typography>
          <Carousel>
            <TrendCard
              title="Your average"
              stat={userInsights.userDailyAverage}
              today={todayStats.totalCarbonFootprint}
            />
            {externalInsights.insights.map((insight, idx) => (
              <TrendCard
                key={idx}
                title={insight.title}
                stat={insight.emissions}
                today={todayStats.totalCarbonFootprint}
                source={insight.source}
              />
            ))}
            <TrendCard
              title="Users average"
              stat={userInsights.dailyGlobalAverage}
              today={todayStats.totalCarbonFootprint}
            />
            <TrendCard
              title="Users diet average"
              stat={userInsights.dailyGlobalAverageUserDietaryRequirement}
              today={todayStats.totalCarbonFootprint}
            />
          </Carousel>
        </Grid>
        <Grid xs={10} sm={11} md={10} lg={9} xl={8}>
          <Typography textAlign="center" variant="h4">
            Progress
          </Typography>
          <Carousel>
            <GoalLinePlot
              goal={userSettings.cfLimit}
              actualData={userProgress.all}
              title="Daily carbon footprint, all meals"
            />
            {Array.from(userProgress.meal).map(([key, value]) => (
              <GoalLinePlot
                key={key}
                goal={userSettings.cfLimit}
                actualData={value}
                title={`Daily carbon footprint, ${new Meal(key).getName()}`}
              />
            ))}
          </Carousel>
        </Grid>
      </Grid>
    </Grid>
  );
}

export default Home;
