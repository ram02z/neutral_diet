import { useEffect, useState } from 'react';

import { Auth, User, onIdTokenChanged } from 'firebase/auth';

import { IdTokenHook } from './types';

export default (auth: Auth): IdTokenHook => {
  const [error, setError] = useState<Error>();
  const [user, setUser] = useState<User | null>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const listener = onIdTokenChanged(
      auth,
      async (user) => {
        setLoading(true);
        setUser(user);
        console.log(user);
        setLoading(false);
      },
      setError,
    );

    return () => {
      listener();
    };
  }, [auth]);

  return [user, loading, error];
};
