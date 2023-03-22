import { Fragment } from 'react';
import { BrowserRouter } from 'react-router-dom';

import CssBaseline from '@mui/material/CssBaseline';

import Pages from '@/routes/Pages';
import Navigation from '@/sections/Navigation';
import NotificationService from "@/core/notifications";

import Header from './sections/Header';

function App() {
  window.addEventListener("load", async () => {
    await NotificationService.registerWorker();
    const permission = await Notification.requestPermission();
    if (permission == "denied") {
      console.error("permission denied")
      return
    }
    const token = await NotificationService.getWebToken();
    console.log(token);
    NotificationService.onNotifications((notification) => {
      console.log(notification.messageId);
    })
  })
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
