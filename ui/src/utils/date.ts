import dayjs from 'dayjs';
import isToday from 'dayjs/plugin/isToday';
import isYesterday from 'dayjs/plugin/isYesterday';

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
