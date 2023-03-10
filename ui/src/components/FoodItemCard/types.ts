import dayjs from 'dayjs';

export type FormValues = {
  date: dayjs.Dayjs;
  quantity: string;
  unit: number;
  meal: number;
};
