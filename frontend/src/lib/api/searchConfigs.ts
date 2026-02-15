import { clientWithSvelteFetch } from '$lib/api/client';
import type { components } from '$lib/api/openapi';
import { error } from '@sveltejs/kit';

export const PAGE_SIZE = 20;

export const fetchConfigs = async (
    fetch: typeof globalThis.fetch,
    params: URLSearchParams | FormData
): Promise<components['schemas']['ConfigsResponse']> => {
    if (Array.from(params.keys()).length > 0) {
        const rankby = params.get('sort-by');
        const pageParam = params.get('page')?.toString();
        const pageNum = Math.max(1, Number.parseInt(pageParam ?? '1', 10) || 1);

        const filterTags = Array.from(params.entries())
            .map(([k ]) => k)
            .filter(
                (k) => k.startsWith('feature_')
            );

        const controller_filter = params.get('controller_type');
        if (controller_filter) {
            filterTags.push(controller_filter.toString());
        }

        const apiclient = clientWithSvelteFetch(fetch);
        const r = await apiclient.POST('/v1/search/configs', {
            body: {
                limit: PAGE_SIZE,
                query_text: (params.get('searchtext') ?? '') as string,
                raw: false,
                page: pageNum,
                rank: {
                    by: rankby as 'vote',
                    trending_period: 30
                },
                filter: {
                    tags: filterTags,
                    app_id: params.get('appid')?.toString() ?? undefined
                },
                include: {
                    votes: true,
                    tags: true
                }
            }
        });
        // log.debug('config search action', 'API response', r);
        if (r.error) {
            error(r.error.status || 502, {
                message: r.error.title || 'Failed to complete search',
                ...r.error
            });
        }
        if (!r.data) {
            error(502, {
                message: 'No data received from search endpoint'
            });
        }
        return r.data;


    }
    return {
    };
};
