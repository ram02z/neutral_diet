import { Fragment } from 'react';
import { BrowserRouter } from 'react-router-dom';

import CssBaseline from '@mui/material/CssBaseline';

import Pages from '@/routes/Pages';
import Navigation from '@/sections/Navigation';

import Header from './sections/Header';

function App() {
  return (
    <Fragment>
      <CssBaseline />
      <BrowserRouter>
        <Header title="Neutral Diet" />
        <Pages />
        <Navigation />
      </BrowserRouter>
    </Fragment>
  );
}

export default App;
