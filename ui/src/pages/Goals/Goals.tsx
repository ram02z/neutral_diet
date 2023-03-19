import { useEffect, useMemo, useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useRecoilRefresher_UNSTABLE, useRecoilValue, useSetRecoilState } from 'recoil';

import { Button, Tooltip, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import dayjs from 'dayjs';

import client from '@/api/user_service';
import AddGoalDialog from '@/components/AddGoalDialog';
import { FormValues } from '@/components/AddGoalDialog/types';
import Carousel from '@/components/Carousel';
import GoalLinePlot from '@/components/GoalLinePlot';
import GoalList from '@/components/GoalList';
import { RecommendGoalsButton } from '@/components/RecommendGoalDialog';
import { Meal } from '@/core/meal';
import {
  CurrentUserHeadersState,
  SelectedUserGoalState,
  UserGoalProgressState,
  UserGoalsState,
  UserInsightsState,
} from '@/store/user';
import { toSerializableDate } from '@/utils/date';

function Goals() {
  const setUserGoals = useSetRecoilState(UserGoalsState);
  const userHeaders = useRecoilValue(CurrentUserHeadersState);
  const selectedUserGoal = useRecoilValue(SelectedUserGoalState);
  const selectedUserGoalData = useRecoilValue(UserGoalProgressState);
  const userInsights = useRecoilValue(UserInsightsState);
  const refreshUserInsights = useRecoilRefresher_UNSTABLE(UserInsightsState);
  const [openAddDialog, setOpenAddDialog] = useState(false);
  const newUser = useMemo(() => userInsights.userDailyAverage === 0, [userInsights]);

  // eslint-disable-next-line react-hooks/exhaustive-deps
  useEffect(() => refreshUserInsights(), []);

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    const target = parseFloat(data.targetCarbonFootprint);
    const today = dayjs();
    client
      .addCarbonFootprintGoal(
        {
          carbonFootprintGoal: {
            description: data.description,
            startDate: toSerializableDate(today),
            endDate: toSerializableDate(data.endDate),
            startCarbonFootprint: userInsights.userDailyAverage,
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
                endDate: data.endDate,
                startCarbonFootprint: userInsights.userDailyAverage,
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
    <Grid
      container
      spacing={2}
      direction="column"
      alignItems="center"
      justifyContent="center"
      textAlign="center"
      width={{ xs: '60vw', md: '40vw' }}
    >
      <Grid xs={12}>
        <Typography variant="h4">Carbon footprint goals</Typography>
      </Grid>
      <Grid xs={12}>
        <GoalList />
      </Grid>
      <Grid xs={12}>
        <Tooltip title={newUser ? 'Add food to your log or try a recommended goal' : ''}>
          <span>
            <Button variant="contained" onClick={handleOpenAddDialog} disabled={newUser}>
              Add goal
            </Button>
          </span>
        </Tooltip>
      </Grid>
      <Grid xs={12}>
        <RecommendGoalsButton />
      </Grid>
      <Grid xs={12} mt={2}>
        <Typography variant="h5">Progress</Typography>
      </Grid>
      <Grid xs={12}>
        {selectedUserGoal && selectedUserGoalData ? (
          <Carousel>
            <GoalLinePlot
              goal={selectedUserGoal.targetCarbonFootprint}
              actualData={selectedUserGoalData.all}
              title={`${selectedUserGoal.description}, all meals`}
            />
            {Array.from(selectedUserGoalData.meal).map(([key, value]) => (
              <GoalLinePlot
                key={key}
                goal={selectedUserGoal.targetCarbonFootprint}
                actualData={value}
                title={`${selectedUserGoal.description}, ${new Meal(key).getName()}`}
              />
            ))}
          </Carousel>
        ) : (
          <>
            <Typography color="text.secondary">No goal selected.</Typography>
            <Typography color="text.secondary">Click on a goal to view your progress!</Typography>
          </>
        )}
      </Grid>
      <AddGoalDialog
        openDialog={openAddDialog}
        handleClose={handleCloseAddDialog}
        onSubmit={onSubmit}
        startCarbonFootprint={userInsights.userDailyAverage}
      />
    </Grid>
  );
}

export default Goals;
