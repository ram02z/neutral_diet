import { createPromiseClient } from '@bufbuild/connect-web';

import { FoodService } from '@/api/gen/neutral_diet/food/v1/api_connectweb';

import transport from './transport';

const client = createPromiseClient(FoodService, transport);

export default client;
