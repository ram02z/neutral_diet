import { useRecoilValue, useSetRecoilState } from 'recoil';

import { LocalUserSettingsState } from '@/store/user';

import StyledSlider from './StyledSlider';

function CarbonFootprintSlider() {
  const localUserSettings = useRecoilValue(LocalUserSettingsState);
  const setLocalUserSettings = useSetRecoilState(LocalUserSettingsState);

  const marks = [
    { value: 0.0, label: '0.0kg' },
    { value: 10.0, label: '10.0kg' },
  ];

  const valueText = (value: number) => `${value}kg`;

  const handleChange = (_: Event, newValue: number | number[]) => {
    setLocalUserSettings((old) => {
      return {
        ...old,
        cfLimit: newValue as number,
        dirty: true,
      };
    });
  };

  return (
    <StyledSlider
      label="Carbon footprint limit (CO2e/day)"
      marks={marks}
      value={localUserSettings.cfLimit}
      step={0.5}
      handleChange={handleChange}
      valueText={valueText}
    />
  );
}

export default CarbonFootprintSlider;
