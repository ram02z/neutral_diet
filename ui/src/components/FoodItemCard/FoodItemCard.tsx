import { useMemo, useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil';

import { Add, Info } from '@mui/icons-material';
import { Box, Card, CardActions, CardContent, IconButton, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { Stack } from '@mui/system';

import { useSnackbar } from 'notistack';

import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import client from '@/api/user_service';
import AddFoodItemDialog from '@/components/AddFoodItemDialog';
import FoodItemInfoDialog from '@/components/FoodItemInfoDialog';
import RegionChip from '@/components/RegionChip';
import { MIN_CARD_WIDTH } from '@/config';
import UserRegion from '@/core/regions';
import { WeightUnit } from '@/core/weight';
import { FoodHistoryState, FoodItemInfoQuery } from '@/store/food';
import { CurrentUserHeadersState, FoodItemLogDateState, FoodItemLogSerializableDateState, LocalFoodItemLogState } from '@/store/user';

import { FormValues } from './types';

export const ESTIMATED_CARD_HEIGHT = 160;

type FoodItemCardProps = {
  foodItem: AggregateFoodItem;
};

export function FoodItemCard({ foodItem }: FoodItemCardProps) {
  const [foodHistory, setFoodHistory] = useRecoilState(FoodHistoryState);
  const setDate = useSetRecoilState(FoodItemLogDateState);
  const serializableDate = useRecoilValue(FoodItemLogSerializableDateState);
  const [, setFoodItemLog] = useRecoilState(LocalFoodItemLogState(serializableDate));
  const [inHistory, setInHistory] = useState(false);
  const [openAddDialog, setOpenAddDialog] = useState(false);
  const [openInfoDialog, setOpenInfoDialog] = useState(false);
  const { enqueueSnackbar } = useSnackbar();
  const userHeaders = useRecoilValue(CurrentUserHeadersState);
  const foodItemInfo = useRecoilValue(
    FoodItemInfoQuery({ foodItemId: foodItem.id, region: foodItem.region }),
  );
  const region = new UserRegion(foodItem.region);

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    setDate(data.date);
    const weight = parseFloat(data.weight);
    const weightUnit = new WeightUnit(data.weightUnit);
    client
      .addFoodItem(
        {
          foodLogItem: {
            foodItemId: foodItem.id,
            weight: weight,
            weightUnit: data.weightUnit,
            date: { year: data.date.year(), month: data.date.month() + 1, day: data.date.date() },
            region: foodItem.region,
            meal: data.meal,
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
              weight: { value: weight, unit: weightUnit },
              carbonFootprint: res.carbonFootprint,
              region: foodItem.region,
              meal: data.meal,
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
    <Card sx={{ minWidth: MIN_CARD_WIDTH }}>
      <CardContent>
        <Grid container columns={5}>
          <Stack spacing={1} sx={{ pt: 1, pl: 1 }}>
            <Typography sx={{ textTransform: 'capitalize' }} variant="h5" component="div">
              {foodItem.foodName.toLowerCase()}
            </Typography>
            <Typography variant="subtitle1" color="text.secondary">
              <b>{parseFloat(foodItem.medianCarbonFootprint.toFixed(3))}</b>
              <Typography variant="caption">
                CO<sub>2</sub>/kg
              </Typography>
            </Typography>
          </Stack>
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
        <Box sx={{ marginLeft: 'auto' }}>
          <RegionChip region={region} />
        </Box>
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
