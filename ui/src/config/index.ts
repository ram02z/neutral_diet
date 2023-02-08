import { Notifications } from './types';

const notifications: Notifications = {
  options: {
    anchorOrigin: {
      vertical: 'bottom',
      horizontal: 'center',
    },
    preventDuplicate: true,
  },
  maxSnack: 1,
};

const MIN_WIDTH = '250px';

export { notifications, MIN_WIDTH };
