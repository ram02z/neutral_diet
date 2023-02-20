import dayjs from 'dayjs';

export type FormValues = {
  date: dayjs.Dayjs;
  weight: string;
  weightUnit: string;
};
