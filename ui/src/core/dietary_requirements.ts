import { UserSettings_DietaryRequirement } from '@/api/gen/neutral_diet/user/v1/user_pb';

class DietaryRequirement {
  value: UserSettings_DietaryRequirement;

  constructor(dr: UserSettings_DietaryRequirement) {
    this.value = dr;
  }

  // Uses data adjusted for age and sex per 2000kcal
  // From doi.org/10.1007/s10584-014-1169-1
  getMeanEmissionsPerDay(): number {
    switch (this.value) {
      case UserSettings_DietaryRequirement.PESCATARIAN:
        return 3.91;
      case UserSettings_DietaryRequirement.VEGETARIAN:
        return 3.81;
      case UserSettings_DietaryRequirement.VEGAN:
        return 2.89;
      default:
        return 5.76;
    }
  }

  getSettingName(): string {
    switch (this.value) {
      case UserSettings_DietaryRequirement.PESCATARIAN:
        return 'Pescatarian';
      case UserSettings_DietaryRequirement.VEGETARIAN:
        return 'Vegetarian';
      case UserSettings_DietaryRequirement.VEGAN:
        return 'Vegan';
      default:
        return 'N/A';
    }
  }
}

export default DietaryRequirement;
