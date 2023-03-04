import { SubmitHandler } from 'react-hook-form';
import RenderIfVisible from 'react-render-if-visible';
import { useRecoilState, useRecoilValue } from 'recoil';

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

import ClearHistoryButton from '@/components/ClearHistoryButton';
import FoodItemCard from '@/components/FoodItemCard';
import { ESTIMATED_CARD_HEIGHT } from '@/components/FoodItemCard/FoodItemCard';
import SortFilterMenu from '@/components/SortFilterMenu';
import { FormValues } from '@/components/SortFilterMenu/types';
import {
  FilteredSearchFoodItemsState,
  SearchTextState,
  SearchTypeState,
  SelectedSearchFiltersState,
} from '@/store/search';
import { SearchType } from '@/store/search/types';

function Search() {
  const searchFoodItems = useRecoilValue(FilteredSearchFoodItemsState);
  const [searchText, setSearchText] = useRecoilState(SearchTextState);
  const [searchType, setSearchType] = useRecoilState(SearchTypeState);
  const [searchFilters, setSearchFilters] = useRecoilState(SelectedSearchFiltersState);
  const handleSubmit = () => {
    setSearchType(SearchType.Global);
  };

  const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => {
    setSearchText(e.target.value as string);
    setSearchType(SearchType.History);
  };

  const clearSearch = () => {
    setSearchText('');
    setSearchType(SearchType.History);
  };

  const onFilterSubmit: SubmitHandler<FormValues> = (data) => {
    setSearchFilters({ typologies: data.typologyNames, subTypologies: data.subTypologyNames });
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
      {searchType === SearchType.History && (
        <>
          <Grid>
            <Typography variant="h4">History</Typography>
          </Grid>
          <Grid>
            <SortFilterMenu onSubmit={onFilterSubmit} currentSearchFilters={searchFilters} />
          </Grid>
          {searchFoodItems.map((foodItem, idx) => (
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
          {searchText.length == 0 && searchFoodItems.length > 0 && (
            <Grid>
              <ClearHistoryButton />
            </Grid>
          )}
        </>
      )}
      {searchType === SearchType.Global && (
        <>
          <Grid>
            <Typography variant="h4">Search Results</Typography>
          </Grid>
          <Grid>
            <SortFilterMenu onSubmit={onFilterSubmit} currentSearchFilters={searchFilters} />
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
