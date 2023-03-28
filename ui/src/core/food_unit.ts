import { Unit } from '@/api/gen/neutral_diet/user/v1/food_item_log_pb';

export type QuantityUnit = {
  value: number;
  unit: FoodUnit;
};

export class FoodUnit {
  value: Unit;

  constructor(unit: Unit) {
    this.value = unit;
  }

  getName(): string {
    switch (this.value) {
      case Unit.UNSPECIFIED:
        return 'kilogram';
      case Unit.GRAM:
        return 'gram';
      case Unit.OUNCE:
        return 'ounce';
      case Unit.POUND:
        return 'pound';
      case Unit.PINT:
        return 'pint';
      case Unit.GALLON:
        return 'gallon';
      case Unit.MILLILITRE:
        return 'millilitre';
      case Unit.LITRE:
        return 'litre';
    }
  }

  getShortName(): string {
    switch (this.value) {
      case Unit.UNSPECIFIED:
        return 'kg';
      case Unit.GRAM:
        return 'g';
      case Unit.OUNCE:
        return 'oz';
      case Unit.POUND:
        return 'lb';
      case Unit.PINT:
        return 'pt';
      case Unit.GALLON:
        return 'gal';
      case Unit.MILLILITRE:
        return 'ml';
      case Unit.LITRE:
        return 'l';
    }
  }
}
