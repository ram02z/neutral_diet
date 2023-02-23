import { useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useRecoilValue, useSetRecoilState } from 'recoil';

import { Delete, Info } from '@mui/icons-material';
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
import FoodItemInfoDialog from '@/components/FoodItemInfoDialog';
import { MIN_WIDTH } from '@/config';
import { ReverseWeightUnitNameMap, Weight } from '@/core/weight';
import { CurrentUserHeadersState, FoodItemLogDateState, LocalFoodItemLogState } from '@/store/user';
import { LocalFoodLogItem } from '@/store/user/types';
import { FoodItemInfoQuery } from '@/store/food';

export const ESTIMATED_CARD_HEIGHT = 160;

type FoodItemCardProps = {
  foodLogItem: LocalFoodLogItem;
};

export function FoodItemLogCard({ foodLogItem }: FoodItemCardProps) {
  const confirm = useConfirm();
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false);
  const [openInfoDialog, setOpenInfoDialog] = useState(false);
  const userHeaders = useRecoilValue(CurrentUserHeadersState);
  const date = useRecoilValue(FoodItemLogDateState);
  const { enqueueSnackbar } = useSnackbar();
  const setFoodItemLog = useSetRecoilState(LocalFoodItemLogState(date));
  const foodItemInfo = useRecoilValue(FoodItemInfoQuery(foodLogItem.foodItemId))

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    const weight = parseFloat(data.weight);
    const weightUnit = ReverseWeightUnitNameMap.get(data.weightUnit);
    client
      .updateFoodItem(
        {
          id: foodLogItem.dbId,
          weight: weight,
          weightUnit: weightUnit,
        },
        { headers: userHeaders },
      )
      .then((res) => {
        setFoodItemLog((old) => {
          return old.map((item) => {
            if (item.dbId == foodLogItem.dbId) {
              return {
                dbId: foodLogItem.dbId,
                foodItemId: foodLogItem.foodItemId,
                name: foodLogItem.name,
                weight: new Weight(weight, weightUnit),
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
    handleCloseDeleteDialog();
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

  const handleOpenDeleteDialog = () => {
    setOpenDeleteDialog(true);
  };

  const handleCloseDeleteDialog = () => {
    setOpenDeleteDialog(false);
  };

  const handleOpenInfoDialog = () => {
    setOpenInfoDialog(true);
  };

  const handleCloseInfoDialog = () => {
    setOpenInfoDialog(false);
  };

  return (
    <Card sx={{ minWidth: MIN_WIDTH }}>
      <CardActionArea onClick={handleOpenDeleteDialog}>
        <CardContent>
          <Grid container columns={5}>
            <Grid xs={4} sx={{ pt: 1, pl: 1 }}>
              <Typography sx={{ textTransform: 'capitalize' }} variant="h5" component="div">
                {foodLogItem.name.toLowerCase()}
              </Typography>
              <Typography variant="subtitle1" color="text.secondary" component="div">
                {foodLogItem.weight.getFormattedName()}
              </Typography>
            </Grid>
            <Grid
              xs
              display="flex"
              justifyContent="flex-end"
              alignItems="center"
              textAlign="center"
              sx={{ pt: 1, pl: 1 }}
            >
              <Typography variant="h6" color="text.primary">
                {parseFloat(foodLogItem.carbonFootprint.toFixed(3))}
              </Typography>
              <Typography variant="overline">
                CO<sub>2</sub>/kg
              </Typography>
            </Grid>
          </Grid>
        </CardContent>
      </CardActionArea>
      <CardActions disableSpacing>
        <IconButton
          onClick={handleOpenInfoDialog}
          aria-label="info"
          disabled={!foodItemInfo}
        >
          <Info />
        </IconButton>
        <IconButton onClick={handleDelete}>
          <Delete />
        </IconButton>
      </CardActions>
      <EditFoodItemDialog
        onSubmit={onSubmit}
        openDialog={openDeleteDialog}
        handleClose={handleCloseDeleteDialog}
        currentWeight={foodLogItem.weight}
      />
      <FoodItemInfoDialog
        openDialog={openInfoDialog}
        handleClose={handleCloseInfoDialog}
        foodItemInfo={foodItemInfo}
      />
    </Card>
  );
}
