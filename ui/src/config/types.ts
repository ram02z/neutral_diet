import { SnackbarProps } from 'notistack';

type Notifications = {
  options: SnackbarProps;
  maxSnack: number;
  preventDuplicate: boolean;
};

export type { Notifications };
