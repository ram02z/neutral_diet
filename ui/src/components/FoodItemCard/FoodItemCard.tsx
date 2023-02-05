import { useMemo, useState } from 'react';
import { useRecoilState } from 'recoil';

import { Add } from '@mui/icons-material';
import { Card, CardContent, IconButton, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

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
        <Grid container columns={3}>
          <Grid xs={2} sx={{ pt: 1, pl: 1 }}>
            <Typography sx={{ textTransform: 'capitalize' }} variant="h5" component="div">
              {foodItem.foodName.toLowerCase()}
            </Typography>
            <Typography variant="subtitle1" color="text.secondary" component="div">
              {foodItem.medianCarbonFootprint}
            </Typography>
          </Grid>
          <Grid
            xs
            display="flex"
            justifyContent="flex-end"
            alignItems="center"
            sx={{ pt: 1, pl: 1 }}
          >
            <IconButton onClick={addToHistory} aria-label="add" sx={{ float: 'right' }}>
              <Add />
            </IconButton>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default FoodItemCard;
