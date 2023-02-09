import { useMemo, useState } from 'react';

import { Button } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';

import dayjs from 'dayjs';

import { getDateString } from '@/utils/date';

function Diary() {
  const [isForcePickerOpen, setIsOpen] = useState(false);
  const [date, setDate] = useState(dayjs());
  const isToday = useMemo(() => date.isSame(dayjs(), 'date'), [date]);

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
          <LocalizationProvider dateAdapter={AdapterDayjs}>
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
          </LocalizationProvider>
        </Grid>
        <Grid>
          <Button disabled={isToday} onClick={tommorrow}>
            Next
          </Button>
        </Grid>
      </Grid>
    </Grid>
  );
}

export default Diary;
