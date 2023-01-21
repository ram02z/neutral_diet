import { AuthError, UserCredential } from 'firebase/auth';

export type AuthActionHook<M> = [M, UserCredential | undefined, boolean, AuthError | undefined];

export type DefaultSignInHook = AuthActionHook<
  (email: string, password: string) => Promise<UserCredential | undefined>
>;

export type DefaultSignUpHook = AuthActionHook<
  (displayName: string, email: string, password: string) => Promise<UserCredential | undefined>
>;
