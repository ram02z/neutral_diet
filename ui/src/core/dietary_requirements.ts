import { UserSettings_DietaryRequirement } from '@/api/gen/neutral_diet/user/v1/user_pb';

class DietaryRequirement {
  value: UserSettings_DietaryRequirement;

  constructor(dr: UserSettings_DietaryRequirement) {
    this.value = dr;
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
