import { FC } from 'react';
import { Navigate } from 'react-router';
import { useSetRecoilState } from 'recoil';

import { useCurrentUser } from '@/hooks/useCurrentUser';
import routes from '@/routes';
import { Pages } from '@/routes/types';
import { PrivateRoutePathState } from '@/store/location';

const PrivateRoute: FC<{ children: JSX.Element }> = ({ children }) => {
  const user = useCurrentUser();
  const setPath = useSetRecoilState(PrivateRoutePathState);
  if (user === null) {
    setPath(location.pathname);
    return <Navigate to={routes[Pages.Auth].path} replace />;
  }

  return children;
};

export default PrivateRoute;
