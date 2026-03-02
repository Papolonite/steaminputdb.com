import type { Handle } from '@sveltejs/kit';
import { randomUUID } from 'node:crypto';

import { building } from '$app/environment';
import { ANSI, log } from '$lib/log';
import { getStatusText, httpActiveConnections, httpRequestDuration, httpRequestsTotal } from '$lib/metrics';
import { sequence } from '@sveltejs/kit/hooks';


const logHook: Handle = async ({ event, resolve }) => {
    if (building) {
        return resolve(event);
    }
    const start = Date.now();
    const requestId = randomUUID();
    let statusCode = 500;

    try {
        httpActiveConnections.inc();
        const response = await resolve(event);
        statusCode = response.status;
        response.headers.set('x-request-id', requestId);
        return response;
    } finally {
        const durationMs = Date.now() - start;
        const durationSeconds = durationMs / 1000;

        const method = event.request.method;
        const path = event.url.pathname;
        const statusText = getStatusText(statusCode);

        httpRequestsTotal.labels(method, path, statusText).inc();
        httpRequestDuration.labels(method, path).observe(durationSeconds);

        const remoteAddr = event.getClientAddress?.() ?? 'unknown';
        const statusStr = String(statusCode).charAt(0);
        let statusColor: string;
        switch (statusStr) {
            case '2':
                statusColor = ANSI.green;
                break;
            case '3':
                statusColor = ANSI.cyan;
                break;
            case '4':
                statusColor = ANSI.yellow;
                break;
            case '5':
                statusColor = ANSI.red;
                break;
            default:
                statusColor = ANSI.white;
        }

        log.debug(
            'request',
            'status_code',
            `${statusColor}${statusCode}${ANSI.reset}`,
            'method',
            event.request.method,
            'path',
            event.url.pathname,
            'duration_ms',
            durationMs,
            'request_id',
            requestId,
            'remote_addr',
            remoteAddr
        );
        httpActiveConnections.dec();
    }
};

const themeHook: Handle = async ({ event, resolve }) => {
    const theme = event.cookies.get('theme');
    if (!theme) {
        return resolve(event);
    }
    return resolve(event,
        {
            transformPageChunk: ({ html }) => html
                .replace('<html ', `<html data-theme="${theme}" style="color-scheme: ${theme}"`)
        }
    );
};

export const handle: Handle = sequence(
    logHook,
    themeHook
);
