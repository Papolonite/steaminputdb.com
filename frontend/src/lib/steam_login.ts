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
