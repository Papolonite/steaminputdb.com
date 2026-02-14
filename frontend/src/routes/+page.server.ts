import { resolve } from '$app/paths';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions = {
    search: async (event) => {
        const params = await event.request.formData();
        throw redirect(302, resolve( `/config/search?term=${params.get('searchtext')}`));
    }
} satisfies Actions;
