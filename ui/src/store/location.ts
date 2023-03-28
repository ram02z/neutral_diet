import { atom } from 'recoil';

export const PrivateRoutePathState = atom<string | null>({
  key: 'PrivateRoutePathState',
  default: null,
});
