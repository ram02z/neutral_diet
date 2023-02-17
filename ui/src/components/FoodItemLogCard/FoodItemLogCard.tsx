import { useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useRecoilValue, useSetRecoilState } from 'recoil';

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
import EditFoodItemDialog from '@/components/EditFoodItemDialog';
import { FormValues } from '@/components/FoodItemCard/types';
import { MIN_WIDTH } from '@/config';
import { CurrentUserHeadersState, FoodItemLogDateState, LocalFoodItemLogState } from '@/store/user';
import { LocalFoodLogItem } from '@/store/user/types';

type FoodItemCardProps = {
  foodLogItem: LocalFoodLogItem;
};

function FoodItemLogCard({ foodLogItem }: FoodItemCardProps) {
  const confirm = useConfirm();
  const [openDialog, setOpenDialog] = useState(false);
  const userHeaders = useRecoilValue(CurrentUserHeadersState);
  const date = useRecoilValue(FoodItemLogDateState);
  const { enqueueSnackbar } = useSnackbar();
  const setFoodItemLog = useSetRecoilState(LocalFoodItemLogState(date));

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    const weight = parseFloat(data.weight);
    client
      .updateFoodItem(
        {
          id: foodLogItem.dbId,
          weight: weight,
        },
        { headers: userHeaders },
      )
      .then((res) => {
        setFoodItemLog((old) => {
          return old.map((item) => {
            if (item.dbId == foodLogItem.dbId) {
              return {
                dbId: foodLogItem.dbId,
                name: foodLogItem.name,
                weight: weight,
                carbonFootprint: res.carbonFootprint,
              };
            }
            return { ...item };
          });
        });
        enqueueSnackbar('Updated food entry', { variant: 'success' });
      })
      .catch((err) => {
        enqueueSnackbar("Couldn't update food entry", { variant: 'error' });
        console.error(err);
      });
    handleClose();
  };

  const handleDelete = async () => {
    confirm({
      title: 'Delete food item',
      content: (
        <div>
          <Alert severity="error">
            After you have deleted a food item from your log, it will be permanently deleted.
          </Alert>
        </div>
      ),
      cancellationButtonProps: { color: 'info' },
      confirmationText: 'Delete',
      confirmationButtonProps: { color: 'error', variant: 'contained' },
    }).then(() => {
      client
        .deleteFoodItem({ id: foodLogItem.dbId }, { headers: userHeaders })
        .then(() => {
          setFoodItemLog((old) => old.filter((item) => item.dbId !== foodLogItem.dbId));
          enqueueSnackbar('Deleted food entry.', { variant: 'success' });
        })
        .catch((err) => {
          enqueueSnackbar('Could not delete food entry.', { variant: 'error' });
          console.error(err);
        });
    });
  };

  const handleClickOpen = () => {
    setOpenDialog(true);
  };

  const handleClose = () => {
    setOpenDialog(false);
  };

  return (
    <Card sx={{ minWidth: MIN_WIDTH }}>
      <CardActionArea onClick={handleClickOpen}>
        <CardContent>
          <Grid container columns={5}>
            <Grid xs={4} sx={{ pt: 1, pl: 1 }}>
              <Typography sx={{ textTransform: 'capitalize' }} variant="h5" component="div">
                {foodLogItem.name.toLowerCase()}
              </Typography>
              <Typography variant="subtitle1" color="text.secondary" component="div">
                {foodLogItem.carbonFootprint}
              </Typography>
            </Grid>
          </Grid>
        </CardContent>
      </CardActionArea>
      <CardActions disableSpacing>
        <IconButton onClick={handleDelete}>
          <Delete />
        </IconButton>
      </CardActions>
      <EditFoodItemDialog
        onSubmit={onSubmit}
        openDialog={openDialog}
        handleClose={handleClose}
        currentWeight={foodLogItem.weight.toString()}
      />
    </Card>
  );
}

export default FoodItemLogCard;
