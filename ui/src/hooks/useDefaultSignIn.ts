import { useCallback, useState } from 'react';
import { useSetRecoilState } from 'recoil';

import { Auth, UserCredential, signInWithEmailAndPassword } from 'firebase/auth';

import { CurrentUserDisplayName } from '@/store/user';

import { DefaultSignInHook } from './types';

function useDefaultSignIn(auth: Auth): DefaultSignInHook {
  const [error, setError] = useState<boolean>(false);
  const [loggedInUser, setLoggedInUser] = useState<UserCredential>();
  const setUserDisplayName = useSetRecoilState(CurrentUserDisplayName);
  const [loading, setLoading] = useState<boolean>(false);

  const signIn = useCallback(
    async (email: string, password: string) => {
      setLoading(true);
      setError(false);
      try {
        const userCredential = await signInWithEmailAndPassword(auth, email, password);
        setUserDisplayName(userCredential.user.displayName);
        setLoggedInUser(userCredential);

        return userCredential;
      } catch (err) {
        console.error(err);
        setError(true);
      } finally {
        setLoading(false);
      }
    },
    [auth, setUserDisplayName],
  );

  return [signIn, loggedInUser, loading, error];
}

export default useDefaultSignIn;
