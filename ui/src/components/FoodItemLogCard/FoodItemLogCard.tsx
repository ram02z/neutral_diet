import { useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useRecoilValue, useSetRecoilState } from 'recoil';

import { Delete } from '@mui/icons-material';
import {
  Card,
  CardActionArea,
  CardActions,
  CardContent,
  IconButton,
  Typography,
} from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';

import { useSnackbar } from 'notistack';

import { ID_TOKEN_HEADER } from '@/api/transport';
import client from '@/api/user_service';
import EditFoodItemDialog from '@/components/EditFoodItemDialog';
import { FormValues } from '@/components/FoodItemCard/types';
import { MIN_WIDTH } from '@/config';
import { CurrentUserTokenIDState, FoodItemLogDateState, LocalFoodItemLogState } from '@/store/user';
import { LocalFoodLogItem } from '@/store/user/types';

type FoodItemCardProps = {
  foodLogItem: LocalFoodLogItem;
};

function FoodItemLogCard({ foodLogItem }: FoodItemCardProps) {
  const [openDialog, setOpenDialog] = useState(false);
  const idToken = useRecoilValue(CurrentUserTokenIDState);
  const date = useRecoilValue(FoodItemLogDateState);
  const { enqueueSnackbar } = useSnackbar();
  const setFoodItemLog = useSetRecoilState(LocalFoodItemLogState(date));

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    const weight = parseFloat(data.weight);
    if (idToken) {
      const headers = new Headers();
      headers.set(ID_TOKEN_HEADER, idToken);
      client
        .updateFoodItem(
          {
            id: foodLogItem.dbId,
            weight: weight,
          },
          { headers: headers },
        )
        .then((res) => {
          setFoodItemLog((old) => {
            return old.map((item) => {
              if (item.dbId == foodLogItem.dbId) {
                return {
                  dbId: foodLogItem.dbId,
                  name: foodLogItem.name,
                  weight: weight,
                  carbonFootprint: res.carbonFootprint,
                };
              }
              return { ...item };
            });
          });
          enqueueSnackbar('Updated food entry', { variant: 'success' });
        })
        .catch((err) => {
          enqueueSnackbar("Couldn't update food entry", { variant: 'error' });
          console.error(err);
        });
    }
    handleClose();
  };

  const handleDelete = () => {
    if (idToken) {
      const headers = new Headers();
      headers.set(ID_TOKEN_HEADER, idToken);
      client.deleteFoodItem({ id: foodLogItem.dbId }, { headers: headers }).then(() => {
        setFoodItemLog((old) => old.filter((item) => item.dbId !== foodLogItem.dbId));
      });
    }
  };

  const handleClickOpen = () => {
    setOpenDialog(true);
  };

  const handleClose = () => {
    setOpenDialog(false);
  };

  return (
    <Card sx={{ minWidth: MIN_WIDTH }}>
      <CardActionArea onClick={handleClickOpen}>
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
      </CardActionArea>
      <CardActions disableSpacing>
        <IconButton onClick={handleDelete}>
          <Delete />
        </IconButton>
      </CardActions>
      <EditFoodItemDialog
        onSubmit={onSubmit}
        openDialog={openDialog}
        handleClose={handleClose}
        currentWeight={foodLogItem.weight.toString()}
      />
    </Card>
  );
}

export default FoodItemLogCard;
