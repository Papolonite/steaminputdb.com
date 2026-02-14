import { clientWithSvelteFetch, type ResponseType } from '$lib/api/client';
import type { components } from '$lib/api/openapi';
import { log } from '$lib/log';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';


export const load: PageServerLoad = async ({ params, fetch }) => {

    const appid = params.appid;

    const app_id = parseInt(appid, 10);
    if (isNaN(app_id) || !app_id) {
        error(400, 'Invalid app ID');
    }

    const client = clientWithSvelteFetch(fetch);
    let infoResp: Awaited<ResponseType<'GET', '/v1/steam/appinfo'>>;
    try {
        infoResp = await client.GET('/v1/steam/appinfo', {
            params: {
                query: {
                    app_id,
                    raw: false
                }
            }
        });
    } catch (err) {
        log.error('Failed to fetch app details', 'app_id', app_id, 'error', err);
        error(500, {
            message: 'An unexpected error occurred while fetching app details',
            err
        });
    }
    if (infoResp.error) {
        log.error('Failed to fetch app details', 'app_id', app_id, 'error', infoResp.error);
        error(infoResp.error.status || 503, {
            message: infoResp.error.detail || 'Failed to fetch app details',
            error: infoResp.error
        });
    }
    if (!infoResp.data || !infoResp.data.name) {
        error(404, 'App not found');
    }

    return {
        appInfo: infoResp.data as components['schemas']['AppsSearchItem']
    };

};
