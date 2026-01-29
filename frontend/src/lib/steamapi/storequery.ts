import { GET } from './client';
import type { CStoreQuerySearchSuggestionsRequest, CStoreQuerySearchSuggestionsResponse } from './proto/service_storequery.pb';

export function SearchSuggestions(
    req: CStoreQuerySearchSuggestionsRequest,
    fetchFn: typeof fetch = fetch,
    apiKey?: string
): Promise<CStoreQuerySearchSuggestionsResponse> {
    return GET(
        {
            interface: 'IStoreQueryService',
            method: 'SearchSuggestions'
        },
        req,
        fetchFn,
        apiKey
    );
}
