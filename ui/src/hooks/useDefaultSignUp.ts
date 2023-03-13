import { useCallback, useState } from 'react';
import { useSetRecoilState } from 'recoil';

import { Auth, UserCredential, createUserWithEmailAndPassword } from 'firebase/auth';

import client from '@/api/user_service';
import { CurrentUserDisplayName } from '@/store/user';

import { DefaultSignUpHook } from './types';

function useDefaultSignUp(auth: Auth): DefaultSignUpHook {
  const [error, setError] = useState<boolean>(false);
  const [registeredUser, setRegisteredUser] = useState<UserCredential>();
  const setUserDisplayName = useSetRecoilState(CurrentUserDisplayName);
  const [loading, setLoading] = useState<boolean>(false);

  const signUp = useCallback(
    async (displayName: string, email: string, password: string) => {
      setLoading(true);
      setError(false);
      try {
        const userCredential = await createUserWithEmailAndPassword(auth, email, password);
        setUserDisplayName(displayName);
        // Add user to backend DB
        userCredential.user.getIdToken().then((idToken) => {
          const headers = new Headers();
          headers.set('Authorization', `Bearer ${idToken}`);
          client.createUser({ firebaseUid: userCredential.user.uid }, { headers: headers });
        });

        setRegisteredUser(userCredential);

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

  return [signUp, registeredUser, loading, error];
}

export default useDefaultSignUp;
