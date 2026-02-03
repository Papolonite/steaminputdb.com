import { PUBLIC_API_BASE_URL, PUBLIC_API_BASE_URL_LOCAL, PUBLIC_DEV } from '$env/static/public';
import createClient, { type FetchResponse } from 'openapi-fetch';
import type { paths } from './openapi';

export const client = createClient<paths>({
    baseUrl: PUBLIC_DEV ? PUBLIC_API_BASE_URL_LOCAL : PUBLIC_API_BASE_URL
});


export type ResponseType<
    M extends keyof typeof client,
    P extends keyof paths,
> = P extends keyof paths
    ? Lowercase<M> extends keyof paths[P]
        ? paths[P][Lowercase<M>] extends Record<string | number, unknown>
            ? FetchResponse<paths[P][Lowercase<M>], unknown, `${string}/${string}`>
            : never
        : never
    : never;
