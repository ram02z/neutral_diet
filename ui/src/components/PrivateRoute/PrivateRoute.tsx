import { FC } from 'react';
import { Navigate } from 'react-router';

import { useCurrentUser } from '@/hooks/useCurrentUser';
import routes from '@/routes';
import { Pages } from '@/routes/types';
import { useSetRecoilState } from 'recoil';
import { PrivateRoutePathState } from '@/store/location';

const PrivateRoute: FC<{ children: JSX.Element }> = ({ children }) => {
  const user = useCurrentUser();
  const setPath = useSetRecoilState(PrivateRoutePathState);
  setPath(user === null ? location.pathname : null)

  return user === null ? (
    <Navigate to={routes[Pages.Auth].path} replace/>
  ) : (
    children
  );
};

export default PrivateRoute;
