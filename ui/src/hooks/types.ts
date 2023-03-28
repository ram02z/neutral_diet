import dayjs from 'dayjs';
import { CustomParameters, UserCredential } from 'firebase/auth';

export type AuthActionHook<M> = [M, UserCredential | undefined, boolean, boolean];

export type DefaultSignInHook = AuthActionHook<
  (email: string, password: string) => Promise<UserCredential | undefined>
>;

export type DefaultSignUpHook = AuthActionHook<
  (displayName: string, email: string, password: string) => Promise<UserCredential | undefined>
>;

export type SignInWithPopupHook = AuthActionHook<
  (
    scopes?: string[],
    customOAuthParameters?: CustomParameters,
  ) => Promise<UserCredential | undefined>
>;

export type HighlightedDaysHook = [
  (date: dayjs.Dayjs) => Promise<number[]>,
  number[],
  boolean,
  React.MutableRefObject<AbortController | null>,
];
