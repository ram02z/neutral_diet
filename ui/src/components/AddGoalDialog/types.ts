import dayjs from "dayjs";

export type FormValues = {
  description: string;
  endDate: dayjs.Dayjs;
  targetCarbonFootprint: string;
};
