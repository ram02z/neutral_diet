import { useState } from 'react';
import { useRecoilValue, useSetRecoilState } from 'recoil';

import { MAX_CF_LIMIT } from '@/config';
import { LocalUserSettingsState } from '@/store/user';

import StyledSlider from './StyledSlider';

function CarbonFootprintSlider() {
  const localUserSettings = useRecoilValue(LocalUserSettingsState);
  const setLocalUserSettings = useSetRecoilState(LocalUserSettingsState);
  const [value, setValue] = useState(localUserSettings.cfLimit);

  const marks = [
    { value: 0.0, label: '0.0kg' },
    { value: MAX_CF_LIMIT, label: `${MAX_CF_LIMIT.toFixed(1)}kg` },
  ];

  const valueText = (value: number) => `${value}kg`;

  const handleChange = (_: Event, newValue: number | number[]) => {
    setValue(newValue as number);
  };

  const handleChangeCommitted = () => {
    if (value === localUserSettings.cfLimit) return;

    setLocalUserSettings((old) => {
      return {
        ...old,
        cfLimit: value as number,
        dirty: true,
      };
    });
  };

  return (
    <StyledSlider
      label="Carbon footprint limit (CO2e/day)"
      marks={marks}
      value={value}
      step={0.5}
      handleChange={handleChange}
      handleChangeCommitted={handleChangeCommitted}
      valueText={valueText}
    />
  );
}

export default CarbonFootprintSlider;
