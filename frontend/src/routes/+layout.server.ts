import type { LayoutServerLoad } from './$types';


export const load: LayoutServerLoad = async ({ parent, cookies }) => {
    const res = {
        ...(await parent()),
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

    return {
        ...res,
        steamId
    };

};
