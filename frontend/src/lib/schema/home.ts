const websiteId = 'https://www.steaminputdb.com/#website';
const webPageId = 'https://www.steaminputdb.com/#webpage';
const searchActionId = 'https://www.steaminputdb.com/#search-action';

export function createHomeSchemaJsonLd() {
    return JSON.stringify({
        '@context': 'https://schema.org',
        '@graph': [
            {
                '@id': websiteId,
                '@type': 'WebSite',
                'name': 'SteamInputDB',
                'url': 'https://www.steaminputdb.com/',
                'description': 'Community-driven database of Steam Input configurations using the Steam API.',
                'inLanguage': 'en',
                'potentialAction': {
                    '@id': searchActionId,
                    '@type': 'SearchAction',
                    'target': 'https://www.steaminputdb.com/config/search?searchtext={searchtext}',
                    'query-input': 'required name=searchtext'
                }
            },
            {
                '@id': webPageId,
                '@type': 'WebPage',
                'name': 'SteamInputDB',
                'url': 'https://www.steaminputdb.com/',
                'description': 'Community-driven database of Steam Input configurations using the Steam API.',
                'isPartOf': {
                    '@id': websiteId
                },
                'primaryImageOfPage': {
                    '@type': 'ImageObject',
                    'url': 'https://www.steaminputdb.com/ogimage.png'
                }
            }
        ]
    });
}
