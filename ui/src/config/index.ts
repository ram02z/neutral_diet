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

export const MIN_CARD_WIDTH = '250px';

export const MIN_CF_LIMIT = 0.1;

export const MAX_CF_LIMIT = 10.0;

export const NAVIGATION_DRAWER_WIDTH = '240px';
