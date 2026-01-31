import { PUBLIC_API_BASE_URL, PUBLIC_API_BASE_URL_LOCAL, PUBLIC_DEV } from '$env/static/public';
import createClient from 'openapi-fetch';
import type { paths } from './openapi';


export const client = createClient<paths>({
    baseUrl: PUBLIC_DEV ? PUBLIC_API_BASE_URL_LOCAL : PUBLIC_API_BASE_URL
});
