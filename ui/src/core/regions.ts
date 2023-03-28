import { Region } from '@/api/gen/neutral_diet/food/v1/region_pb';

class UserRegion {
  value: Region;

  constructor(r: Region) {
    this.value = r;
  }

  getSettingName(): string {
    switch (this.value) {
      case Region.ASIA:
        return 'Asia';
      case Region.AFRICA:
        return 'Africa';
      case Region.AMERICA:
        return 'America';
      case Region.EUROPE:
        return 'Europe';
      case Region.MEDITERRANEAN:
        return 'Mediterranean';
      case Region.OCEANIA:
        return 'Oceania';
      case Region.WORLD:
        return 'Rest of world';
      case Region.UNSPECIFIED:
        return 'All';
    }
  }
}

export default UserRegion;
