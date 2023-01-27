import StyledSlider from './StyledSlider';

function CarbonFootprintSlider() {
  const marks = [
    { value: 0.0, label: '0.0kg' },
    { value: 10.0, label: '10.0kg' },
  ];
  const valueText = (value: number) => `${value}kg`;

  return (
    <StyledSlider
      label="Carbon footprint limit (CO2e/day)"
      marks={marks}
      defaultValue={0.0}
      step={0.5}
      valueText={valueText}
    />
  );
}

export default CarbonFootprintSlider;
