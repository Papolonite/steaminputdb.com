import { BuildRequestTypeString, BuildResponseTypeString, CtorForTypeString, type Endpoint, type RequestTypeEP, type ResponseTypeEP } from './magic';


export async function GET<EP extends Endpoint, Req extends RequestTypeEP<EP>>(
    endpoint: EP,
    req?: Req,
    fetchFn: typeof fetch = fetch,
    apiKey?: string
): Promise<ResponseTypeEP<EP>> {
    try {
        const params = new URLSearchParams();
        if (req) {
            const requestCtor = CtorForTypeString(BuildRequestTypeString(endpoint));
            const reqBytes = requestCtor.encode(req as never).finish(); // TODO: fix type or whatever
            params.append('input_protobuf_encoded', btoa(String.fromCharCode(...reqBytes)));
        }
        if (apiKey) {
            params.append('key', apiKey);
        }
        const URL = `https://api.steampowered.com/${endpoint.interface}/${endpoint.method}/v${endpoint.version ?? '1'}/?${params.toString()}`;
        const resp = await fetchFn(URL, {
            method: 'GET'
        });
        if (!resp.ok) {
            throw new Error('Steam API HTTP Error', {
                cause: {
                    status: resp.status,
                    statusText: resp.statusText
                }
            });
        }
        const buf = new Uint8Array(await resp.arrayBuffer());
        const respCtor = CtorForTypeString(BuildResponseTypeString(endpoint));
        return respCtor.decode(buf) as ResponseTypeEP<EP>;
    } catch (e) {
        throw new Error('Failed to fetch from Steam API', { cause: e });
    }
}
