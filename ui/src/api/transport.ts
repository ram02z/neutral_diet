import { createConnectTransport } from '@bufbuild/connect-web';

const transport = createConnectTransport({
  baseUrl: `${import.meta.env.VITE_BACKEND_HOST}/api`,
});

export default transport;
