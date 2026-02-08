import { clientWithSvelteFetch, type ResponseType } from '$lib/api/client';
import type { components } from '$lib/api/openapi';
import { log } from '$lib/log';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';


export const load: PageServerLoad = async ({ params, fetch }) => {

    const fileid = params.fileid;

    const file_id = parseInt(fileid, 10);
    if (isNaN(file_id) || !file_id) {
        throw error(400, 'Invalid file ID');
    }

    let infoResp: Awaited<ResponseType<'GET', '/v1/steam/filedetails'>>;
    try {
        infoResp = await clientWithSvelteFetch(fetch).GET('/v1/steam/filedetails', {
            params: {
                query: {
                    file_id,
                    playtime_stats: 14,
                    raw: false
                }
            }
        });
    } catch (err) {
        log.error('Failed to fetch file details', 'file_id', file_id, 'error', err);
        throw error(500, {
            message: 'An unexpected error occurred while fetching file details',
            err
        });
    }
    if (infoResp.error) {
        log.error('Failed to fetch file details', 'file_id', file_id, 'error', infoResp.error);
        throw error(infoResp.error.status || 503, {
            message: infoResp.error.detail || 'Failed to fetch file details',
            error: infoResp.error
        });
    }
    if (!infoResp.data) {
        throw error(404, 'File not found');
    }

    const fileInfo = infoResp.data as components['schemas']['ConfigResponseItem'];
    const resData: {
        fileInfo: components['schemas']['ConfigResponseItem'];
        nonSteam: boolean;
        appInfo?: components['schemas']['AppInfo'];
    } = {
        fileInfo,
        nonSteam: ! (!!fileInfo.app_id && Number.isInteger(fileInfo.app_id))
    };

    if (fileInfo.app_id) {
        try {
            const appInfoResp = await clientWithSvelteFetch(fetch).GET('/v1/steam/appinfo', {
                params: {
                    query: {
                        app_id: fileInfo.app_id,
                        raw: false
                    }
                }
            });
            if (appInfoResp.error) {
                log.error(
                    'Failed to fetch app details',
                    'file_id', file_id,
                    'app_id', fileInfo.app_id,
                    'status', appInfoResp.error.status,
                    'error', appInfoResp.error
                );
            }
            if (!appInfoResp.data) {
                log.error('No app details data received',' file_id', file_id, 'app_id', fileInfo.app_id);
            }
            resData.appInfo = appInfoResp.data;
        } catch (err) {
            log.error('Failed to fetch app details', 'file_id', file_id, 'app_id', fileInfo.app_id, 'error', err);
        }
    }

    return resData;

};
