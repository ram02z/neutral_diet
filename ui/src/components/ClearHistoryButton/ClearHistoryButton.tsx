import { FoodHistoryState } from '@/store/food';
import { Delete } from '@mui/icons-material';
import { Alert, IconButton } from '@mui/material';
import { useConfirm } from 'material-ui-confirm';
import { useSetRecoilState } from 'recoil';

function ClearHistoryButton() {
  const confirm = useConfirm();
  const setFoodHistory = useSetRecoilState(FoodHistoryState);

  const clearHistory = () => {
    confirm({
      title: 'Clear history',
      content: (
        <div>
          <Alert severity="error">
          Are you sure you want to clear your history?
          </Alert>
        </div>
      ),
      cancellationButtonProps: { color: 'info' },
      confirmationText: 'Clear',
      confirmationButtonProps: { color: 'error', variant: 'contained' },
    }).then(() => {
      setFoodHistory([])
    });
  }

  return (
    <IconButton onClick={clearHistory}>
      <Delete></Delete>
    </IconButton>
  );
}

export default ClearHistoryButton;
