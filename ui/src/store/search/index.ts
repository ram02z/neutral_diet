import { atom, selector } from 'recoil';

import { SortMethod } from '@/core/sort';
import { FoodHistoryState, FoodItemsState } from '@/store/food';

import { SearchFilters, SearchSortMethod, SearchType } from './types';

export const SearchTypeState = atom({
  key: 'SearchTypeState',
  default: SearchType.History,
});

export const SearchSortState = atom({
  key: 'SearchSortState',
  default: SearchSortMethod.NameDescending,
});

export const SearchTextState = atom({
  key: 'SearchTextState',
  default: '',
});

export const SelectedSearchFiltersState = atom<SearchFilters>({
  key: 'SelectedSearchFiltersState',
  default: { typologies: [], subTypologies: [] },
});

export const SearchSortMethodsState = selector({
  key: 'SearchSortMethodsState',
  get: () => {
    return Object.values(SearchSortMethod)
      .filter((x) => typeof x === 'number')
      .map((m) => new SortMethod(m as SearchSortMethod));
  },
});

export const FilteredSearchFoodItemsState = selector({
  key: 'FilteredSearchFoodItemsState',
  get: ({ get }) => {
    const searchType = get(SearchTypeState);
    let foodItems;
    switch (searchType) {
      case SearchType.History: {
        foodItems = get(FoodHistoryState);
        break;
      }
      case SearchType.Global: {
        foodItems = get(FoodItemsState);
        break;
      }
    }
    const searchText = get(SearchTextState);
    const searchFilters = get(SelectedSearchFiltersState);

    return foodItems.filter((foodItem) => {
      if (
        (searchFilters.typologies.length > 0 &&
          !searchFilters.typologies.includes(foodItem.typology)) ||
        (searchFilters.subTypologies.length > 0 &&
          !searchFilters.subTypologies.includes(foodItem.subTypology))
      )
        return;
      return foodItem.name.toLowerCase().includes(searchText.toLowerCase());
    });
  },
});

export const SortedSearchFoodItemsState = selector({
  key: 'SortedSearchFoodItemsState',
  get: ({ get }) => {
    const sortMethod = get(SearchSortState);
    const filteredSearch = get(FilteredSearchFoodItemsState);

    switch (sortMethod) {
      case SearchSortMethod.NameDescending:
        return [...filteredSearch].sort((itemA, itemB) =>
          new Intl.Collator('en').compare(itemA.name, itemB.name),
        );
      case SearchSortMethod.NameAscending:
        return [...filteredSearch].sort(
          (itemA, itemB) => new Intl.Collator('en').compare(itemA.name, itemB.name) * -1,
        );
      case SearchSortMethod.EmissionsDescending:
        return [...filteredSearch].sort((itemA, itemB) =>
          itemA.carbonFootprint < itemB.carbonFootprint ? 1 : -1,
        );
      case SearchSortMethod.EmissionsAscending:
        return [...filteredSearch].sort((itemA, itemB) =>
          itemA.carbonFootprint > itemB.carbonFootprint ? 1 : -1,
        );
    }
  },
});
