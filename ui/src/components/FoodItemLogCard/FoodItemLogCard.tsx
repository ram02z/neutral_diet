import { Card, CardContent, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import { FoodLogItemResponse } from '@/api/gen/neutral_diet/user/v1/food_item_log_pb';
import { MIN_WIDTH } from '@/config';

type FoodItemCardProps = {
  foodLogItem: FoodLogItemResponse;
};

function FoodItemLogCard({ foodLogItem }: FoodItemCardProps) {
  return (
    <Card sx={{ minWidth: MIN_WIDTH }}>
      <CardContent>
        <Grid container columns={5}>
          <Grid xs={4} sx={{ pt: 1, pl: 1 }}>
            <Typography sx={{ textTransform: 'capitalize' }} variant="h5" component="div">
              {foodLogItem.name.toLowerCase()}
            </Typography>
            <Typography variant="subtitle1" color="text.secondary" component="div">
              {foodLogItem.carbonFootprint}
            </Typography>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default FoodItemLogCard;
