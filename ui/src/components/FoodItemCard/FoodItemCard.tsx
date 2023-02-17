import { useMemo, useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil';

import { Add } from '@mui/icons-material';
import { Card, CardContent, IconButton, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import { useSnackbar } from 'notistack';

import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import client from '@/api/user_service';
import AddFoodItemDialog from '@/components/AddFoodItemDialog';
import { MIN_WIDTH } from '@/config';
import { FoodHistoryState } from '@/store/food';
import { CurrentUserHeadersState, FoodItemLogDateState, LocalFoodItemLogState } from '@/store/user';

import { FormValues } from './types';

type FoodItemCardProps = {
  foodItem: AggregateFoodItem;
};

function FoodItemCard({ foodItem }: FoodItemCardProps) {
  const [foodHistory, setFoodHistory] = useRecoilState(FoodHistoryState);
  const [date, setDate] = useRecoilState(FoodItemLogDateState);
  const setFoodItemLog = useSetRecoilState(LocalFoodItemLogState(date));
  const [inHistory, setInHistory] = useState(false);
  const [openDialog, setOpenDialog] = useState(false);
  const { enqueueSnackbar } = useSnackbar();
  const userHeaders = useRecoilValue(CurrentUserHeadersState);

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    setDate(data.date);
    const weight = parseFloat(data.weight);
    client
      .addFoodItem(
        {
          foodLogItem: {
            foodItemId: foodItem.id,
            weight: weight,
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
              name: foodItem.foodName,
              weight: weight,
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
    handleClose();
  };

  const handleClickOpen = () => {
    setOpenDialog(true);
  };

  const handleClose = () => {
    setOpenDialog(false);
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
            <Typography variant="subtitle1" color="text.secondary" component="div">
              {foodItem.medianCarbonFootprint}
            </Typography>
          </Grid>
          <Grid
            xs
            display="flex"
            justifyContent="flex-end"
            alignItems="center"
            sx={{ pt: 1, pl: 1 }}
          >
            <IconButton onClick={handleClickOpen} aria-label="add" sx={{ float: 'right' }}>
              <Add />
            </IconButton>
          </Grid>
        </Grid>
      </CardContent>
      <AddFoodItemDialog onSubmit={onSubmit} openDialog={openDialog} handleClose={handleClose} />
    </Card>
  );
}

export default FoodItemCard;
