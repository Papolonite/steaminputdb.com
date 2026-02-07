import { clientWithSvelteFetch } from '$lib/api/client';
import { log } from '$lib/log';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async ({ data  }) => {
    if (!data?.userInfo && data?.steamId) {
        try {
            const resp = await clientWithSvelteFetch(fetch).GET('/v1/steam/userinfo');
            if (resp.error) {
                log.error('Failed to fetch user info in layout load', 'status', resp.response.status, 'error', resp.error);
            }
            if (!resp.data) {
                log.error('No user info data received in layout load');
            }
            data.userInfo = resp.data;
        } catch (e) {
            log.error('Failed to fetch user info in layout load','err', e);
        }
    }
    return data || {};
};

