export type Source = {
  reference: string;
  year: number;
  region: number;
};

export type FoodItemInfo = {
  typologyName: string;
  subTypologyName: string;
  nonUniqueSources: bigint;
  sources: Source[];
};
