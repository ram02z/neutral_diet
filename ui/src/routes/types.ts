import { FC } from 'react';
import { PathRouteProps } from 'react-router-dom';

import type { SvgIconProps } from '@mui/material/SvgIcon';

enum Pages {
  Home,
  Diary,
  Search,
  Account,
  LogIn,
  SignUp,
  NotFound,
}

type PathRouteCustomProps = {
  title?: string;
  path: string;
  component: FC;
  icon?: FC<SvgIconProps>;
  navigation: boolean;
};

type Routes = Record<Pages, PathRouteProps & PathRouteCustomProps>;

export type { Routes };
export { Pages };
