import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ parent }) => {

    const { steamId } = await parent();

    if (!steamId) {
        throw redirect(302, '/');
    }
    throw redirect(302, `/user/${steamId}`);

};
