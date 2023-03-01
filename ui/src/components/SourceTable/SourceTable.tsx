import { Paper, Table, TableBody, TableCell, TableHead, TableRow } from '@mui/material';
import TableContainer from '@mui/material/TableContainer';

import { Source } from '@/components/FoodItemInfoDialog/types';
import { MIN_CARD_WIDTH } from '@/config';

type SourceTableProps = {
  sources: Source[];
};

function SourceTable({ sources }: SourceTableProps) {
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: MIN_CARD_WIDTH }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Reference</TableCell>
            <TableCell align="right">Region</TableCell>
            <TableCell align="right">Year</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {sources.map((source) => (
            <TableRow
              key={source.reference}
              sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
            >
              <TableCell component="th" scope="row">
                {source.reference}
              </TableCell>
              <TableCell align="right">{source.regionName}</TableCell>
              <TableCell align="right">{source.year}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}

export default SourceTable;
