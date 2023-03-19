import { useState } from 'react';
import { useRecoilValue } from 'recoil';

import { TabContext, TabList, TabPanel } from '@mui/lab';
import { Box, Stack, Tab, Typography } from '@mui/material';

import { GoalCard } from '@/components/GoalCard';
import { UserGoalsState } from '@/store/user';

function GoalList() {
  const [tab, setTab] = useState('1');
  const userGoals = useRecoilValue(UserGoalsState);

  const handleTabChange = (_: React.SyntheticEvent, newValue: string) => {
    setTab(newValue);
  };

  return (
    <TabContext value={tab}>
      <Stack direction="column" justifyContent="center" alignItems="center" spacing={2}>
      <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
        <TabList onChange={handleTabChange}>
          <Tab label="Active" value="1" />
          <Tab label="Completed" value="2" />
        </TabList>
      </Box>
      <TabPanel value="1" sx={{ width: '100%' }}>
        {userGoals.active.length == 0 && (
          <Typography color="text.secondary">No active goals.</Typography>
        )}
        {userGoals.active.map((goal, idx) => (
          <GoalCard key={idx} goal={goal} active />
        ))}
      </TabPanel>
      <TabPanel value="2" sx={{ width: '100%' }}>
        {userGoals.completed.length == 0 && (
          <Typography color="text.secondary">No completed goals.</Typography>
        )}
        {userGoals.completed.map((goal, idx) => (
          <GoalCard key={idx} goal={goal} active={false} />
        ))}
      </TabPanel>
      </Stack>
    </TabContext>
  );
}

export default GoalList;
