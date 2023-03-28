import { atom, selector, selectorFamily } from 'recoil';

import client from '@/api/food_service';
import { FoodItemInfo } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import { Region } from '@/api/gen/neutral_diet/food/v1/region_pb';
import UserRegion from '@/core/regions';
import { persistAtom } from '@/store';
import { LocalUserSettingsState } from '@/store/user';

import { FoodItem, FoodItemInfoQueryParams } from './types';

export const RegionsState = selector({
  key: 'RegionsState',
  get: () => {
    return Object.values(Region)
      .filter((x) => typeof x === 'number')
      .map((r) => new UserRegion(r as Region));
  },
});

export const FoodItemsState = selector<FoodItem[]>({
  key: 'FoodItems',
  get: async ({ get }) => {
    const userSettings = get(LocalUserSettingsState);
    const response = await client.listAggregateFoodItems({ region: userSettings.region });
    return response.foodItems.map((a) => {
      return {
        id: a.id,
        name: a.foodName,
        carbonFootprint: a.medianCarbonFootprint,
        region: a.region,
        typology: a.typologyName,
        subTypology: a.subTypologyName,
      };
    });
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

export const FoodHistoryState = atom<FoodItem[]>({
  key: 'FoodHistoryState',
  default: [],
  effects: [persistAtom],
});
