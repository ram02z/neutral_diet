import { atom } from 'recoil';

const NavigationState = atom<number>({
  key: 'NavigationState',
  default: 0,
});

export default NavigationState;
