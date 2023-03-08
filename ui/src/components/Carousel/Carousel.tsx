import { ReactNode, useState } from 'react';

import { KeyboardArrowLeft, KeyboardArrowRight } from '@mui/icons-material';
import { Box, Button, MobileStepper } from '@mui/material';
import { useTheme } from '@mui/system';

type CarouselProps = {
  children: ReactNode[];
};

function Carousel(props: CarouselProps) {
  const theme = useTheme();
  const [activeStep, setActiveStep] = useState(0);
  const maxSteps = props.children.length;

  const handleNext = () => {
    setActiveStep((prevActiveStep) => (prevActiveStep + 1) % maxSteps);
  };

  const handleBack = () => {
    setActiveStep((prevActiveStep) => (maxSteps + (prevActiveStep - 1)) % maxSteps);
  };

  return (
    <Box>
      <Box sx={{ width: '100%', p: 2 }}>{props.children[activeStep]}</Box>
      <MobileStepper
        steps={maxSteps}
        position="static"
        activeStep={activeStep}
        nextButton={
          <Button size="small" onClick={handleNext}>
            Next
            {theme.direction === 'rtl' ? <KeyboardArrowLeft /> : <KeyboardArrowRight />}
          </Button>
        }
        backButton={
          <Button size="small" onClick={handleBack}>
            {theme.direction === 'rtl' ? <KeyboardArrowRight /> : <KeyboardArrowLeft />}
            Back
          </Button>
        }
      />
    </Box>
  );
}

export default Carousel;
