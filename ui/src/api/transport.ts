import { createConnectTransport } from '@bufbuild/connect-web';

const transport = createConnectTransport({
  baseUrl: import.meta.env.VITE_BACKEND_BASE_URL,
});

export default transport;
