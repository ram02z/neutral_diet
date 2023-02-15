import { useMemo, useState } from 'react';
import { useRecoilState, useRecoilValue } from 'recoil';

import { Button } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';

import dayjs from 'dayjs';

import FoodItemLogCard from '@/components/FoodItemLogCard';
import { FoodItemLogDateState, FoodItemLogQuery } from '@/store/user';
import { getDateString } from '@/utils/date';

function Diary() {
  const [isForcePickerOpen, setIsOpen] = useState(false);
  const [date, setDate] = useRecoilState(FoodItemLogDateState);
  const isToday = useMemo(() => date.isSame(dayjs(), 'date'), [date]);
  // TODO: handle errors
  const foodItemLog = useRecoilValue(FoodItemLogQuery);

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
      {foodItemLog.map((foodLogItem, idx) => (
        <Grid key={idx} xs={8} sm={7} md={6} lg={5} xl={4}>
          <FoodItemLogCard foodLogItem={foodLogItem} />
        </Grid>
      ))}
    </Grid>
  );
}

export default Diary;
