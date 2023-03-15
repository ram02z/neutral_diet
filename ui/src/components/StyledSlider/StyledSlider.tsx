import { FormControl, Slider, SliderValueLabelProps, Tooltip, Typography } from '@mui/material';

type SliderMark = {
  value: number;
  label: string;
};

type StyledSliderProps = {
  label: string;
  value: number;
  step: number;
  marks: Array<SliderMark>;
  handleChange: (event: Event, value: number | number[]) => void;
  handleChangeCommitted: (
    event: Event | React.SyntheticEvent<Element, Event>,
    value: number | number[],
  ) => void;
  valueText?: (value: number) => string;
};

function ValueLabelComponent(props: SliderValueLabelProps) {
  const { children, value } = props;

  return (
    <Tooltip enterTouchDelay={0} placement="top" title={value}>
      {children}
    </Tooltip>
  );
}

function StyledSlider({
  label,
  value,
  step,
  marks,
  valueText,
  handleChange,
  handleChangeCommitted,
}: StyledSliderProps) {
  const min = Math.min(...marks.map((m) => m.value));
  const max = Math.max(...marks.map((m) => m.value));
  return (
    <FormControl fullWidth>
      <Typography gutterBottom>{label}</Typography>
      <Slider
        valueLabelDisplay="auto"
        slots={{ valueLabel: ValueLabelComponent }}
        aria-label={label}
        getAriaValueText={valueText}
        onChange={handleChange}
        onChangeCommitted={handleChangeCommitted}
        value={value}
        step={step}
        marks={marks}
        min={min}
        max={max}
      />
    </FormControl>
  );
}

export default StyledSlider;
