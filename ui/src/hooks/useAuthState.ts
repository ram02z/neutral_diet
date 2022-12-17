import { useEffect, useState } from 'react';

import { Auth, AuthError, User, onAuthStateChanged } from 'firebase/auth';

import { AuthStateHook } from './types';

type AuthStateOptions = {
  onUserChanged?: (user: User | null) => Promise<void>;
};

export default (auth: Auth, options?: AuthStateOptions): AuthStateHook => {
  const [error, setError] = useState<Error>();
  const [user, setUser] = useState<User | null>();
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    const listener = onAuthStateChanged(
      auth,
      async (user) => {
        setError(undefined);
        if (options?.onUserChanged) {
          setLoading(true);
          // onUserChanged function to process custom claims on any other trigger function
          try {
            await options.onUserChanged(user);
          } catch (e) {
            setError(e as AuthError);
          } finally {
            setLoading(false);
          }
        }
        setUser(user);
      },
      setError,
    );

    return () => {
      listener();
    };
  }, [auth, options]);

  return [user, loading, error];
};
