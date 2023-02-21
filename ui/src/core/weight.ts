import { WeightUnit } from '@/api/gen/neutral_diet/user/v1/food_item_log_pb';

export const DEFAULT_WEIGHT_UNIT = 'kilogram';

export const WeightUnitNameMap = new Map<number, string>([
  [WeightUnit.KILOGRAM, 'kilogram'],
  [WeightUnit.GRAM, 'gram'],
  [WeightUnit.OUNCE, 'ounce'],
  [WeightUnit.POUND, 'pound'],
]);

export const ShortWeightUnitNameMap = new Map<number, string>([
  [WeightUnit.KILOGRAM, 'kg'],
  [WeightUnit.GRAM, 'g'],
  [WeightUnit.OUNCE, 'oz'],
  [WeightUnit.POUND, 'lb'],
])

export const ReverseWeightUnitNameMap = new Map([...WeightUnitNameMap].map(([k, v]) => [v, k]));

export class Weight {
  value: number;
  weightUnit: WeightUnit;

  constructor(value: number, weightUnit: WeightUnit | undefined) {
    this.value = value;
    this.weightUnit = weightUnit ?? WeightUnit.KILOGRAM;
  }

  getFormattedName(): string {
    return `${this.value}${this.getShortWeightUnitName()}`
  }

  getWeightUnitName(): string {
    return WeightUnitNameMap.get(this.weightUnit) ?? DEFAULT_WEIGHT_UNIT;
  }

  getShortWeightUnitName(): string {
    return ShortWeightUnitNameMap.get(this.weightUnit) ?? "";
  }
}