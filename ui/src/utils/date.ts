import dayjs from 'dayjs';
import isToday from 'dayjs/plugin/isToday';
import isYesterday from 'dayjs/plugin/isYesterday';

import { SerializableDate } from '@/store/user/types';

export function getDateString(date: dayjs.Dayjs): string {
  dayjs.extend(isToday);
  dayjs.extend(isYesterday);

  if (date.isToday()) {
    return 'Today';
  } else if (date.isYesterday()) {
    return 'Yesterday';
  } else if (dayjs().isSame(date, 'year')) {
    return date.format('dddd, MMM D');
  } else {
    return date.format('dddd, MMM D YYYY');
  }
}

export function toSerializableDate(date: dayjs.Dayjs): SerializableDate {
  return { year: date.year(), month: date.month() + 1, day: date.date() };
}
