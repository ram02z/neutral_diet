import { Notifications } from './types';

const notifications: Notifications = {
  options: {
    anchorOrigin: {
      vertical: 'bottom',
      horizontal: 'center',
    },
  },
  maxSnack: 1,
  preventDuplicate: true,
};

const MIN_WIDTH = '250px';

export { notifications, MIN_WIDTH };
