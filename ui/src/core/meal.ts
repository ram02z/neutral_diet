import BreakfastDiningIcon from '@mui/icons-material/BreakfastDining';
import DinnerDiningIcon from '@mui/icons-material/DinnerDining';
import LunchDiningIcon from '@mui/icons-material/LunchDining';
import TapasIcon from '@mui/icons-material/Tapas';

import { Meal as MealProto } from '@/api/gen/neutral_diet/user/v1/food_item_log_pb';

export class Meal {
  value: MealProto;

  constructor(meal: MealProto) {
    this.value = meal;
  }

  getIcon() {
    switch (this.value) {
      case MealProto.UNSPECIFIED:
        return TapasIcon;
      case MealProto.BREAKFAST:
        return BreakfastDiningIcon;
      case MealProto.LUNCH:
        return LunchDiningIcon;
      case MealProto.DINNER:
        return DinnerDiningIcon;
    }
  }

  getName(): string {
    switch (this.value) {
      case MealProto.UNSPECIFIED:
        return 'snacks';
      case MealProto.BREAKFAST:
        return 'breakfast';
      case MealProto.LUNCH:
        return 'lunch';
      case MealProto.DINNER:
        return 'dinner';
    }
  }

  getOrder(): number {
    switch (this.value) {
      case MealProto.UNSPECIFIED:
        return 4;
      case MealProto.BREAKFAST:
        return 1;
      case MealProto.LUNCH:
        return 2;
      case MealProto.DINNER:
        return 3;
    }
  }
}
