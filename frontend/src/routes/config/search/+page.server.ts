import type { components } from '$lib/api/openapi';
import { log } from '$lib/log';
import { isHttpError } from '@sveltejs/kit';
import { fetchConfigs } from '../../../lib/api/searchConfigs';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({  url, fetch }) => {
    log.debug('Config search page load', 'searchParams', Array.from(url.searchParams.entries()));

    const loadRes: {
        hasSearched?: boolean;
        results?: components['schemas']['ConfigsResponse'];
        searchError?: {
            status?: number;
            message?: string;
        } & Record<string, unknown>;
    } = {};

    if (url.searchParams.size > 0) {
        loadRes.hasSearched = true;

        try {
            loadRes.results = await fetchConfigs(fetch, url.searchParams);
        } catch (e) {
            log.error('Error fetching search results', 'error', e);
            if (isHttpError(e)) {
                loadRes.searchError = {
                    status: e.status || 502,
                    message: e.body?.message || 'Error contacting search endpoint',
                    error: `${e.body}`
                };
            } else {
                loadRes.searchError = {
                    status: 502,
                    message: 'Error contacting search endpoint',
                    error: `${e}`
                };
            }
        }
    }

    return loadRes;
};
