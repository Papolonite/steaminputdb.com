import { log } from '$lib/log';
import type { Actions } from './$types';

export const actions = {
    search: async (event) => {
        // TODO register the user
        log.debug('search', event);
        return {};
    }
} satisfies Actions;
