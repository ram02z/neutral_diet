import { useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useRecoilState, useRecoilValue } from 'recoil';

import { TabContext, TabList, TabPanel } from '@mui/lab';
import { Box, Button, Stack, Tab, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import dayjs from 'dayjs';

import client from '@/api/user_service';
import AddGoalDialog from '@/components/AddGoalDialog';
import { FormValues } from '@/components/AddGoalDialog/types';
import Carousel from '@/components/Carousel';
import GoalCard from '@/components/GoalCard';
import GoalLinePlot from '@/components/GoalLinePlot';
import { Meal } from '@/core/meal';
import {
  CurrentUserHeadersState,
  SelectedUserGoalState,
  UserGoalProgressState,
  UserGoalsState,
  UserInsightsState,
} from '@/store/user';
import { toSerializableDate } from '@/utils/date';

function GoalList() {
  const [tab, setTab] = useState('1');
  const [userGoals, setUserGoals] = useRecoilState(UserGoalsState);
  const selectedUserGoal = useRecoilValue(SelectedUserGoalState);
  const selectedUserGoalData = useRecoilValue(UserGoalProgressState);
  const userInsights = useRecoilValue(UserInsightsState);
  const userHeaders = useRecoilValue(CurrentUserHeadersState);
  const [openAddDialog, setOpenAddDialog] = useState(false);

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    const target = parseFloat(data.targetCarbonFootprint);
    const today = toSerializableDate(dayjs().subtract(7, 'day'));
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

  const handleTabChange = (_: React.SyntheticEvent, newValue: string) => {
    setTab(newValue);
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
      direction="column"
      justifyContent="center"
      alignItems="center"
      width="100%"
      spacing={2}
    >
      <TabContext value={tab}>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <TabList onChange={handleTabChange} aria-label="lab API tabs example">
            <Tab label="Active" value="1" />
            <Tab label="Completed" value="2" />
          </TabList>
        </Box>
        <TabPanel value="1" sx={{ width: '100%' }}>
          {userGoals.active.map((goal, idx) => (
            <Grid key={idx} xs={12}>
              <GoalCard goal={goal} active />
            </Grid>
          ))}
        </TabPanel>
        <TabPanel value="2" sx={{ width: '100%' }}>
          {userGoals.completed.map((goal, idx) => (
            <Grid key={idx} xs={12}>
              <GoalCard goal={goal} active={false} />
            </Grid>
          ))}
        </TabPanel>
      </TabContext>
      <Grid>
        <Button variant="contained" onClick={handleOpenAddDialog}>
          Add goal
        </Button>
      </Grid>
      <Grid mt={2} xs={12}>
        <Typography textAlign="center" variant="h5">
          Progress
        </Typography>
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
            <Typography color="text.secondary">Click on a goal to view progress!</Typography>
          </>
        )}
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

export default GoalList;
