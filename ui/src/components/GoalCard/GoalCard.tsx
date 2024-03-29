import { useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil';

import { Delete } from '@mui/icons-material';
import {
  Alert,
  Card,
  CardActionArea,
  CardActions,
  CardContent,
  IconButton,
  Typography,
} from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import { useConfirm } from 'material-ui-confirm';
import { useSnackbar } from 'notistack';

import client from '@/api/user_service';
import { MIN_CARD_WIDTH } from '@/config';
import { CurrentUserHeadersState, SelectedUserGoalState, UserGoalsState } from '@/store/user';
import { LocalUserGoal } from '@/store/user/types';
import { getDateString } from '@/utils/date';

type GoalCardProps = {
  goal: LocalUserGoal;
  active: boolean;
};

/**
 * Card containing information about a carbon footprint goal.
 * Handles the goal deletion logic.
 */
function GoalCard({ goal, active }: GoalCardProps) {
  const setUserGoals = useSetRecoilState(UserGoalsState);
  const [selectedUserGoal, setSelectedUserGoal] = useRecoilState(SelectedUserGoalState);
  const userHeaders = useRecoilValue(CurrentUserHeadersState);
  const confirm = useConfirm();
  const { enqueueSnackbar } = useSnackbar();

  const handleDelete = async () => {
    confirm({
      title: 'Delete goal',
      content: (
        <div>
          <Alert variant="filled" severity="error">
            After you have deleted a goal from your account, it will be permanently deleted.
          </Alert>
        </div>
      ),
      cancellationButtonProps: { color: 'secondary' },
      confirmationText: 'Delete',
      confirmationButtonProps: { color: 'error' },
    }).then(() => {
      client
        .deleteCarbonFootprintGoal({ id: goal.dbId }, { headers: userHeaders })
        .then(() => {
          setUserGoals((old) => {
            if (active) {
              return {
                active: old.active.filter((item) => item.dbId !== goal.dbId),
                completed: old.completed,
              };
            } else {
              return {
                active: old.active,
                completed: old.completed.filter((item) => item.dbId !== goal.dbId),
              };
            }
          });
          if (selectedUserGoal?.dbId === goal.dbId) setSelectedUserGoal(null);
          enqueueSnackbar('Deleted goal.', { variant: 'success' });
        })
        .catch((err) => {
          enqueueSnackbar('Could not delete goal.', { variant: 'error' });
          console.error(err);
        });
    });
  };

  const handleSelectGoal = () => {
    setSelectedUserGoal(goal);
  };

  return (
    <Card sx={{ minWidth: MIN_CARD_WIDTH }}>
      <CardActionArea onClick={handleSelectGoal}>
        <CardContent>
          <Grid container columns={5} direction="column">
            <Grid>
              <Typography variant="h5" component="div">
                <Typography variant="caption" color="text.secondary" display="block" gutterBottom>
                  Goal
                </Typography>
                {goal.description}
              </Typography>
            </Grid>
            <Grid container columns={2} direction="row" justifyContent="space-between">
              <Grid>
                <Typography variant="caption" color="text.secondary">
                  Start
                </Typography>
                <Typography variant="h6" color="text.primary">
                  {getDateString(goal.startDate)}
                </Typography>
              </Grid>
              <Grid>
                <Typography variant="caption" color="text.secondary">
                  Target
                </Typography>
                <Typography variant="h6" color="text.primary">
                  {getDateString(goal.endDate)}
                </Typography>
              </Grid>
            </Grid>
            <Grid container columns={2} direction="row" justifyContent="space-between">
              <Grid>
                <Typography variant="h6" color="text.primary">
                  {goal.startCarbonFootprint.toFixed(3)}
                  <Typography variant="overline">
                    CO<sub>2</sub>/kg
                  </Typography>
                </Typography>
              </Grid>
              <Grid>
                <Typography variant="h6" color="text.primary">
                  {goal.targetCarbonFootprint.toFixed(3)}
                  <Typography variant="overline">
                    CO<sub>2</sub>/kg
                  </Typography>
                </Typography>
              </Grid>
            </Grid>
          </Grid>
        </CardContent>
      </CardActionArea>
      <CardActions disableSpacing>
        <IconButton onClick={handleDelete}>
          <Delete />
        </IconButton>
      </CardActions>
    </Card>
  );
}

export default GoalCard;
