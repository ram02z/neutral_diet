import { useSetRecoilState } from 'recoil';

import { Delete } from '@mui/icons-material';
import { Alert, IconButton } from '@mui/material';

import { useConfirm } from 'material-ui-confirm';

import { FoodHistoryState } from '@/store/food';

/**
 * Button to clear search history with confirmation.
 */
function ClearHistoryButton() {
  const confirm = useConfirm();
  const setFoodHistory = useSetRecoilState(FoodHistoryState);

  const clearHistory = () => {
    confirm({
      title: 'Clear history',
      content: (
        <div>
          <Alert variant="filled" severity="error">
            Are you sure you want to clear your history?
          </Alert>
        </div>
      ),
      cancellationButtonProps: { color: 'secondary' },
      confirmationText: 'Clear',
      confirmationButtonProps: { color: 'error' },
    }).then(() => {
      setFoodHistory([]);
    });
  };

  return (
    <IconButton onClick={clearHistory}>
      <Delete></Delete>
    </IconButton>
  );
}

export default ClearHistoryButton;
