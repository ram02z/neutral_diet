import { useEffect, useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useRecoilRefresher_UNSTABLE, useRecoilState, useRecoilValue } from 'recoil';

import { Button, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import dayjs from 'dayjs';

import client from '@/api/user_service';
import AddGoalDialog from '@/components/AddGoalDialog';
import { FormValues } from '@/components/AddGoalDialog/types';
import Carousel from '@/components/Carousel';
import CircularProgressWithLabel from '@/components/CircularProgressWithLabel';
import GoalLinePlot from '@/components/GoalLinePlot';
import TrendCard from '@/components/TrendCard';
import DietaryRequirement from '@/core/dietary_requirements';
import { Meal } from '@/core/meal';
import {
  CurrentUserHeadersState,
  LocalFoodItemLogStats,
  LocalUserSettingsState,
  UserGoalsState,
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
  const dietaryRequirement = new DietaryRequirement(userSettings.dietaryRequirement);
  const [openAddDialog, setOpenAddDialog] = useState(false);
  const [userGoals, setUserGoals] = useRecoilState(UserGoalsState);
  const userHeaders = useRecoilValue(CurrentUserHeadersState);

  // eslint-disable-next-line react-hooks/exhaustive-deps
  useEffect(() => refreshUserInsights(), []);

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    const target = parseFloat(data.targetCarbonFootprint)
    const today = toSerializableDate(dayjs());
    client
      .addCarbonFootprintGoal(
        {
          carbonFootprintGoal: {
            description: data.description,
            startDate: today,
            endDate: toSerializableDate(data.endDate),
            startCarbonFootprint: userInsights.overallUserAverage,
            targetCarbonFootprint: target,
          },
        },
        { headers: userHeaders },
      )
      .then((res) => {
        setUserGoals((old) => {
          return {
            completed: old.completed,
            active: [
              ...old.active,
              {
                dbId: res.id,
                description: data.description,
                startDate: today,
                endDate: toSerializableDate(data.endDate),
                startCarbonFootprint: userInsights.overallUserAverage,
                targetCarbonFootprint: target,
              },
            ],
          };
        });
      });
    handleCloseAddDialog();
  };

  const handleOpenAddDialog = () => {
    setOpenAddDialog(true);
  };

  const handleCloseAddDialog = () => {
    setOpenAddDialog(false);
  };

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
      <Grid xs={10} sm={11} md={10} lg={9} xl={8}>
        <Typography textAlign="center" variant="h4">
          Goals
        </Typography>
        <Button onClick={handleOpenAddDialog}>Add goal</Button>
      </Grid>
      <Grid container direction="column" justifyContent="center" alignItems="center" xs={11}>
        <Grid xs={10} sm={11} md={10} lg={9} xl={8}>
          <Typography textAlign="center" variant="h4">
            Trends
          </Typography>
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
      <AddGoalDialog
        openDialog={openAddDialog}
        handleClose={handleCloseAddDialog}
        onSubmit={onSubmit}
        startCarbonFootprint={userInsights.overallUserAverage}
      />
    </Grid>
  );
}

export default Home;
