import { atom, useRecoilState } from 'recoil';

import type { Actions } from './types';

const navigationState = atom<number>({
  key: 'navigation-state',
  default: 0,
});

function useNavigation(): [number, Actions] {
  const [value, setValue] = useRecoilState(navigationState);

  function change(_: React.SyntheticEvent<Element, Event>, value: number) {
    setValue(value);
  }

  function reset() {
    setValue(-1);
  }

  return [value, { change, reset }];
}

export default useNavigation;
