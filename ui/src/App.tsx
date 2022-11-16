import { Fragment } from 'react';
import { BrowserRouter } from 'react-router-dom';

import CssBaseline from '@mui/material/CssBaseline';

import Header from '@/sections/Header';
import Pages from '@/routes/Pages';
import Navigation from '@/sections/Navigation';

function App() {
  return (
    <Fragment>
      <CssBaseline />
      <BrowserRouter>
        <Header />
        <Pages />
        <Navigation />
      </BrowserRouter>
    </Fragment>
  );
}

export default App;
