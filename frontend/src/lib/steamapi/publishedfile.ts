import { GET } from './client';
import type {
    CPublishedFileGetDetailsRequest,
    CPublishedFileGetDetailsResponse,
    CPublishedFileQueryFilesRequest,
    CPublishedFileQueryFilesResponse
} from './proto/service_publishedfile.pb';

export function QueryFiles(
    req: CPublishedFileQueryFilesRequest,
    fetchFn: typeof fetch = fetch,
    apiKey?: string
): Promise<CPublishedFileQueryFilesResponse> {
    return GET(
        {
            interface: 'IPublishedFileService',
            method: 'QueryFiles'
        },
        req,
        fetchFn,
        apiKey
    );
}

export function GetFileDetails(
    req: CPublishedFileGetDetailsRequest,
    fetchFn: typeof fetch = fetch,
    apiKey?: string
): Promise<CPublishedFileGetDetailsResponse> {
    return GET(
        {
            interface: 'IPublishedFileService',
            method: 'GetDetails'
        },
        req,
        fetchFn,
        apiKey
    );
}
