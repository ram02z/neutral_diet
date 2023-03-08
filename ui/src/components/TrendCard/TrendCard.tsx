import { Card, CardContent, Grid, Stack, Typography } from '@mui/material';
import TrendingUpIcon from '@mui/icons-material/TrendingUp';
import TrendingDownIcon from '@mui/icons-material/TrendingDown';

type TrendCardProps = {
  title: string;
  stat: number;
  today: number;
};

function TrendCard({ title, stat, today }: TrendCardProps) {
  const statDiff = (stat - today) / today * 100.0
  return (
    <Card sx={{ width: 310 }}>
      <CardContent>
        <Grid container columns={5}>
          <Stack spacing={1} sx={{ pt: 1, pl: 1 }}>
            <Typography sx={{ textTransform: 'capitalize' }} variant="h5" component="div">
              {title}
            </Typography>
            <Typography variant="subtitle1" color="text.secondary">
              <b>{parseFloat(stat.toFixed(3))}</b>
              <Typography variant="caption">
                CO<sub>2</sub>/kg
              </Typography>
            </Typography>
            <Grid container direction="row" alignItems="center">
            <Grid sx={{ pr: 1 }}>
            {statDiff > 0 ? <TrendingUpIcon color="success"/> : <TrendingDownIcon color="error"/>}
            </Grid>
            <Grid>
            <Typography variant="subtitle1" color={statDiff > 0 ? "success" : "error"}>
              <b>{Math.abs(statDiff).toFixed(2)}%</b>
            </Typography>
            </Grid>
            </Grid>
          </Stack>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default TrendCard;
