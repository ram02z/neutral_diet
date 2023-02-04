import { useState } from 'react';
import { useRecoilValue } from 'recoil';

import { Close, SearchRounded } from '@mui/icons-material';
import { FormControl, IconButton, InputAdornment, OutlinedInput } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { Box } from '@mui/system';

import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import FoodItemCard from '@/components/FoodItemCard';
import { FoodItemsState } from '@/store/food';

function Search() {
  const foodItems = useRecoilValue(FoodItemsState);
  const [searchFoodItems, setSearchFoodItems] = useState<AggregateFoodItem[]>([]);
  const [searchText, setSearchText] = useState('');

  const handleSubmit = (e) => {
    if (e?.key === 'Enter') {
      const newFoodItems = foodItems.filter((foodItem) => {
        return foodItem.foodName.toLowerCase().includes(searchText.toLowerCase());
      });
      setSearchFoodItems(newFoodItems);
    }
  };

  const clearSearch = () => {
    setSearchText('');
    setSearchFoodItems([]);
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
        <Grid xs={8} sm={7} md={6} lg={5} xl={4}>
          <FormControl fullWidth variant="outlined">
            <OutlinedInput
              placeholder="Search for food"
              value={searchText}
              onKeyDown={handleSubmit}
              onChange={(e) => setSearchText(e.target.value as string)}
              startAdornment={
                <InputAdornment position="start">
                  <SearchRounded />
                </InputAdornment>
              }
              endAdornment={
                searchText.length ? (
                  <InputAdornment position="end">
                    <IconButton onClick={clearSearch}>
                      <Close />
                    </IconButton>
                  </InputAdornment>
                ) : undefined
              }
            />
          </FormControl>
        </Grid>
        {searchFoodItems.map((foodItem, idx) => (
          <Grid key={idx} xs={8} sm={7} md={6} lg={5} xl={4}>
            <FoodItemCard foodItem={foodItem} />
          </Grid>
        ))}
      </Grid>
    </Box>
  );
}

export default Search;
