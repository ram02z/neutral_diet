import {
  Button,
  Chip,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Divider,
  Typography,
} from '@mui/material';
import { Stack } from '@mui/system';

import SourceTable from '@/components/SourceTable';
import { MIN_WIDTH } from '@/config';

import { FoodItemInfo } from './types';

type FoodItemInfoDialogProps = {
  openDialog: boolean;
  handleClose: () => void;
  foodItemInfo?: FoodItemInfo;
};

function FoodItemInfoDialog({ openDialog, handleClose, foodItemInfo }: FoodItemInfoDialogProps) {
  return (
    <Dialog fullWidth open={openDialog} onClose={handleClose}>
      <DialogTitle textAlign="center">Food information</DialogTitle>
      <DialogContent>
        <Stack minWidth={MIN_WIDTH} textAlign="center" spacing={3}>
          <Divider>
            <Chip label="Typology" />
          </Divider>
          <Typography variant="subtitle1" sx={{ textTransform: 'capitalize' }}>
            {foodItemInfo?.typologyName.toLowerCase()}
          </Typography>
          {foodItemInfo?.subTypologyName && (
            <>
              <Divider>
                <Chip label="Sub-Typology" />
              </Divider>
              <Typography variant="subtitle1" sx={{ textTransform: 'capitalize' }}>
                {foodItemInfo?.subTypologyName.toLowerCase()}
              </Typography>
            </>
          )}
          <Divider>
            <Chip label="Number of non-unique sources" />
          </Divider>
          <Typography variant="subtitle1">{foodItemInfo?.nonUniqueSources.toString()}</Typography>
          <Divider>
            <Chip label="Unique sources" />
          </Divider>
          <SourceTable sources={foodItemInfo?.sources ?? []} />
        </Stack>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose}>Close</Button>
      </DialogActions>
    </Dialog>
  );
}

export default FoodItemInfoDialog;
