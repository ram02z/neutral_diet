import { useRecoilValue, useSetRecoilState } from 'recoil';

import { Delete } from '@mui/icons-material';
import { Card, CardActions, CardContent, IconButton, Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import { ID_TOKEN_HEADER } from '@/api/transport';
import client from '@/api/user_service';
import { MIN_WIDTH } from '@/config';
import { CurrentUserTokenIDState, LocalFoodItemLogState } from '@/store/user';
import { LocalFoodLogItem } from '@/store/user/types';

type FoodItemCardProps = {
  foodLogItem: LocalFoodLogItem;
};

function FoodItemLogCard({ foodLogItem }: FoodItemCardProps) {
  const idToken = useRecoilValue(CurrentUserTokenIDState);
  const setFoodItemLog = useSetRecoilState(LocalFoodItemLogState);

  const handleDelete = () => {
    if (idToken) {
      const headers = new Headers();
      headers.set(ID_TOKEN_HEADER, idToken);
      client.deleteFoodItem({ id: foodLogItem.remoteId }, { headers: headers }).then(() => {
        setFoodItemLog((old) => old.filter((item) => item.remoteId !== foodLogItem.remoteId));
      });
    }
  };

  return (
    <Card sx={{ minWidth: MIN_WIDTH }}>
      <CardContent>
        <Grid container columns={5}>
          <Grid xs={4} sx={{ pt: 1, pl: 1 }}>
            <Typography sx={{ textTransform: 'capitalize' }} variant="h5" component="div">
              {foodLogItem.name.toLowerCase()}
            </Typography>
            <Typography variant="subtitle1" color="text.secondary" component="div">
              {foodLogItem.carbonFootprint}
            </Typography>
          </Grid>
        </Grid>
      </CardContent>
      <CardActions disableSpacing>
        <IconButton onClick={handleDelete}>
          <Delete />
        </IconButton>
      </CardActions>
    </Card>
  );
}

export default FoodItemLogCard;
