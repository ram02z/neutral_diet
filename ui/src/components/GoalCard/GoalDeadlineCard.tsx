import TrendingDownIcon from '@mui/icons-material/TrendingDown';
import TrendingFlatIcon from '@mui/icons-material/TrendingFlat';
import TrendingUpIcon from '@mui/icons-material/TrendingUp';
import { Card, CardContent, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import { MIN_CARD_WIDTH } from '@/config';
import { LocalUserGoal } from '@/store/user/types';
import { getDateString } from '@/utils/date';

type GoalDeadlineCardProps = {
  goal: LocalUserGoal;
  currentUserAverage: number;
};

function GoalDeadlineCard({ goal, currentUserAverage }: GoalDeadlineCardProps) {
  let statDiff = 0;
  if (currentUserAverage > 0) {
    statDiff = ((goal.targetCarbonFootprint - currentUserAverage) / currentUserAverage) * 100.0;
  }
  let textColor = 'text.secondary';
  if (statDiff > 0) {
    textColor = 'success.main';
  } else if (statDiff < 0) {
    textColor = 'error.main';
  }
  return (
    <Card sx={{ minWidth: MIN_CARD_WIDTH, display: 'flex' }}>
      <Grid container columns={5} direction="column">
        <CardContent>
          <Grid>
            <Typography variant="h5" component="div">
              <Typography variant="caption" color="text.secondary" display="block" gutterBottom>
                Goal
              </Typography>
              {goal.description}
            </Typography>
          </Grid>
          <Grid container direction="row" alignItems="center">
            <Grid sx={{ pr: 1 }}>
              {statDiff > 0 && <TrendingUpIcon color="success" />}
              {statDiff == 0 && <TrendingFlatIcon color="secondary" />}
              {statDiff < 0 && <TrendingDownIcon color="error" />}
            </Grid>
            <Grid>
              <Typography variant="subtitle1" color={textColor}>
                <b>{Math.abs(statDiff).toFixed(2)}%</b>
              </Typography>
            </Grid>
          </Grid>
        </CardContent>
      </Grid>
      <Grid direction="column" alignSelf="center" marginLeft="auto">
        <Grid textAlign="center">
          <Typography variant="h2" color="info.light">
            {goal.endDate.diff(goal.startDate, 'day')}
          </Typography>
          <Typography variant="h5" color="info.main">DAYS</Typography>
          <Typography variant="subtitle2" color="info.main">LEFT</Typography>
        </Grid>
      </Grid>
    </Card>
  );
}

export default GoalDeadlineCard;
