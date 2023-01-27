import { ComponentType, StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { HelmetProvider } from 'react-helmet-async';
import { RecoilRoot } from 'recoil';

import { ConfirmProvider } from 'material-ui-confirm';
import { SnackbarProvider } from 'notistack';

import { notifications } from '@/config';
import ThemeProvider from '@/theme/Provider';

const container = document.getElementById('root') as HTMLElement;
const root = createRoot(container);

function render(App: ComponentType) {
  root.render(
    <StrictMode>
      <RecoilRoot>
        <HelmetProvider>
          <ThemeProvider>
            <ConfirmProvider>
              <SnackbarProvider
                maxSnack={notifications.maxSnack}
                anchorOrigin={notifications.options.anchorOrigin}
              >
                <App />
              </SnackbarProvider>
            </ConfirmProvider>
          </ThemeProvider>
        </HelmetProvider>
      </RecoilRoot>
    </StrictMode>,
  );
}

export default render;
