export type LocalUserSettings = {
  region: string;
  cfLimit: number;
  dietaryRequirement: number;
  dirty: boolean;
};

export type LocalFoodLogItem = {
  remoteId: number;
  name: string;
  weight: number;
  carbonFootprint: number;
};
