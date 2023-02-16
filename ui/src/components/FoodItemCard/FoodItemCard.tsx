import { useMemo, useState } from 'react';
import { Controller, SubmitHandler, useForm } from 'react-hook-form';
import { useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil';

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
import { DatePicker } from '@mui/x-date-pickers/DatePicker';

import dayjs from 'dayjs';
import { useSnackbar } from 'notistack';

import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import { ID_TOKEN_HEADER } from '@/api/transport';
import client from '@/api/user_service';
import { MIN_WIDTH } from '@/config';
import { FoodHistoryState } from '@/store/food';
import { CurrentUserTokenIDState, FoodItemLogDateState, LocalFoodItemLogState } from '@/store/user';

type FoodItemCardProps = {
  foodItem: AggregateFoodItem;
};

type FormValues = {
  date: dayjs.Dayjs;
  weight: string;
};

function FoodItemCard({ foodItem }: FoodItemCardProps) {
  const [foodHistory, setFoodHistory] = useRecoilState(FoodHistoryState);
  const [date, setDate] = useRecoilState(FoodItemLogDateState);

  const setFoodItemLog = useSetRecoilState(LocalFoodItemLogState(date));
  const [inHistory, setInHistory] = useState(false);
  const [openDialog, setOpenDialog] = useState(false);
  const { enqueueSnackbar } = useSnackbar();
  const idToken = useRecoilValue(CurrentUserTokenIDState);
  const { handleSubmit, control } = useForm<FormValues>();

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    setDate(data.date);
    const weight = parseFloat(data.weight);
    const carbonFootprint = weight * foodItem.medianCarbonFootprint;
    if (idToken) {
      const headers = new Headers();
      headers.set(ID_TOKEN_HEADER, idToken);
      client
        .addFoodItem(
          {
            foodLogItem: {
              foodItemId: foodItem.id,
              weight: weight,
              carbonFootprint: carbonFootprint,
              date: { year: data.date.year(), month: data.date.month() + 1, day: data.date.date() },
            },
          },
          { headers: headers },
        )
        .then((res) => {
          if (!inHistory) {
            setFoodHistory((oldFoodHistory) => [...oldFoodHistory, foodItem]);
          }
          setFoodItemLog((old) => {
            return [
              ...old,
              {
                remoteId: res.id,
                name: foodItem.foodName,
                weight: weight,
                carbonFootprint: carbonFootprint,
              },
            ];
          });
          enqueueSnackbar('Added food to diary', { variant: 'success' });
        })
        .catch((err) => {
          enqueueSnackbar("Couldn't add food", { variant: 'error' });
          console.error(err);
        });
    }
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
      <Dialog open={openDialog} onClose={handleClose}>
        <form onSubmit={handleSubmit(onSubmit)}>
          <DialogTitle textAlign="center">Add food</DialogTitle>
          <Stack sx={{ px: 5 }} spacing={3}>
            <Controller
              control={control}
              name="date"
              defaultValue={dayjs()}
              rules={{
                required: true,
              }}
              render={({ field: { ref, onChange, value }, fieldState: { error } }) => (
                <DatePicker
                  value={value}
                  onChange={onChange}
                  inputRef={ref}
                  label="Date"
                  inputFormat="YYYY-MM-DD"
                  maxDate={dayjs()}
                  renderInput={(inputProps) => (
                    <TextField {...inputProps} error={!!error} helperText={error?.message} />
                  )}
                />
              )}
            />
            <Controller
              control={control}
              name="weight"
              defaultValue="1.0"
              rules={{ required: true, min: 0.001 }}
              render={({ field: { onChange, value }, fieldState: { error } }) => (
                <TextField
                  error={!!error}
                  helperText={error?.message}
                  onChange={onChange}
                  value={value}
                  type="number"
                  label="Weight"
                  inputProps={{ step: 0.001 }}
                />
              )}
            />
          </Stack>
          <DialogActions>
            <Button onClick={handleClose}>Cancel</Button>
            <Button variant="contained" type="submit">
              Add
            </Button>
          </DialogActions>
        </form>
      </Dialog>
    </Card>
  );
}

export default FoodItemCard;
