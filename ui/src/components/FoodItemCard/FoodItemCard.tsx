import { Add } from '@mui/icons-material';
import { Card, CardActionArea, CardContent, IconButton, Typography } from '@mui/material';

import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';

type FoodItemCardProps = {
  foodItem: AggregateFoodItem;
};

function FoodItemCard({ foodItem }: FoodItemCardProps) {
  return (
    <Card sx={{ minWidth: 250 }}>
      <CardActionArea>
        <CardContent>
          <IconButton sx={{ float: 'right' }} aria-label="add">
            <Add />
          </IconButton>
          <Typography sx={{ textTransform: 'capitalize' }} variant="h5" component="div">
            {foodItem.foodName.toLowerCase()}
          </Typography>
          <Typography variant="subtitle1" color="text.secondary" component="div">
            {foodItem.medianCarbonFootprint}
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}

export default FoodItemCard;
