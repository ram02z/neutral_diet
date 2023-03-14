import { useRecoilValue, useSetRecoilState } from 'recoil';

import { Delete } from '@mui/icons-material';
import {
  Alert,
  Box,
  Card,
  CardActionArea,
  CardActions,
  CardContent,
  Checkbox,
  IconButton,
  Typography,
} from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import { useConfirm } from 'material-ui-confirm';
import { useSnackbar } from 'notistack';

import client from '@/api/user_service';
import { MIN_CARD_WIDTH } from '@/config';
import { CurrentUserHeadersState, UserGoalsState } from '@/store/user';
import { LocalUserGoal } from '@/store/user/types';

type GoalCardProps = {
  goal: LocalUserGoal;
  active: boolean;
};

function GoalCard({ goal, active }: GoalCardProps) {
  const setUserGoals = useSetRecoilState(UserGoalsState);
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
          enqueueSnackbar('Deleted food entry.', { variant: 'success' });
        })
        .catch((err) => {
          enqueueSnackbar('Could not delete food entry.', { variant: 'error' });
          console.error(err);
        });
    });
  };

  const handleCheck = async () => {
    client
      .updateCarbonFootprintGoal({ id: goal.dbId, completed: !active }, { headers: userHeaders })
      .then(() => {
        setUserGoals((old) => {
          if (active) {
            return {
              active: old.active.filter((item) => item.dbId !== goal.dbId),
              completed: [...old.completed, goal],
            };
          } else {
            return {
              active: [...old.active, goal],
              completed: old.completed.filter((item) => item.dbId !== goal.dbId),
            };
          }
        });
        enqueueSnackbar(`Marked goal as ${active ? 'completed' : 'active'}.`, {
          variant: 'success',
        });
      })
      .catch((err) => {
        enqueueSnackbar(`Could not move goal to ${active ? 'completed' : 'active'}.`, {
          variant: 'error',
        });
        console.error(err);
      });
  };

  return (
    <Card sx={{ minWidth: MIN_CARD_WIDTH }}>
      <CardActionArea>
        <CardContent>
          <Grid container columns={5} direction="column">
            <Grid>
              <Typography variant="h5" component="div">
                <Typography variant="caption" color="text.secondary" display="block" gutterBottom>
                  Description
                </Typography>
                {goal.description}
              </Typography>
            </Grid>
            <Grid>
              <Typography variant="h6" color="text.primary">
                Start: {goal.startCarbonFootprint.toFixed(3)}
                <Typography variant="overline">
                  CO<sub>2</sub>/kg
                </Typography>
              </Typography>
              <Typography variant="h6" color="text.primary">
                Target: {goal.targetCarbonFootprint.toFixed(3)}
                <Typography variant="overline">
                  CO<sub>2</sub>/kg
                </Typography>
              </Typography>
            </Grid>
            <Grid>
              <Typography variant="h6" color="text.primary">
                {goal.endDate?.day}/{goal.endDate?.month}/{goal.endDate?.year}
              </Typography>
            </Grid>
          </Grid>
        </CardContent>
      </CardActionArea>
      <CardActions disableSpacing>
        <Box sx={{ marginRight: 'auto' }}>
          <IconButton onClick={handleDelete}>
            <Delete />
          </IconButton>
        </Box>
        <Box sx={{ marginLeft: 'auto' }}>
          <Checkbox
            onChange={handleCheck}
            checked={!active}
            size="medium"
            inputProps={{ 'aria-label': 'controlled' }}
          />
        </Box>
      </CardActions>
    </Card>
  );
}

export default GoalCard;
