import { atom, selector } from 'recoil';

import client from '@/api/food_service';

export const RegionsState = atom({
  key: 'RegionsState',
  default: selector({
    key: 'Regions',
    get: async () => {
      const response = await client.listRegions({});
      await new Promise(r => setTimeout(r, 2000));
      return response.regions;
    },
  }),
});
