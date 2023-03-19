import { useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useRecoilValue, useSetRecoilState } from 'recoil';

import { Button } from '@mui/material';

import dayjs from 'dayjs';

import client from '@/api/user_service';
import { FormValues } from '@/components/AddGoalDialog/types';
import Insights from '@/core/insights';
import {
  CurrentUserHeadersState,
  LocalUserSettingsState,
  UserGoalsState,
  UserInsightsState,
} from '@/store/user';
import { toSerializableDate } from '@/utils/date';

import RecommendGoalDialog from './RecommendGoalDialog';

function RecommendGoalsButton() {
  const setUserGoals = useSetRecoilState(UserGoalsState);
  const userHeaders = useRecoilValue(CurrentUserHeadersState);
  const userInsights = useRecoilValue(UserInsightsState);
  const userSettings = useRecoilValue(LocalUserSettingsState);
  const [openRecommendDialog, setOpenRecommendDialog] = useState(false);
  const externalInsights = new Insights(userSettings.dietaryRequirement).insights;

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    const target = parseFloat(data.targetCarbonFootprint);
    const today = dayjs();
    client
      .addCarbonFootprintGoal(
        {
          carbonFootprintGoal: {
            description: data.description,
            startDate: toSerializableDate(today),
            endDate: toSerializableDate(data.endDate),
            startCarbonFootprint: userInsights.userDailyAverage,
            targetCarbonFootprint: target,
          },
        },
        { headers: userHeaders },
      )
      .then((res) => {
        setUserGoals((old) => {
          return {
            completed: old.completed,
            active: [
              ...old.active,
              {
                dbId: res.id,
                description: data.description,
                startDate: today,
                endDate: data.endDate,
                startCarbonFootprint: userInsights.userDailyAverage,
                targetCarbonFootprint: target,
              },
            ],
          };
        });
      });
    handleCloseRecommendDialog();
  };

  const handleOpenRecommendDialog = () => {
    setOpenRecommendDialog(true);
  };

  const handleCloseRecommendDialog = () => {
    setOpenRecommendDialog(false);
  };

  return (
    <>
      <Button variant="contained" onClick={handleOpenRecommendDialog}>
        Recommend a goal
      </Button>
      <RecommendGoalDialog
        openDialog={openRecommendDialog}
        handleClose={handleCloseRecommendDialog}
        onSubmit={onSubmit}
        startCarbonFootprint={userInsights.userDailyAverage}
        goals={externalInsights}
      />
    </>
  );
}

export default RecommendGoalsButton;
