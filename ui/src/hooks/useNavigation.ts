import { useRecoilState } from 'recoil';

import NavigationState from '@/store/navigation';
import { Actions } from '@/store/navigation/types';

export function useNavigation(): [number, Actions] {
  const [value, setValue] = useRecoilState(NavigationState);

  function change(_: React.SyntheticEvent<Element, Event>, value: number) {
    setValue(value);
  }

  function reset() {
    setValue(-1);
  }

  return [value, { change, reset }];
}
