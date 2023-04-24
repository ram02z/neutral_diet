import dayjs from 'dayjs';
import isToday from 'dayjs/plugin/isToday';
import isYesterday from 'dayjs/plugin/isYesterday';

import { SerializableDate } from '@/store/user/types';

/**
 * Returns a formatted string representation of the current date.
 * If the date is today, it returns 'Today'.
 * If the date is yesterday, it returns 'Yesterday'.
 * If the date is from the same year as the current date, it returns the day of the week, month and day.
 * Otherwise, it returns the day of the week, month, day and year.
 */
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

export function toDayJsDate(date: SerializableDate): dayjs.Dayjs {
  const d = new Date(date.year, date.month - 1, date.day);
  return dayjs(d);
}
