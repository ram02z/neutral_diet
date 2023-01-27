import { useEffect, useState } from 'react';

import Loading from '@/components/Loading';
import { useCurrentUserSettings } from '@/hooks/useCurrentUserSettings';

import StyledSlider from './StyledSlider';

function CarbonFootprintSlider() {
  const userSettings = useCurrentUserSettings();
  const [cfLimit, setCfLimit] = useState(0.0);

  const marks = [
    { value: 0.0, label: '0.0kg' },
    { value: 10.0, label: '10.0kg' },
  ];

  const valueText = (value: number) => `${value}kg`;

  const handleChange = (event: Event, newValue: number | number[]) => {
    setCfLimit(newValue as number);
  };

  useEffect(() => {
    if (userSettings) setCfLimit(userSettings.cfLimit);
  }, [userSettings]);

  if (userSettings === undefined || userSettings === null) {
    return (
      <>
        <Loading />
      </>
    );
  } else {
    return (
      <StyledSlider
        label="Carbon footprint limit (CO2e/day)"
        marks={marks}
        value={cfLimit}
        step={0.5}
        handleChange={handleChange}
        valueText={valueText}
      />
    );
  }
}

export default CarbonFootprintSlider;
