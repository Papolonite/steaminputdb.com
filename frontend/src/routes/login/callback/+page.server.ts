import { client as apiclient, type ResponseType } from '$lib/api/client';
import { log } from '$lib/log';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ parent, url }) => {
    const parentData = await parent();

    const openIdParams = url.searchParams.entries().reduce(
        (acc, [k, v]) => {
            if (k.startsWith('openid.')) {
                acc[k] = v;
                return acc;
            } else {
                return acc;
            }
        },
        {} as Record<string, string>
    );
    if (!Object.keys(openIdParams).length) {
        log.error('Missing OpenID parameters in login callback');
        throw error(400, { message: 'Missing OpenID parameters' });
    }

    const loginPromise = (async () => {
        let r: Awaited<ResponseType<'POST', '/v1/steam/login'>>;
        try {
            r = await apiclient.POST('/v1/steam/login', {
                body: Object.entries(openIdParams).reduce((acc, [k, v]) => {
                    const key = k.split('.')?.pop();
                    if (!key) {
                        return acc;
                    }
                    acc[key] = v;
                    return acc;
                    // eslint-disable-next-line @typescript-eslint/no-explicit-any
                }, {} as any)
            });
        } catch (err) {
            log.error('Error contacting login endpoint', 'error', err);
            throw error(500, { message: 'Error contacting login endpoint', error: `${err}` });
        }
        if (r.error) {
            log.error('Login endpoint returned an error',  'error', r.error);
            throw error(r.error.status || 500, {
                ...r.error,
                message: r.error.title || 'Failed to complete Steam login'
            });
        }
        if (!r.data) {
            log.error('No data received from login endpoint');
            throw error(500, { message: 'No data received from login endpoint' });
        }
        if (!r.data.token) {
            log.error('Invalid login response from server: missing token');
            throw error(500, { message: 'Invalid login response from server' });
        }
        const mid = r.data.token.split('.')?.[1];
        if (!mid) {
            log.error('Invalid JWT token received from login endpoint');
            throw error(500, { message: 'Invalid JWT token received' });
        }
        const decoded = atob(mid);
        const expiresIn = JSON.parse(decoded).exp - Math.floor(Date.now() / 1000);
        // cookies.set('token', r.data.token, { path: '/', maxAge: expiresIn });

        // throw redirect(302, '/');
        return `token=${r.data.token}; Path=/; Max-Age=${expiresIn}; SameSite=Lax`;
    })();

    return {
        ...parentData,
        loginPromise
    };
};
