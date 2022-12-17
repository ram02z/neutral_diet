import { signOut } from 'firebase/auth';

import Loading from '@/components/Loading';
import SignIn from '@/components/SignIn';
import SignUp from '@/components/SignUp';
import { auth } from '@/core/firebase';
import useAuthState from '@/hooks/useAuthState';

const logout = () => {
  signOut(auth);
};

function Account() {
  const [user, loading, error] = useAuthState(auth);

  if (loading) {
    return (
      <>
        <Loading />
      </>
    );
  }

  if (error) {
    return (
      <div>
        <p>Error:</p>
        <>{error}</>
      </div>
    );
  }
  if (user) {
    return (
      <div>
        <p>Current User: {user.displayName}</p>
        <button onClick={logout}>Log out</button>
      </div>
    );
  }

  return (
    <>
      <SignIn /> <SignUp />
    </>
  );
}

export default Account;
