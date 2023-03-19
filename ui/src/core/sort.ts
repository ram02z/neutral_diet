import { SearchSortMethod } from '@/store/search/types';

export class SortMethod {
  value: SearchSortMethod;

  constructor(method: SearchSortMethod) {
    this.value = method;
  }

  getName(): string {
    switch (this.value) {
      case SearchSortMethod.NameDescending:
        return 'A to Z';
      case SearchSortMethod.NameAscending:
        return 'Z to A';
      case SearchSortMethod.EmissionsAscending:
        return 'Emissions, Low to High';
      case SearchSortMethod.EmissionsDescending:
        return 'Emissions, High to Low';
    }
  }
}
