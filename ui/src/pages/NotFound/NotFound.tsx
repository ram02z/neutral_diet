import { Typography } from '@mui/material';
import useHeader from '@/store/header';

import { FullSizeCenteredFlexBox } from '@/components/styled';

function NotFound() {
  const [,headerActions] = useHeader();
  headerActions.change("Neutral Diet")

  return (
    <>
      <FullSizeCenteredFlexBox>
        <Typography variant="h3">NotFound</Typography>
      </FullSizeCenteredFlexBox>
    </>
  );
}

export default NotFound;
