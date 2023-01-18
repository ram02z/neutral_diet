import { useEffect, useState } from 'react';

import { Box } from '@mui/system';
import { DataGrid, GridColDef } from '@mui/x-data-grid';

import client from '@/api/food_service';
import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import { FullSizeCenteredFlexBox } from '@/components/styled';

const columns: GridColDef[] = [
  { field: 'typologyName', headerName: 'Category', flex: 1 },
  { field: 'subTypologyName', headerName: 'Sub Category', flex: 1 },
  { field: 'foodName', headerName: 'Name', flex: 1 },
  { field: 'n', headerName: 'N', flex: 1 },
  { field: 'medianCarbonFootprint', headerName: 'CO2e/kg', flex: 1 },
];

function Home() {
  const [foodItems, setFoodItems] = useState<AggregateFoodItem[]>([]);
  useEffect(() => {
    client
      .listAggregateFoodItems({})
      .then((response) => setFoodItems(response.foodItems))
      .catch((e) => console.error(e.message));
  }, []);

  return (
    <>
      <FullSizeCenteredFlexBox>
        <Box sx={{ height: 1000, width: '100%' }}>
          <DataGrid
            rows={foodItems}
            columns={columns}
            getRowId={(row: AggregateFoodItem) => row.id}
          />
        </Box>
      </FullSizeCenteredFlexBox>
    </>
  );
}

export default Home;
