import { useMemo, useState } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { useRecoilState, useRecoilValue } from 'recoil';

import { Add } from '@mui/icons-material';
import {
  Button,
  Card,
  CardContent,
  Dialog,
  DialogActions,
  DialogTitle,
  IconButton,
  TextField,
  Typography,
} from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { Stack } from '@mui/system';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';

import dayjs from 'dayjs';

import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import { ID_TOKEN_HEADER } from '@/api/transport';
import client from '@/api/user_service';
import { MIN_WIDTH } from '@/config';
import { FoodHistoryState } from '@/store/food';
import { CurrentUserTokenIDState } from '@/store/user';

type FoodItemCardProps = {
  foodItem: AggregateFoodItem;
};

type AddFoodItemProps = {
  date: dayjs.Dayjs;
  weight: number;
};

function FoodItemCard({ foodItem }: FoodItemCardProps) {
  const [foodHistory, setFoodHistory] = useRecoilState(FoodHistoryState);
  const [inHistory, setInHistory] = useState(false);
  const [openDialog, setOpenDialog] = useState(false);
  const [date, setDate] = useState<dayjs.Dayjs | null>(dayjs());
  const idToken = useRecoilValue(CurrentUserTokenIDState);
  const {
    register,
    formState: { errors },
    handleSubmit,
  } = useForm();

  const onSubmit = (data: AddFoodItemProps) => {
    if (idToken) {
      const headers = new Headers();
      headers.set(ID_TOKEN_HEADER, idToken);
      client
        .addFoodItem(
          {
            foodLogItem: {
              foodItemId: foodItem.id,
              weight: data.weight,
              carbonFootprint: data.weight * foodItem.medianCarbonFootprint,
              date: { year: data.date.year(), month: data.date.month() + 1, day: data.date.date() },
            },
          },
          { headers: headers },
        )
        .catch((err) => console.error(err));
    }
    handleClose();
  };

  const handleClickOpen = () => {
    setOpenDialog(true);
  };

  const handleClose = () => {
    setOpenDialog(false);
  };

  const handleChangeDate = (newDate: dayjs.Dayjs | null) => {
    setDate(newDate);
  };

  useMemo(() => {
    setInHistory(foodHistory.some((value) => value.id === foodItem.id));
  }, [foodHistory, foodItem]);

  // const addToHistory = () => {
  //   if (inHistory) return;
  //   setFoodHistory((oldFoodHistory) => [...oldFoodHistory, foodItem]);
  // };

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
      <Dialog open={openDialog} onClose={handleClose}>
        <form onSubmit={handleSubmit(onSubmit)}>
          <DialogTitle textAlign="center">Add food</DialogTitle>
          <Stack sx={{ px: 5 }} spacing={3}>
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <DatePicker
                {...register('date', {
                  required: true,
                  value: date,
                  setValueAs: () => date,
                })}
                label="Date"
                inputFormat="YYYY-MM-DD"
                value={date}
                maxDate={dayjs()}
                onChange={handleChangeDate}
                renderInput={(params) => <TextField {...params} />}
              />
            </LocalizationProvider>
            <TextField
              {...register('weight', {
                required: true,
                min: 0.001,
                setValueAs: (value) => parseFloat(value),
              })}
              inputProps={{ step: 0.01 }}
              type="number"
              label="Weight"
              error={errors.weight ? true : false}
            />
          </Stack>
          <DialogActions>
            <Button onClick={handleClose}>Cancel</Button>
            <Button type="submit">Add</Button>
          </DialogActions>
        </form>
      </Dialog>
    </Card>
  );
}

export default FoodItemCard;
