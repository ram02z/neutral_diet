import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import BookIcon from '@mui/icons-material/Book';
import HomeIcon from '@mui/icons-material/Home';
import SearchIcon from '@mui/icons-material/Search';

import asyncComponentLoader from '@/utils/loader';

import { Pages, Routes } from './types';

const routes: Routes = {
  [Pages.Home]: {
    component: asyncComponentLoader(() => import('@/pages/Home')),
    path: '/',
    title: 'Home',
    icon: HomeIcon,
    navigation: true,
    requireAuth: true,
    subComponents: [],
  },
  [Pages.Diary]: {
    component: asyncComponentLoader(() => import('@/pages/Diary')),
    path: '/diary',
    title: 'Diary',
    icon: BookIcon,
    navigation: true,
    requireAuth: true,
    subComponents: [],
  },
  [Pages.Search]: {
    component: asyncComponentLoader(() => import('@/pages/Search')),
    path: '/search',
    title: 'Search',
    icon: SearchIcon,
    navigation: true,
    requireAuth: true,
    subComponents: [],
  },
  [Pages.Account]: {
    component: asyncComponentLoader(() => import('@/pages/Account')),
    path: '/account',
    title: 'Account',
    icon: AccountCircleIcon,
    navigation: false,
    requireAuth: true,
    subComponents: [
      {
        component: asyncComponentLoader(() => import('@/pages/Settings')),
        path: 'settings',
      },
      {
        component: asyncComponentLoader(() => import('@/pages/Goals')),
        path: 'goals',
      },
    ],
  },
  [Pages.Auth]: {
    component: asyncComponentLoader(() => import('@/pages/Auth')),
    path: '/auth',
    navigation: false,
    requireAuth: false,
    subComponents: [
      {
        component: asyncComponentLoader(() => import('@/pages/LogIn')),
        path: 'login',
      },
      {
        component: asyncComponentLoader(() => import('@/pages/SignUp')),
        path: 'signup',
      },
    ],
  },
  [Pages.NotFound]: {
    component: asyncComponentLoader(() => import('@/pages/NotFound')),
    path: '*',
    navigation: false,
    requireAuth: false,
    subComponents: [],
  },
};

export default routes;
