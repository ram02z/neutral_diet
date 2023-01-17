import { useState } from 'react';

import { auth } from '@/core/firebase';
import useDefaultSignIn from '@/hooks/useDefaultSignIn';

import Loading from '@/components/Loading';

function LogIn() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [signIn, _, loading, error] = useDefaultSignIn(auth);

  // TODO: handle errors by type
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
      <p>Welcome back</p>
      <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
      <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button onClick={() => signIn(email, password)}>Continue</button>
    </>
  );
}

export default LogIn;
