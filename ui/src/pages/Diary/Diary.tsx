import { useEffect, useMemo, useState } from 'react';
import RenderIfVisible from 'react-render-if-visible';
import { useRecoilState, useRecoilValue } from 'recoil';

import { Badge, Button, Divider, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { Stack } from '@mui/system';
import { CalendarPickerSkeleton, PickersDay } from '@mui/x-date-pickers';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';

import dayjs from 'dayjs';

import FoodItemLogCard from '@/components/FoodItemLogCard';
import { ESTIMATED_CARD_HEIGHT } from '@/components/FoodItemLogCard/FoodItemLogCard';
import LinearProgressWithLabel from '@/components/LinearProgressWithLabel';
import useHighlightedDays from '@/hooks/useHighlightedDays';
import {
  FoodItemLogDateState,
  FoodItemLogSerializableDateState,
  LocalFoodItemLogState,
  LocalFoodItemLogStats,
  LocalFoodItemLogStatsByMeal,
  LocalUserSettingsState,
  MealsState,
} from '@/store/user';
import { getDateString } from '@/utils/date';

function Diary() {
  const [isForcePickerOpen, setIsOpen] = useState(false);
  const [date, setDate] = useRecoilState(FoodItemLogDateState);
  const serializableDate = useRecoilValue(FoodItemLogSerializableDateState);
  const isToday = useMemo(() => date.isSame(dayjs(), 'date'), [date]);
  // TODO: handle errors
  const foodItemLog = useRecoilValue(LocalFoodItemLogState(serializableDate));
  const userSettings = useRecoilValue(LocalUserSettingsState);
  const stats = useRecoilValue(LocalFoodItemLogStats(serializableDate));
  const statsByMeal = useRecoilValue(LocalFoodItemLogStatsByMeal(serializableDate));
  const meals = useRecoilValue(MealsState);
  const [fetchHighlightedDays, highlightedDays, loading, requestAbortController] =
    useHighlightedDays();

  const yesterday = () => {
    setDate(date.subtract(1, 'day'));
  };

  const tommorrow = () => {
    setDate(date.add(1, 'day'));
  };

  useEffect(() => {
    fetchHighlightedDays(date);
    // abort request on unmount
    // eslint-disable-next-line react-hooks/exhaustive-deps
    return () => requestAbortController.current?.abort();
  }, [date, fetchHighlightedDays, requestAbortController]);

  const handleMonthChange = (date: dayjs.Dayjs) => {
    requestAbortController.current?.abort();
    fetchHighlightedDays(date);
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
            loading={loading}
            onMonthChange={handleMonthChange}
            renderLoading={() => <CalendarPickerSkeleton />}
            maxDate={dayjs()}
            onChange={(newValue) => {
              setDate(newValue ?? date);
            }}
            renderDay={(day, _value, DayComponentProps) => {
              const isSelected =
                !DayComponentProps.outsideCurrentMonth && highlightedDays.includes(day.date());

              return (
                <Badge
                  key={day.toString()}
                  overlap="circular"
                  badgeContent={isSelected ? '👍' : undefined}
                >
                  <PickersDay {...DayComponentProps} />
                </Badge>
              );
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
            <Grid key={i} xs={8} lg={7} xl={6}>
              <Stack direction="row" justifyContent="space-between">
                <Typography
                  sx={{ textTransform: 'capitalize' }}
                  variant="h5"
                  color="secondary.dark"
                >
                  {meal.getName()}
                </Typography>
                <Typography variant="h5" color="text.secondary">
                  <strong>{statsByMeal.get(meal.value)?.toFixed(3) ?? '0.000'}</strong>
                  <Typography variant="caption">
                    CO<sub>2</sub>/kg
                  </Typography>
                </Typography>
              </Stack>
              <Divider flexItem />
            </Grid>
            {foodItemLog.map((foodLogItem, j) => {
              return (
                foodLogItem.meal == meal.value && (
                  <Grid key={j} xs={8} lg={7} xl={6}>
                    <RenderIfVisible defaultHeight={ESTIMATED_CARD_HEIGHT}>
                      <FoodItemLogCard foodLogItem={foodLogItem} />
                    </RenderIfVisible>
                  </Grid>
                )
              );
            })}
          </>
        );
      })}
    </Grid>
  );
}

export default Diary;
