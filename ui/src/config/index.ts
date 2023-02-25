import { Notifications } from './types';

export const notifications: Notifications = {
  options: {
    anchorOrigin: {
      vertical: 'bottom',
      horizontal: 'center',
    },
    preventDuplicate: true,
  },
  maxSnack: 1,
};

export const MIN_WIDTH = '250px';

export const MAX_CF_LIMIT = 10.0;
