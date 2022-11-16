import { atom, useRecoilState } from 'recoil';

import type { Actions } from './types';

const navigationState = atom<string>({
  key: 'sidebar-openness-state',
  default: "",
});

function useHeader(): [string, Actions] {
  const [value, setValue] = useRecoilState(navigationState);

  function change(value: string) {
    setValue(value)
  }

  return [value, { change }];
}

export default useHeader;
