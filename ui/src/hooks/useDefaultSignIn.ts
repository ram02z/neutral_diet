import { useCallback, useState } from 'react';

import {
  Auth,
  AuthError,
  UserCredential,
  signInWithEmailAndPassword,
} from 'firebase/auth';
import { DefaultSignInHook } from './types';


function useDefaultSignIn(auth: Auth): DefaultSignInHook {
  const [error, setError] = useState<AuthError>();
  const [loggedInUser, setLoggedInUser] = useState<UserCredential>();
  const [loading, setLoading] = useState<boolean>(false);

  const signIn = useCallback(
    async (email: string, password: string) => {
      setLoading(true);
      setError(undefined);
      try {
        const user = await signInWithEmailAndPassword(auth, email, password);
        setLoggedInUser(user);

        return user;
      } catch (err) {
        setError(err as AuthError);
      } finally {
        setLoading(false);
      }
    },
    [auth],
  );

  return [signIn, loggedInUser, loading, error];
}

export default useDefaultSignIn;
