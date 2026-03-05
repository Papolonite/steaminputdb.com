export function createAppSchemaJsonLd(params: {
    appId: string;
    appName: string;
    imageUrl?: string;
}) {
    const pageUrl = `https://www.steaminputdb.com/app/${params.appId}`;

    return JSON.stringify({
        '@context': 'https://schema.org',
        '@graph': [
            {
                '@id': 'https://www.steaminputdb.com/#website',
                '@type': 'WebSite',
                'name': 'SteamInputDB',
                'url': 'https://www.steaminputdb.com/',
                'potentialAction': {
                    '@id': `${pageUrl}#search-action`,
                    '@type': 'SearchAction',
                    'target': `${pageUrl}?searchtext={searchtext}`,
                    'query-input': 'required name=searchtext'
                }
            },
            {
                '@id': `${pageUrl}#webpage`,
                '@type': 'WebPage',
                'name': `SteamInputDB - ${params.appName}`,
                'url': pageUrl,
                'description': `Search for Steam Input configurations for ${params.appName}`,
                'isPartOf': {
                    '@id': 'https://www.steaminputdb.com/#website'
                },
                ...(params.imageUrl
                    ? {
                        primaryImageOfPage: {
                            '@type': 'ImageObject',
                            'url': params.imageUrl
                        }
                    }
                    : {})
            }
        ]
    });
}
