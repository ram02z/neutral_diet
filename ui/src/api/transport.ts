import { createConnectTransport } from '@bufbuild/connect-web';

const transport = createConnectTransport({
  baseUrl: `http://${import.meta.env.VITE_BACKEND_HOST}:${import.meta.env.VITE_BACKEND_PORT}/api`,
});

export default transport;
