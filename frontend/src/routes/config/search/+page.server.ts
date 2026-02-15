import { clientWithSvelteFetch } from '$lib/api/client';
import type { components } from '$lib/api/openapi';
import { log } from '$lib/log';
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
        let rankby = url.searchParams.get('sort-by');
        switch (rankby) {
            case 'trend':
                break;
            case 'playtime':
                rankby = 'lifetime_avg_playtime';
                break;
                // TODO:
        }
        if (!rankby) {
            rankby = 'trend';
        }

        const filterTags = Array.from(url.searchParams.entries())
            .map(([k ]) => k)
            .filter(
                (k) => k.startsWith('feature_')
            );

        const controller_filter = url.searchParams.get('controller_type');
        if (controller_filter) {
            filterTags.push(controller_filter);
        }

        const apiclient = clientWithSvelteFetch(fetch);
        try {
            const r = await apiclient.POST('/v1/search/configs', {
                body: {
                    limit: 20,
                    query_text: url.searchParams.get('searchtext') as string,
                    raw: false,
                    page: 1,
                    rank: {
                        by: rankby as 'trend',
                        trending_period: 30
                    },
                    filter: {
                        tags: filterTags
                    },
                    include: {
                        votes: true,
                        tags: true
                    }
                }
            });
                // log.debug('config search action', 'API response', r);
            if (r.error) {
                loadRes.searchError = {
                    status: r.error.status || 502,
                    message: r.error.title  || 'Failed to complete search',
                    ...r.error
                };
            }
            if (!r.data) {
                log.error('No data received from search endpoint');
                loadRes.searchError = {
                    status: 502,
                    message: 'No data received from search endpoint'
                };
            }
            loadRes.results = r.data;

        } catch (e) {
            loadRes.searchError = {
                status: 502,
                message: 'Error contacting search endpoint',
                error: `${e}`
            };
        }

    }

    return loadRes;
};
