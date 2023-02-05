import { useMemo, useState } from 'react';
import { useRecoilState } from 'recoil';

import { Add } from '@mui/icons-material';
import { Card, CardContent, IconButton, Typography } from '@mui/material';

import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import { FoodHistoryState } from '@/store/food';

type FoodItemCardProps = {
  foodItem: AggregateFoodItem;
};

function FoodItemCard({ foodItem }: FoodItemCardProps) {
  const [foodHistory, setFoodHistory] = useRecoilState(FoodHistoryState);
  const [inHistory, setInHistory] = useState(false);

  useMemo(() => {
    setInHistory(foodHistory.some((value) => value.id === foodItem.id));
  }, [foodHistory, foodItem]);

  const addToHistory = () => {
    if (inHistory) return;
    setFoodHistory((oldFoodHistory) => [...oldFoodHistory, foodItem]);
  };

  return (
    <Card sx={{ minWidth: 250 }}>
        <CardContent>
          <IconButton
            onClick={addToHistory}
            sx={{ float: 'right' }}
            aria-label="add"
          >
            <Add />
          </IconButton>
          <Typography sx={{ textTransform: 'capitalize' }} variant="h5" component="div">
            {foodItem.foodName.toLowerCase()}
          </Typography>
          <Typography variant="subtitle1" color="text.secondary" component="div">
            {foodItem.medianCarbonFootprint}
          </Typography>
        </CardContent>
    </Card>
  );
}

export default FoodItemCard;
