import { UserSettings_DietaryRequirement } from '@/api/gen/neutral_diet/user/v1/user_pb';

export type Insight = {
  title: string;
  emissions: number;
  source: string;
};

class Insights {
  dietaryRequirement: UserSettings_DietaryRequirement;

  constructor(dr: UserSettings_DietaryRequirement) {
    this.dietaryRequirement = dr;
  }

  #getEmissionsByDietaryRequiremnent(): number {
    switch (this.dietaryRequirement) {
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

  get insights(): Insight[] {
    const dre = this.#getEmissionsByDietaryRequiremnent();
    return [
      {
        title: 'Diet average (UK)',
        emissions: dre,
        source: 'https://doi.org/10.1007/s10584-014-1169-1',
      },
      {
        title: 'LiveWell 2020 plate',
        emissions: 4.25,
        source: 'https://www.wwf.org.uk/what-we-do/livewell',
      },
      {
        title: 'LiveWell 2030 plate',
        emissions: 4.09,
        source: 'https://www.wwf.org.uk/what-we-do/livewell',
      },
    ];
  }
}

export default Insights;
