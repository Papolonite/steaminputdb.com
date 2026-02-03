import { register } from '$lib/metrics';
import type { RequestHandler } from './$types';


export const GET: RequestHandler = async (event) => {
    const addr = event.getClientAddress();

    if (
        !isLoopbackAddress(addr)
        && !event.url.host.includes('docker')
        && !(event.url.hostname === ('frontend'))
    ) {
        return new Response('Forbidden', { status: 403 });
    }

    const metrics = await register.metrics();

    return new Response(metrics, {
        headers: {
            'Content-Type': register.contentType
        }
    });
};
function isLoopbackAddress(addr: string) {
    if (addr === '::1') {
        return true;
    }
    if (addr.startsWith('127.')) {
        return true;
    }
    if (addr.startsWith('::ffff:127.')) {
        return true;
    }
    return false;
}
