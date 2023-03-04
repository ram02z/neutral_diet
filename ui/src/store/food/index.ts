import { atom, selector, selectorFamily } from 'recoil';

import client from '@/api/food_service';
import { AggregateFoodItem, FoodItemInfo } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import { Region } from '@/api/gen/neutral_diet/food/v1/region_pb';
import UserRegion from '@/core/regions';
import { LocalUserSettingsState } from '@/store/user';

import { FoodItemInfoQueryParams } from './types';

export const RegionsState = selector({
  key: 'RegionsState',
  get: () => {
    return Object.values(Region)
      .filter((x) => typeof x === 'number')
      .map((r) => new UserRegion(r as Region));
  },
});

export const FoodItemsState = selector({
  key: 'FoodItems',
  get: async ({ get }) => {
    const userSettings = get(LocalUserSettingsState);
    const response = await client.listAggregateFoodItems({ region: userSettings.region });
    return response.foodItems;
  },
});

export const FoodItemInfoQuery = selectorFamily<FoodItemInfo | undefined, FoodItemInfoQueryParams>({
  key: 'FoodItemInfoQuery',
  get:
    ({ foodItemId, region }) =>
    async () => {
      const response = await client.getFoodItemInfo({ id: foodItemId, region: region });
      return response.foodItemInfo;
    },
  cachePolicy_UNSTABLE: {
    eviction: 'keep-all',
  },
});

export const TypologiesState = selector({
  key: 'TypologiesState',
  get: async () => {
    const response = await client.listTypologyNames({});
    return response.names;
  },
});

export const SubTypologiesState = selector({
  key: 'SubTypologiesState',
  get: async () => {
    const response = await client.listSubTypologyNames({});
    return response.names;
  },
});

// TODO: should food history be managed by user?
export const FoodHistoryState = atom<AggregateFoodItem[]>({
  key: 'FoodHistoryState',
  default: [],
  // FIXME: disabled persist since enum gets stored as string
  // effects: [persistAtom],
});
