export type Source = {
  reference: string;
  year: number;
  regionName: string;
};

export type FoodItemInfo = {
  typologyName: string;
  subTypologyName: string;
  nonUniqueSources: bigint;
  sources: Source[];
};
