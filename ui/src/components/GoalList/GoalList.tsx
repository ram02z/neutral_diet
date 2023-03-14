import { useState } from 'react';
import { useRecoilValue } from 'recoil';

import { TabContext, TabList, TabPanel } from '@mui/lab';
import { Box, Tab } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import GoalCard from '@/components/GoalCard';
import { UserGoalsState } from '@/store/user';

function GoalList() {
  const [tab, setTab] = useState('1');
  const userGoals = useRecoilValue(UserGoalsState);

  const handleTabChange = (_: React.SyntheticEvent, newValue: string) => {
    setTab(newValue);
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
    </Grid>
  );
}

export default GoalList;
