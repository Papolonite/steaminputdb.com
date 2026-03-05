import { resolve } from '$app/paths';
import { redirect, type RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async (event) => {
    event.cookies.delete('token', { path: '/' });

    const tld = event.url.hostname.split('.').slice(-2).join('.');
    event.cookies.set('token', '', {
        path: '/',
        domain: `${tld}`,
        httpOnly:  true,
        secure:  true,
        sameSite: 'lax',
        expires: undefined,
        maxAge: 1
    });

    throw redirect(303, resolve('/'));
};
