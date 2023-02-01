import { useCallback, useState } from 'react';

import { Auth, UserCredential, signInWithEmailAndPassword } from 'firebase/auth';

import { DefaultSignInHook } from './types';

function useDefaultSignIn(auth: Auth): DefaultSignInHook {
  const [error, setError] = useState<boolean>(false);
  const [loggedInUser, setLoggedInUser] = useState<UserCredential>();
  const [loading, setLoading] = useState<boolean>(false);

  const signIn = useCallback(
    async (email: string, password: string) => {
      setLoading(true);
      setError(false);
      try {
        const user = await signInWithEmailAndPassword(auth, email, password);
        setLoggedInUser(user);

        return user;
      } catch (err) {
        console.error(err);
        setError(true);
      } finally {
        setLoading(false);
      }
    },
    [auth],
  );

  return [signIn, loggedInUser, loading, error];
}

export default useDefaultSignIn;
