import { atom, selector, selectorFamily } from 'recoil';

import client from '@/api/food_service';
import { AggregateFoodItem, FoodItemInfo } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import { persistAtom } from '@/store';

export const RegionsState = atom({
  key: 'RegionsState',
  default: selector({
    key: 'Regions',
    get: async () => {
      const response = await client.listRegions({});
      return response.regions;
    },
  }),
});

export const FoodItemsState = atom({
  key: 'FoodItemsState',
  default: selector({
    key: 'FoodItems',
    get: async () => {
      const response = await client.listAggregateFoodItems({});
      return response.foodItems;
    },
  }),
});

export const FoodItemInfoQuery = selectorFamily<FoodItemInfo | undefined, number>({
  key: 'FoodItemInfoQuery',
  get: (foodItemId) => async () => {
    const response = await client.getFoodItemInfo({ id: foodItemId });
    return response.foodItemInfo;
  },
  cachePolicy_UNSTABLE: {
    eviction: 'keep-all',
  },
});

// TODO: should food history be managed by user?
export const FoodHistoryState = atom<AggregateFoodItem[]>({
  key: 'FoodHistoryState',
  default: [],
  effects: [persistAtom],
});
