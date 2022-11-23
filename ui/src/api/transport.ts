import { createConnectTransport } from '@bufbuild/connect-web';

const transport = createConnectTransport({
  baseUrl: 'http://localhost:8080/api',
});

export default transport;
