import { FC } from 'react';
import { PathRouteProps } from 'react-router-dom';

import type { SvgIconProps } from '@mui/material/SvgIcon';

enum Pages {
  Home,
  Diary,
  Search,
  Account,
  Auth,
  NotFound,
}

type SubComponent = {
  path: string;
  component: FC;
};

type PathRouteCustomProps = {
  title?: string;
  path: string;
  component: FC;
  icon?: FC<SvgIconProps>;
  navigation: boolean;
  requireAuth: boolean;
  subComponents: SubComponent[];
};

type Routes = Record<Pages, PathRouteProps & PathRouteCustomProps>;

export type { Routes };
export { Pages };
