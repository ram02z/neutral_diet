import { Line } from 'react-chartjs-2';

import {
  CategoryScale,
  Chart as ChartJS,
  Legend,
  LineElement,
  LinearScale,
  PointElement,
  Title,
  Tooltip,
  ChartOptions,
  ChartData,
} from 'chart.js';
import { useTheme } from '@mui/material';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

const options: ChartOptions<'line'> = {
  responsive: true,
  spanGaps: true,
  plugins: {
    legend: {
      position: 'top' as const,
    },
  },
};

type GoalLinePlotProps = {
  actualData: Record<string, number>
  goal: number;

}

function GoalLinePlot({ actualData, goal }: GoalLinePlotProps) {
  const labels = Object.keys(actualData);
  const theme = useTheme();
  const data: ChartData<'line'> = {
    labels,
    datasets: [
      {
        label: 'Goal',
        data: labels.map(() => goal),
        borderColor: theme.palette.secondary.main,
        backgroundColor: theme.palette.secondary.dark,
      },
      {
        label: 'Actual',
        data: Object.values(actualData),
        borderColor: theme.palette.primary.main,
        backgroundColor: theme.palette.primary.dark,
      },
    ],
  };

  return (<Line options={options} data={data}/>);
}

export default GoalLinePlot;
