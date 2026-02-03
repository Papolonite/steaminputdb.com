
const steamAuthUrl = 'https://steamcommunity.com/openid/login';
export const buildSteamLoginUrl = (origin: string) => {
    const params = new URLSearchParams({
        'openid.mode': 'checkid_setup',
        'openid.ns': 'http://specs.openid.net/auth/2.0',
        'openid.identity': 'http://specs.openid.net/auth/2.0/identifier_select',
        'openid.claimed_id': 'http://specs.openid.net/auth/2.0/identifier_select',
        'openid.return_to': `${origin}/login/callback`,
        'openid.realm': origin
    });
    return `${steamAuthUrl}?${params.toString()}`;
};
