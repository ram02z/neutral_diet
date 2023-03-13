import { Line } from 'react-chartjs-2';

import { useTheme } from '@mui/material';

import {
  CategoryScale,
  ChartData,
  Chart as ChartJS,
  ChartOptions,
  Legend,
  LineElement,
  LinearScale,
  PointElement,
  Title,
  Tooltip,
} from 'chart.js';
import { MAX_CF_LIMIT } from '@/config';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

const options: ChartOptions<'line'> = {
  responsive: true,
  spanGaps: true,
  plugins: {
    legend: {
      position: 'top' as const,
    },
  },
  scales: {
    y: {
      suggestedMin: 0,
      suggestedMax: MAX_CF_LIMIT,
    },
  },
};

type GoalLinePlotProps = {
  actualData: Record<string, number>;
  goal: number;
  title: string;
};

function GoalLinePlot({ actualData, goal, title }: GoalLinePlotProps) {
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

  const extOptions = { ...options, plugins: { title: { display: true, text: title } } };

  return <Line options={extOptions} data={data} />;
}

export default GoalLinePlot;
