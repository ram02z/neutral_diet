import { useEffect, useState } from 'react';

import client from '@/api/food_service';
import { FoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import { FullSizeCenteredFlexBox } from '@/components/styled';

function Home() {
  const [foodItems, setFoodItems] = useState<FoodItem[]>([]);
  useEffect(() => {
    client
      .listFoodItems({})
      .then((response) => setFoodItems(response.foodItems))
      .catch((e) => console.error(e.message));
  }, []);
  return (
    <>
      <FullSizeCenteredFlexBox>
        {foodItems.map((foodItem, idx) => (
          <p key={idx}>
            {foodItem.name}-{foodItem.emissions}
          </p>
        ))}
      </FullSizeCenteredFlexBox>
    </>
  );
}

export default Home;
