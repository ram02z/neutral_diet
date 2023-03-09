import { Meal as MealProto } from '@/api/gen/neutral_diet/user/v1/food_item_log_pb';

export class Meal {
  value: MealProto;

  constructor(meal: MealProto) {
    this.value = meal;
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
}
