import DietaryRequirement from "@/core/dietary_requirements";

export type LocalUserSettings = {
  region: string;
  cfLimit: number;
  dietaryRequirement: DietaryRequirement;
  dirty: boolean;
};
