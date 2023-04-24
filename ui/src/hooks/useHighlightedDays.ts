import { useCallback, useRef, useState } from 'react';
import { useRecoilValue } from 'recoil';

import dayjs from 'dayjs';

import client from '@/api/user_service';
import { CurrentUserHeadersState } from '@/store/user';

import { HighlightedDaysHook } from './types';

/**
 * Get days in current month where food has been logged.
 */
function useHighlightedDays(): HighlightedDaysHook {
  const [loading, setLoading] = useState(false);
  const [highlightedDays, setHighlightedDays] = useState<number[]>([]);
  const requestAbortController = useRef<AbortController | null>(null);
  const userHeaders = useRecoilValue(CurrentUserHeadersState);

  const fetchHighlightedDays = useCallback(
    async (date: dayjs.Dayjs) => {
      const controller = new AbortController();
      setHighlightedDays([]);
      setLoading(true);
      try {
        const response = await client.getFoodItemLogDays(
          { year: date.year(), month: date.month() + 1 },
          { headers: userHeaders },
        );

        setHighlightedDays(response.days);

        return response.days;
      } catch (err) {
        console.error(err);
        setLoading(false);
        return [];
      } finally {
        setLoading(false);
        requestAbortController.current = controller;
      }
    },
    [userHeaders],
  );

  return [fetchHighlightedDays, highlightedDays, loading, requestAbortController];
}

export default useHighlightedDays;
