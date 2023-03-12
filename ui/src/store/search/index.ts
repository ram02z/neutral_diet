import { atom, selector } from 'recoil';

import { FoodHistoryState, FoodItemsState } from '@/store/food';

import { SearchFilters, SearchType } from './types';

export const SearchTypeState = atom({
  key: 'SearchTypeState',
  default: SearchType.History,
});

export const SearchTextState = atom({
  key: 'SearchTextState',
  default: '',
});

export const SelectedSearchFiltersState = atom<SearchFilters>({
  key: 'SelectedSearchFiltersState',
  default: { typologies: [], subTypologies: [] },
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
