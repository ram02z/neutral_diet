import { useRecoilValue } from 'recoil';

import { SearchRounded } from '@mui/icons-material';
import { InputAdornment, TextField } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { Box } from '@mui/system';

import FoodItemCard from '@/components/FoodItemCard';
import { FoodItemsState } from '@/store/food';
import { useState } from 'react';
import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';

function Search() {
  const foodItems = useRecoilValue(FoodItemsState)
  const [searchFoodItems, setSearchFoodItems] = useState<AggregateFoodItem[]>([]);
  const [search, setSearch] = useState('')

  const handleSubmit = (e) => {
    if (e?.key === 'Enter') {
      const newFoodItems = foodItems.filter((foodItem) => {
        return foodItem.foodName.toLowerCase().includes(search.toLowerCase())
      })
      setSearchFoodItems(newFoodItems)
    }
  };

  // TODO: use react-window
  return (
    <Box>
      <Grid
        container
        direction="column"
        columns={10}
        spacing={4}
        alignItems="center"
        pt="4vh"
        disableEqualOverflow
      >
        <Grid>
          <TextField
            variant="outlined"
            placeholder="Search for food"
            value={search}
            onKeyDown={handleSubmit}
            onChange={(e) => setSearch(e.target.value as string)}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <SearchRounded />
                </InputAdornment>
              ),
            }}
          />
        </Grid>
        {searchFoodItems.map((foodItem, idx) => (
          <Grid key={idx}>
            <FoodItemCard foodItem={foodItem} />
          </Grid>
        ))}
      </Grid>
    </Box>
  );
}

export default Search;
