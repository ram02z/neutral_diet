import { useMemo, useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil';

import { Add, Info } from '@mui/icons-material';
import { Card, CardActions, CardContent, IconButton, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import { useSnackbar } from 'notistack';

import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import client from '@/api/user_service';
import AddFoodItemDialog from '@/components/AddFoodItemDialog';
import FoodItemInfoDialog from '@/components/FoodItemInfoDialog';
import { MIN_WIDTH } from '@/config';
import { ReverseWeightUnitNameMap, Weight } from '@/core/weight';
import { FoodHistoryState, FoodItemInfoQuery } from '@/store/food';
import { CurrentUserHeadersState, FoodItemLogDateState, LocalFoodItemLogState } from '@/store/user';

import { FormValues } from './types';

export const ESTIMATED_CARD_HEIGHT = 160;

type FoodItemCardProps = {
  foodItem: AggregateFoodItem;
};

export function FoodItemCard({ foodItem }: FoodItemCardProps) {
  const [foodHistory, setFoodHistory] = useRecoilState(FoodHistoryState);
  const [date, setDate] = useRecoilState(FoodItemLogDateState);
  const setFoodItemLog = useSetRecoilState(LocalFoodItemLogState(date));
  const [inHistory, setInHistory] = useState(false);
  const [openAddDialog, setOpenAddDialog] = useState(false);
  const [openInfoDialog, setOpenInfoDialog] = useState(false);
  const { enqueueSnackbar } = useSnackbar();
  const userHeaders = useRecoilValue(CurrentUserHeadersState);
  const foodItemInfo = useRecoilValue(FoodItemInfoQuery(foodItem.id));

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    setDate(data.date);
    const weight = parseFloat(data.weight);
    const weightUnit = ReverseWeightUnitNameMap.get(data.weightUnit);
    client
      .addFoodItem(
        {
          foodLogItem: {
            foodItemId: foodItem.id,
            weight: weight,
            weightUnit: weightUnit,
            date: { year: data.date.year(), month: data.date.month() + 1, day: data.date.date() },
          },
        },
        { headers: userHeaders },
      )
      .then((res) => {
        if (!inHistory) {
          setFoodHistory((oldFoodHistory) => [...oldFoodHistory, foodItem]);
        }
        setFoodItemLog((old) => {
          return [
            ...old,
            {
              dbId: res.id,
              foodItemId: foodItem.id,
              name: foodItem.foodName,
              weight: new Weight(weight, weightUnit),
              carbonFootprint: res.carbonFootprint,
            },
          ];
        });
        enqueueSnackbar('Added food to diary', { variant: 'success' });
      })
      .catch((err) => {
        enqueueSnackbar("Couldn't add food", { variant: 'error' });
        console.error(err);
      });
    handleCloseAddDialog();
  };

  const handleOpenAddDialog = () => {
    setOpenAddDialog(true);
  };

  const handleCloseAddDialog = () => {
    setOpenAddDialog(false);
  };

  const handleOpenInfoDialog = () => {
    setOpenInfoDialog(true);
  };

  const handleCloseInfoDialog = () => {
    setOpenInfoDialog(false);
  };

  useMemo(() => {
    setInHistory(foodHistory.some((value) => value.id === foodItem.id));
  }, [foodHistory, foodItem]);

  return (
    <Card sx={{ minWidth: MIN_WIDTH }}>
      <CardContent>
        <Grid container columns={5}>
          <Grid xs={4} sx={{ pt: 1, pl: 1 }}>
            <Typography sx={{ textTransform: 'capitalize' }} variant="h5" component="div">
              {foodItem.foodName.toLowerCase()}
            </Typography>
            <Typography variant="subtitle1" color="text.secondary">
              <b>{parseFloat(foodItem.medianCarbonFootprint.toFixed(3))}</b>
              <Typography variant="caption">
                CO<sub>2</sub>/kg
              </Typography>
            </Typography>
          </Grid>
          <Grid
            xs
            display="flex"
            justifyContent="flex-end"
            alignItems="center"
            sx={{ pt: 1, pl: 1 }}
          >
            <IconButton onClick={handleOpenAddDialog} aria-label="add" sx={{ float: 'right' }}>
              <Add />
            </IconButton>
          </Grid>
        </Grid>
      </CardContent>
      <CardActions disableSpacing>
        <IconButton onClick={handleOpenInfoDialog} aria-label="info" disabled={!foodItemInfo}>
          <Info />
        </IconButton>
      </CardActions>
      <AddFoodItemDialog
        onSubmit={onSubmit}
        openDialog={openAddDialog}
        handleClose={handleCloseAddDialog}
      />
      <FoodItemInfoDialog
        openDialog={openInfoDialog}
        handleClose={handleCloseInfoDialog}
        foodItemInfo={foodItemInfo}
      />
    </Card>
  );
}

export default FoodItemCard;
