import Loading from '../Loading';
import { auth } from '@/core/firebase';
import { useState } from 'react';
import useDefaultSignIn from '@/hooks/useDefaultSignIn';

function SignIn() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [signIn, user, loading, error] = useDefaultSignIn(auth);

  if (error) {
    return (
      <div>
        <p>Error: {error.message}</p>
      </div>
    );
  }

  if (loading) {
    return <Loading />;
  }

  return (
    <>
      <p>Please sign-in:</p>
      <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
      <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button onClick={() => signIn(email, password)}>Sign In</button>
    </>
  );
}

export default SignIn;
