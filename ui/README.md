# UI

## Getting started

### Prerequisites

- [Node.js](https://nodejs.org/en)
- Firebase Project

### Installation

```sh
npm i
```

### Environment

The application includes `.env.development` files with sample working
development environment variables. Not all required environment variables are
included in `.env.development`, i.e. Firebase environment variables.

- `VITE_FIREBASE_APP_ID`
- `VITE_FIREBASE_API_KEY`
    - [Docs](https://firebase.google.com/docs/projects/api-keys)
- `VITE_VAPID_KEY`
    - [Docs](https://firebase.google.com/docs/cloud-messaging/js/client#configure_web_credentials_with)
- `VITE_MESSAGING_SENDER_ID`
    - [Docs](https://firebase.google.com/docs/cloud-messaging/concept-options#credentials)
- `VITE_GOOGLE_CLOUD_PROJECT`
    - [Docs](https://firebase.google.com/docs/projects/learn-more#project-id)
- `VITE_FIREBASE_AUTH_DOMAIN`
- `VITE_GA_MEASUREMENT_ID`

Most the Firebase environment variables can be found using this
short [guide](https://stackoverflow.com/a/52500964). Refer to
[.env.example](./.env.example) for the complete list of required environment
variables.

### Available Scripts

### `npm run dev`

Start a local web server with HMR in development mode.
Open [http://localhost:5173](http://localhost:5173) to view it in your browser.

### `npm run build`

Builds the application for production to the `dist` folder.

### `npm run preview`

Starts a local web server that serves the `dist` folder in production mode.
Open [http://localhost:4173](http://localhost:5173) to view it in your browser.

### `npm run test`

Runs Playwright end-to-end tests. May require additional browser installations.
