import { useEffect, useState } from 'react';

import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';

import client from '@/api/food_service';
import { AggregateFoodItem } from '@/api/gen/neutral_diet/food/v1/food_item_pb';
import { FullSizeCenteredFlexBox } from '@/components/styled';

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
        <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }}>
            <TableHead>
              <TableRow>
                <TableCell>Name</TableCell>
                <TableCell align="right">Carbon footprint (CO2e/kg)</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {foodItems.map((foodItem) => (
                <TableRow
                  key={foodItem.name}
                  sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                  <TableCell component="th" scope="row">
                    {foodItem.name}
                  </TableCell>
                  <TableCell align="right">{foodItem.medianEmissions}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </FullSizeCenteredFlexBox>
    </>
  );
}

export default Home;
