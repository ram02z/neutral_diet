import { useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useRecoilState, useRecoilValue } from 'recoil';

import { TabContext, TabList, TabPanel } from '@mui/lab';
import { Box, Button, Tab } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import dayjs from 'dayjs';

import client from '@/api/user_service';
import AddGoalDialog from '@/components/AddGoalDialog';
import { FormValues } from '@/components/AddGoalDialog/types';
import GoalCard from '@/components/GoalCard';
import { CurrentUserHeadersState, UserGoalsState, UserInsightsState } from '@/store/user';
import { toSerializableDate } from '@/utils/date';

function GoalList() {
  const [tab, setTab] = useState('1');
  const [userGoals, setUserGoals] = useRecoilState(UserGoalsState);
  const userInsights = useRecoilValue(UserInsightsState);
  const userHeaders = useRecoilValue(CurrentUserHeadersState);
  const [openAddDialog, setOpenAddDialog] = useState(false);

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    const target = parseFloat(data.targetCarbonFootprint);
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
        <Button variant="contained" onClick={handleOpenAddDialog}>Add goal</Button>
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
