import { useMemo, useState } from 'react';
import RenderIfVisible from 'react-render-if-visible';
import { useRecoilState, useRecoilValue } from 'recoil';

import { Button, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';

import dayjs from 'dayjs';

import FoodItemLogCard from '@/components/FoodItemLogCard';
import { ESTIMATED_CARD_HEIGHT } from '@/components/FoodItemLogCard/FoodItemLogCard';
import LinearProgressWithLabel from '@/components/LinearProgressWithLabel';
import {
  FoodItemLogDateState,
  FoodItemLogSerializableDateState,
  LocalFoodItemLogState,
  LocalFoodItemLogStats,
  LocalUserSettingsState,
  MealsState,
} from '@/store/user';
import { getDateString, toSerializableDate } from '@/utils/date';

function Diary() {
  const [isForcePickerOpen, setIsOpen] = useState(false);
  const [date, setDate] = useRecoilState(FoodItemLogDateState);
  const serializableDate = useRecoilValue(FoodItemLogSerializableDateState);
  const isToday = useMemo(() => date.isSame(dayjs(), 'date'), [date]);
  // TODO: handle errors
  const foodItemLog = useRecoilValue(LocalFoodItemLogState(serializableDate));
  const userSettings = useRecoilValue(LocalUserSettingsState);
  const stats = useRecoilValue(LocalFoodItemLogStats(toSerializableDate(date)));
  const meals = useRecoilValue(MealsState);

  const yesterday = () => {
    setDate(date.subtract(1, 'day'));
  };

  const tommorrow = () => {
    setDate(date.add(1, 'day'));
  };

  return (
    <Grid
      container
      direction="column"
      columns={10}
      spacing={4}
      alignItems="center"
      pt="4vh"
      pb="8vh"
      disableEqualOverflow
    >
      <Grid container direction="row">
        <Grid>
          <Button onClick={yesterday}>Back</Button>
        </Grid>
        <Grid>
          <DatePicker
            open={isForcePickerOpen}
            onClose={() => setIsOpen(false)}
            value={date}
            maxDate={dayjs()}
            onChange={(newValue) => {
              setDate(newValue ?? date);
            }}
            renderInput={(params) => (
              <div ref={params.inputRef}>
                <input
                  disabled={params.disabled}
                  style={{ display: 'none' }}
                  value={params.inputProps?.value}
                  onChange={params.onChange}
                  {...params.inputProps}
                />

                <Button variant="text" onClick={() => setIsOpen((isOpen) => !isOpen)}>
                  {getDateString(date)}
                </Button>
              </div>
            )}
          />
        </Grid>
        <Grid>
          <Button disabled={isToday} onClick={tommorrow}>
            Next
          </Button>
        </Grid>
      </Grid>
      <Grid textAlign="center" xs={8} lg={7} xl={6}>
        <Grid
          sx={{ p: 0 }}
          container
          justifyContent="center"
          columns={5}
          spacing={2}
          direction="row"
        >
          <Grid>
            <Typography variant="h5" color="text.primary">
              <b>{userSettings.cfLimit.toFixed(3)}</b>
            </Typography>
            <Typography color="text.secondary">Goal</Typography>
          </Grid>
          <Grid>
            <Typography variant="h5" color="text.secondary">
              -
            </Typography>
          </Grid>
          <Grid>
            <Typography variant="h5" color="text.primary">
              <b>{stats.totalCarbonFootprint.toFixed(3)}</b>
            </Typography>
            <Typography color="text.secondary">Food</Typography>
          </Grid>
          <Grid>
            <Typography variant="h5" color="text.secondary">
              =
            </Typography>
          </Grid>
          <Grid>
            <Typography variant="h5" color="text.primary">
              <strong>{stats.carbonFootprintRemaining.toFixed(3)}</strong>
              <Typography variant="caption">
                CO<sub>2</sub>/kg
              </Typography>
            </Typography>
            <Typography color="text.secondary">Remaining</Typography>
          </Grid>
        </Grid>
      </Grid>
      <Grid xs={8} lg={7} xl={6}>
        <LinearProgressWithLabel value={stats.carbonFootprintGoalPercent} />
      </Grid>
      {meals.map((meal, i) => {
        return (
          <>
            <Grid key={i}>
              <Typography>{meal.getName()}</Typography>
            </Grid>
            {foodItemLog.map((foodLogItem, j) => {
              return foodLogItem.meal == meal.value && (
                <Grid key={j} xs={8} lg={7} xl={6}>
                  <RenderIfVisible defaultHeight={ESTIMATED_CARD_HEIGHT}>
                    <FoodItemLogCard foodLogItem={foodLogItem} />
                  </RenderIfVisible>
                </Grid>
              );
            })}
          </>
        );
      })}
    </Grid>
  );
}

export default Diary;
