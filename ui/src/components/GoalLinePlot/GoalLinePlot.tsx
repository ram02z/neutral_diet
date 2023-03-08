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
  const data: ChartData<'line'> = {
    labels,
    datasets: [
      {
        label: 'Goal',
        data: labels.map(() => goal),
        borderColor: 'rgb(255, 99, 132)',
        backgroundColor: 'rgba(255, 99, 132, 0.5)',
      },
      {
        label: 'Actual',
        data: Object.values(actualData),
        borderColor: 'rgb(255, 99, 132)',
        backgroundColor: 'rgba(255, 99, 132, 0.5)',
      },
    ],
  };

  return (<Line options={options} data={data}/>);
}

export default GoalLinePlot;
