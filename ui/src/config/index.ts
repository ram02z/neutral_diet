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

export { notifications };
