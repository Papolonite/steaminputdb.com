import { resolve } from '$app/paths';
import { redirect, type RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async (event) => {
    event.cookies.delete('token', { path: '/' });

    throw redirect(303, resolve('/'));
};
