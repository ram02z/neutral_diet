import { WeightUnit as WeightUnitProto } from '@/api/gen/neutral_diet/user/v1/food_item_log_pb';

export type Weight = {
  value: number;
  unit: WeightUnit;
};

export class WeightUnit {
  value: WeightUnitProto;

  constructor(weightUnit: WeightUnitProto) {
    this.value = weightUnit;
  }

  getName(): string {
    switch (this.value) {
      case WeightUnitProto.GRAM:
        return 'gram';
      case WeightUnitProto.OUNCE:
        return 'ounce';
      case WeightUnitProto.POUND:
        return 'pound';
      case WeightUnitProto.UNSPECIFIED:
        return 'kilogram';
    }
  }

  getShortName(): string {
    switch (this.value) {
      case WeightUnitProto.GRAM:
        return 'g';
      case WeightUnitProto.OUNCE:
        return 'oz';
      case WeightUnitProto.POUND:
        return 'lb';
      case WeightUnitProto.UNSPECIFIED:
        return 'kg';
    }
  }
}
