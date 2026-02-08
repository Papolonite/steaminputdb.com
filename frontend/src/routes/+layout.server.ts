import { log } from '$lib/log';
import type { LayoutServerLoad } from './$types';


export const load: LayoutServerLoad = async ({ cookies }) => {
    const res = {
        theme: cookies.get('theme')
    };


    const token = cookies.get('token');
    if (!token) {
        return res;
    }
    const mid = token.split('.')?.[1];
    if (!mid) {
        return res;
    }
    const decoded = atob(mid);
    const payload = JSON.parse(decoded);
    const steamId = payload.sub as string | undefined;

    log.debug('Layout server load', 'steamid', steamId, 'payload', payload);

    return {
        ...res,
        steamId,
        userInfo: payload
    };

};


