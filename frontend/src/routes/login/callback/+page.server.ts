import { clientWithSvelteFetch } from '$lib/api/client';
import { log } from '$lib/log';
import { error, fail, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ url }) => {

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
        error(400, { message: 'Missing OpenID parameters' });
    }


    return {};
};


// ouh the lengths one goes through to have
// SSR,
// an HTTP-only cookie,
// no API-Key exposed,
// AND an immediate redirect to the frontend in case the real backend is slow...
// 🙄
export const actions = {
    validateLogin: async ({ request, cookies, fetch }) => {
        if (!request.body) {
            log.error('No request body in validateLogin action');
            return fail(400, 'No body in request');
        }

        try {
            const r = await clientWithSvelteFetch(fetch).POST('/v1/steam/login', {
                body: await request.json()
            });

            if (r.error) {
                log.error('Login endpoint returned an error',  'error', r.error);
                return fail(r.error.status || 502, {
                    ...r.error,
                    message: r.error.title || 'Failed to complete Steam login'
                });
            }
            if (!r.data) {
                log.error('No data received from login endpoint');
                return fail(502, { message: 'No data received from login endpoint' });
            }


            const cookieHeaders = r.response.headers.getSetCookie();
            cookieHeaders?.forEach((header) => {
                const parts = header.split(';').map((part) => part.trim());
                const [nameValue, ...attributes] = parts;
                if (!nameValue) {
                    return;
                }
                const nameValueIndex = nameValue.indexOf('=');
                const name = nameValueIndex === -1 ? nameValue : nameValue.slice(0, nameValueIndex);
                const value = nameValueIndex === -1 ? '' : nameValue.slice(nameValueIndex + 1);

                const options: Record<string, string | boolean> = {};
                attributes.forEach((attr) => {
                    const [attrName, attrValue] = attr.split('=');
                    if (!attrName) {
                        return;
                    }
                    options[attrName.toLowerCase()] = attrValue ? attrValue : true;
                });

                const sameSite = typeof options.samesite === 'string'
                    ? options.samesite.toLowerCase()
                    : undefined;
                const expires = typeof options.expires === 'string'
                    ? new Date(options.expires)
                    : undefined;
                const maxAge = typeof options['max-age'] === 'string'
                    ? Number.parseInt(options['max-age'], 10)
                    : undefined;

                cookies.set(name, value, {
                    path: typeof options.path === 'string' ? options.path : '/',
                    domain: typeof options.domain === 'string' ? options.domain : undefined,
                    httpOnly: options.httponly === true,
                    secure: options.secure === true,
                    sameSite: sameSite === 'lax' || sameSite === 'strict' || sameSite === 'none'
                        ? sameSite
                        : undefined,
                    expires: expires && !Number.isNaN(expires.getTime()) ? expires : undefined,
                    maxAge: Number.isFinite(maxAge) ? maxAge : undefined
                });
            });

            return r.data;
        } catch (err) {
            log.error('Error contacting login endpoint', 'error', err);
            return fail(502, { message: 'Error contacting login endpoint', error: `${err}` });
        }
    }
} satisfies Actions;
