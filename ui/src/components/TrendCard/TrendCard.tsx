import TrendingDownIcon from '@mui/icons-material/TrendingDown';
import TrendingFlatIcon from '@mui/icons-material/TrendingFlat';
import TrendingUpIcon from '@mui/icons-material/TrendingUp';
import { Card, CardContent, Grid, Stack, Typography } from '@mui/material';

type TrendCardProps = {
  title: string;
  stat: number;
  today: number;
  source: string;
};

function TrendCard({ title, stat, today, source }: TrendCardProps) {
  let statDiff = 0;
  if (today > 0) {
    statDiff = ((stat - today) / today) * 100.0;
  }
  let textColor = 'text.secondary';
  if (statDiff > 0) {
    textColor = 'success.main';
  } else if (statDiff < 0) {
    textColor = 'error.main';
  }
  return (
    <Card>
      <CardContent>
        <Grid container columns={5}>
          <Stack spacing={1} sx={{ pt: 1, pl: 1 }}>
            <Typography sx={{ textTransform: 'capitalize' }} variant="h6">
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
                {statDiff > 0 && <TrendingUpIcon color="success" />}
                {statDiff == 0 && <TrendingFlatIcon color="secondary" />}
                {statDiff < 0 && <TrendingDownIcon color="error" />}
              </Grid>
              <Grid>
                <Typography variant="subtitle1" color={textColor}>
                  <b>{Math.abs(statDiff).toFixed(2)}%</b>
                </Typography>
              </Grid>
            </Grid>
            <Typography variant="caption" color="text.secondary">
              <b>Source:</b> {source}
            </Typography>
          </Stack>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default TrendCard;
