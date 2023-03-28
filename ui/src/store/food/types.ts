export type FoodItemInfoQueryParams = {
  foodItemId: number;
  region: number;
};

export type FoodItem = {
  id: number;
  name: string;
  carbonFootprint: number;
  region: number;
  typology: string;
  subTypology: string;
};
