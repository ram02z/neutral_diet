import { atom, selector } from 'recoil';

import client from '@/api/food_service';

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
