import { Box, Slider, SliderValueLabelProps, Tooltip, Typography } from '@mui/material';

type SliderMark = {
  value: number;
  label: string;
}

type StyledSliderProps = {
  label: string;
  defaultValue: number;
  step: number;
  marks: Array<SliderMark>;
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

function StyledSlider({ label, defaultValue, step, marks, valueText }: StyledSliderProps) {
  const max = Math.max(...marks.map(m => m.value));
  return (
    <Box>
      <Typography gutterBottom>{label}</Typography>
      <Slider
        valueLabelDisplay="auto"
        slots={{ valueLabel: ValueLabelComponent }}
        aria-label={label}
        getAriaValueText={valueText}
        defaultValue={defaultValue}
        step={step}
        marks={marks}
        max={max}
      />
    </Box>
  );
}

export default StyledSlider;
