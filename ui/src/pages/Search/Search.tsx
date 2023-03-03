import { useEffect, useState } from 'react';
import RenderIfVisible from 'react-render-if-visible';
import { useRecoilValue } from 'recoil';

import { Close, SearchRounded, Tune } from '@mui/icons-material';
import {
  Button,
  FormControl,
  IconButton,
  InputAdornment,
  OutlinedInput,
  Typography,
} from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import ClearHistoryButton from '@/components/ClearHistoryButton';
import FoodItemCard from '@/components/FoodItemCard';
import { ESTIMATED_CARD_HEIGHT } from '@/components/FoodItemCard/FoodItemCard';
import { FoodHistoryState, FoodItemsState } from '@/store/food';

function Search() {
  const foodItems = useRecoilValue(FoodItemsState);
  const foodHistory = useRecoilValue(FoodHistoryState);
  const [searchFoodItems, setSearchFoodItems] = useState<AggregateFoodItem[]>([]);
  const [searchFoodHistory, setSearchFoodHistory] = useState<AggregateFoodItem[]>(foodHistory);
  const [searchText, setSearchText] = useState('');
  const [showHistory, setShowHistory] = useState(true);

  useEffect(() => {
    setSearchFoodHistory(foodHistory);
  }, [foodHistory]);

  const handleSearch = (foodItemArray: AggregateFoodItem[]) => {
    return foodItemArray.filter((foodItem) => {
      return foodItem.foodName.toLowerCase().includes(searchText.toLowerCase());
    });
  };

  const handleSubmit = () => {
    const newFoodItems = handleSearch(foodItems);
    setShowHistory(false);
    setSearchFoodItems(newFoodItems);
  };

  const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => {
    setSearchText(e.target.value as string);
    setSearchFoodItems([]);
    if (e.target.value.length > 0) {
      const newFoodItems = handleSearch(foodHistory);
      setSearchFoodHistory(newFoodItems);
    } else {
      setSearchFoodHistory(foodHistory);
    }
    setShowHistory(true);
  };

  const clearSearch = () => {
    setSearchText('');
    setSearchFoodItems([]);
    setSearchFoodHistory(foodHistory);
    setShowHistory(true);
  };

  return (
    <Grid
      container
      direction="column"
      columns={10}
      spacing={4}
      alignItems="center"
      pt="4vh"
      pb="8vh"
      disableEqualOverflow
    >
      <Grid xs={8} lg={7} xl={6}>
        <FormControl fullWidth variant="outlined">
          <OutlinedInput
            placeholder="Search for food"
            value={searchText}
            onKeyDown={(e) => {
              if (e.key == 'Enter') handleSubmit();
            }}
            onChange={handleChange}
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
      {showHistory ? (
        <>
          <Grid>
            <Typography variant="h4">History</Typography>
          </Grid>
          {searchFoodHistory.map((foodItem, idx) => (
            <Grid key={idx} xs={8} lg={7} xl={6}>
              <RenderIfVisible defaultHeight={ESTIMATED_CARD_HEIGHT}>
                <FoodItemCard foodItem={foodItem} />
              </RenderIfVisible>
            </Grid>
          ))}
          {searchText.length > 0 && (
            <Grid textAlign="center" xs={8} lg={7} xl={6}>
              <Button onClick={handleSubmit} variant="outlined" startIcon={<SearchRounded />}>
                {`Search all foods for "${searchText}"`}
              </Button>
            </Grid>
          )}
          {searchText.length == 0 && foodHistory.length > 0 && (
            <Grid>
              <ClearHistoryButton />
            </Grid>
          )}
        </>
      ) : (
        <>
          <Grid>
            <Typography variant="h4">Search Results</Typography>
          </Grid>
          <Grid>
            <Button variant="contained" startIcon={<Tune />}>
              Sort and filter
            </Button>
          </Grid>
          <Grid xs={8} lg={7} xl={6}>
            <Typography
              variant="subtitle1"
              color="text.secondary"
            >{`${searchFoodItems.length} results`}</Typography>
          </Grid>
          {searchFoodItems.map((foodItem, idx) => (
            <Grid key={idx} xs={8} lg={7} xl={6}>
              <RenderIfVisible defaultHeight={ESTIMATED_CARD_HEIGHT}>
                <FoodItemCard foodItem={foodItem} />
              </RenderIfVisible>
            </Grid>
          ))}
        </>
      )}
    </Grid>
  );
}

export default Search;
