import { useCallback, useState } from 'react';

import {
  Auth,
  AuthProvider,
  CustomParameters,
  GoogleAuthProvider,
  UserCredential,
  signInWithPopup,
} from 'firebase/auth';

import client from '@/api/user_service';

import { SignInWithPopupHook } from './types';

export const useSignInWithGoogle = (auth: Auth): SignInWithPopupHook => {
  const createGoogleAuthProvider = useCallback(
    (scopes?: string[], customOAuthParameters?: CustomParameters) => {
      const provider = new GoogleAuthProvider();
      if (scopes) {
        scopes.forEach((scope) => provider.addScope(scope));
      }
      if (customOAuthParameters) {
        provider.setCustomParameters(customOAuthParameters);
      }
      return provider;
    },
    [],
  );
  return useSignInWithPopup(auth, createGoogleAuthProvider);
};

const useSignInWithPopup = (
  auth: Auth,
  createProvider: (scopes?: string[], customOAuthParameters?: CustomParameters) => AuthProvider,
): SignInWithPopupHook => {
  const [error, setError] = useState<boolean>(false);
  const [loggedInUser, setLoggedInUser] = useState<UserCredential>();
  const [loading, setLoading] = useState<boolean>(false);

  const doSignInWithPopup = useCallback(
    async (scopes?: string[], customOAuthParameters?: CustomParameters) => {
      setLoading(true);
      setError(false);
      try {
        const provider = createProvider(scopes, customOAuthParameters);
        const user = await signInWithPopup(auth, provider);
        // Add user to backend DB
        client.createUser({ firebaseUid: user.user.uid });
        setLoggedInUser(user);

        return user;
      } catch (err) {
        console.error(err);
        setError(true);
      } finally {
        setLoading(false);
      }
    },
    [auth, createProvider],
  );

  return [doSignInWithPopup, loggedInUser, loading, error];
};
