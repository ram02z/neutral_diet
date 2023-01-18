import { useState } from 'react';

import Loading from '@/components/Loading';
import { auth } from '@/core/firebase';
import useDefaultSignUp from '@/hooks/useDefaultSignUp';

function SignUp() {
  const [displayName, setDisplayName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [signUp, _, loading, error] = useDefaultSignUp(auth);

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
      <p>Create your account</p>
      <input
        placeholder="Name"
        type="text"
        value={displayName}
        onChange={(e) => setDisplayName(e.target.value)}
      />
      <input
        placeholder="Email"
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <input
        placeholder="Password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <button onClick={() => signUp(displayName, email, password)}>Continue</button>
    </>
  );
}

export default SignUp;
