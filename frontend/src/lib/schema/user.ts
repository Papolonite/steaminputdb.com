export function createUserSchemaJsonLd(params: {
    userId: string;
    personName: string;
    imageUrl?: string;
}) {
    const pageUrl = `https://www.steaminputdb.com/user/${params.userId}`;

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
                'name': `SteamInputDB - User: ${params.personName}`,
                'url': pageUrl,
                'description': `Search for Steam Input configurations from ${params.personName}`,
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
