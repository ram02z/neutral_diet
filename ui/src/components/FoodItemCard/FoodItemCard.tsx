import { Card, CardActionArea, CardContent } from '@mui/material';

import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';

type FoodItemCardProps = {
  foodItem: AggregateFoodItem;
};

function FoodItemCard({ foodItem }: FoodItemCardProps) {
  return (
    <Card sx={{ minWidth: 300 }}>
      <CardActionArea>
        <CardContent>{foodItem.foodName}</CardContent>
      </CardActionArea>
    </Card>
  );
}

export default FoodItemCard;
