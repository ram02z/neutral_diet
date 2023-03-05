import { Card, CardContent, Grid, Stack, Typography } from '@mui/material';

import { MIN_CARD_WIDTH } from '@/config';

type UserStatCardProps = {
  title: string;
  carbonFootprint: number;
};

function UserStatCard({ title, carbonFootprint }: UserStatCardProps) {
  return (
    <Card sx={{ minWidth: MIN_CARD_WIDTH }}>
      <CardContent>
        <Grid container columns={5}>
          <Stack spacing={1} sx={{ pt: 1, pl: 1 }}>
            <Typography sx={{ textTransform: 'capitalize' }} variant="h5" component="div">
              {title}
            </Typography>
            <Typography variant="subtitle1" color="text.secondary">
              <b>{parseFloat(carbonFootprint.toFixed(3))}</b>
              <Typography variant="caption">
                CO<sub>2</sub>/kg
              </Typography>
            </Typography>
          </Stack>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default UserStatCard;
