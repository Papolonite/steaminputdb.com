import { clientWithSvelteFetch } from '$lib/api/client';
import type { components } from '$lib/api/openapi';
import { log } from '$lib/log';
import { fail } from '@sveltejs/kit';
import type { Actions } from './$types';


export const actions = {
    search: async ({ request, fetch }) => {
        log.debug('config search action', 'request', request);
        const data = await request.formData();
        log.debug('config search action', 'form data', Array.from(data.entries()));

        let rankby = data.get('sort-by');
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

        const apiclient = clientWithSvelteFetch(fetch);

        let results: components['schemas']['ConfigsResponse'] = {};
        try {
            const r = await apiclient.POST('/v1/search/configs', {
                body: {
                    limit: 20,
                    query_text: data.get('searchtext') as string,
                    raw: false,
                    page: 1,
                    rank: {
                        by: rankby as 'trend',
                        trending_period: 14
                    },
                    filter: {
                        // TODO: fix types....
                        controller_type: data.get('controller-type') as 'controller_neptune' || undefined
                        // TODO: additional filters
                    },
                    include: {
                        votes: true
                    }
                }
            });
            log.debug('config search action', 'API response', r);
            if (r.error) {
                return fail(r.error.status || 502, {
                    ...r.error,
                    message: r.error.title || 'Failed to complete search'
                });
            }
            if (!r.data) {
                log.error('No data received from search endpoint');
                return fail(502, { message: 'No data received from search endpoint' });
            }
            results = r.data;

        } catch (e) {
            return fail(502, { message: 'Error contacting search endpoint', error: `${e}` });
        }

        return {
            results
        };
    }
} satisfies Actions;
