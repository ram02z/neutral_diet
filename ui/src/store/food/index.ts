import { atom, selector, selectorFamily } from 'recoil';

import client from '@/api/food_service';
import { AggregateFoodItem, FoodItemInfo } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import { Region } from '@/api/gen/neutral_diet/food/v1/region_pb';
import UserRegion from '@/core/regions';
import { persistAtom } from '@/store';
import { LocalUserSettingsState } from '@/store/user';

export const RegionsState = atom({
  key: 'RegionsState',
  default: selector({
    key: 'RegionsState/Default',
    get: () => {
      return Object.values(Region)
        .filter((x) => typeof x === 'number')
        .map((r) => new UserRegion(r as Region));
    },
  }),
});

export const FoodItemsState = atom({
  key: 'FoodItemsState',
  default: selector({
    key: 'FoodItems',
    get: async ({ get }) => {
      const userSettings = get(LocalUserSettingsState);
      const response = await client.listAggregateFoodItems({ region: userSettings.region });
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
