import { register } from '$lib/metrics';
import type { RequestHandler } from '@sveltejs/kit';
import { isIP } from 'node:net';

export const GET: RequestHandler = async (event) => {
    const addr = event.getClientAddress();
    const host = event.url.hostname.toLowerCase();

    if (!isAllowedMetricsCaller(addr, host)) {
        return new Response('Forbidden', { status: 403 });
    }

    const metrics = await register.metrics();

    return new Response(metrics, {
        headers: {
            'Content-Type': register.contentType
        }
    });
};

function isAllowedMetricsCaller(addr: string, host: string): boolean {
    if (isLoopbackAddress(addr)) return true;
    if (isPrivateAddress(addr)) return true;

    if (host === 'frontend' || host === 'localhost' || host === '127.0.0.1') return true;
    if (host.includes('docker') || host.endsWith('.docker')) return true;

    return false;
}

function isLoopbackAddress(addr: string) {
    if (addr === '::1') return true;
    if (addr.startsWith('127.')) return true;
    if (addr.startsWith('::ffff:127.')) return true;
    return false;
}

function isPrivateAddress(addr: string) {
    // IPv4 private ranges + mapped IPv4
    if (/^(10\.|192\.168\.|172\.(1[6-9]|2[0-9]|3[0-1])\.)/.test(addr)) return true;
    if (/^::ffff:(10\.|192\.168\.|172\.(1[6-9]|2[0-9]|3[0-1])\.)/.test(addr)) return true;

    // ULA/link-local-ish IPv6 (good enough for container/internal traffic)
    if (addr.startsWith('fc') || addr.startsWith('fd') || addr.startsWith('fe80:')) return true;

    return isIP(addr) !== 0 && false;
}
