import { useCallback, useState } from 'react';

import {
  Auth,
  AuthError,
  UserCredential,
  createUserWithEmailAndPassword,
  updateProfile,
} from 'firebase/auth';

import client from '@/api/user_service';

import { DefaultSignUpHook } from './types';

function useDefaultSignUp(auth: Auth): DefaultSignUpHook {
  const [error, setError] = useState<AuthError>();
  const [registeredUser, setRegisteredUser] = useState<UserCredential>();
  const [loading, setLoading] = useState<boolean>(false);

  const signUp = useCallback(
    async (displayName: string, email: string, password: string) => {
      setLoading(true);
      setError(undefined);
      try {
        const user = await createUserWithEmailAndPassword(auth, email, password);
        updateProfile(user.user, { displayName: displayName });
        // Add user to backend DB
        client.createUser({ firebaseUid: user.user.uid });

        setRegisteredUser(user);

        return user;
      } catch (err) {
        setError(err as AuthError);
      } finally {
        setLoading(false);
      }
    },
    [auth],
  );

  return [signUp, registeredUser, loading, error];
}

export default useDefaultSignUp;
