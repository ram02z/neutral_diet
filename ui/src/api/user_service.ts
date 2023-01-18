import { createPromiseClient } from '@bufbuild/connect-web';

import { UserService } from '@/api/gen/neutral_diet/user/v1/api_connectweb';

import transport from './transport';

const client = createPromiseClient(UserService, transport);

export default client;
