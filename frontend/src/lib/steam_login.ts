import { browser } from '$app/environment';

const steamAuthUrl = 'https://steamcommunity.com/openid/login';
export const buildSteamLoginUrl = () => {
    const params = new URLSearchParams({
        'openid.mode': 'checkid_setup',
        'openid.ns': 'http://specs.openid.net/auth/2.0',
        'openid.identity': 'http://specs.openid.net/auth/2.0/identifier_select',
        'openid.claimed_id': 'http://specs.openid.net/auth/2.0/identifier_select',
        'openid.return_to': `${window.location.origin}/login/callback`,
        'openid.realm': window.location.origin
    });
    return `${steamAuthUrl}?${params.toString()}`;
};


export const steamIdFromToken = async () => {
    if (!browser) {
        return undefined;
    }

    const token = (await cookieStore.get('token'))?.value;
    if (!token) {
        return undefined;
    }

    const mid = token.split('.')?.[1];
    if (!mid) {
        return undefined;
    }
    const decoded = atob(mid);
    const payload = JSON.parse(decoded);
    const steamId = payload.sub as string | undefined;
    return steamId;
};
