import type Module from 'module';
import type { Keys as ProtoNames, Types as ProtoTypes } from './proto';
import type { MessageFns } from './proto/common.pb';

// magic! man a love typescript(s typesystem - typescript itself not so much)!

export type RequestTypeNames = Extract<ProtoNames, `${string}Request`>;
export type ResponseTypeNames = Extract<ProtoNames, `${string}Response`>;

export type ServiceInterfaces = Extract<ProtoNames, `${string}ServiceName`>;

type BuildServiceName<S extends string> =
    S extends `${infer Service}Name`
        ? `I${Service}`
        : never;

export type ServiceTypeNames = BuildServiceName<ServiceInterfaces>;

type ServiceTypeNameToServiceName<T extends ServiceTypeNames> =
    T extends `I${infer S}Service` ? S : never;

type ServiceNames = ServiceTypeNameToServiceName<ServiceTypeNames>;


type MethodsForService<S extends string> =
    Extract<RequestTypeNames, `C${S}${string}Request`> extends `C${S}${infer Method}Request`
        ? Method
        : never;

type ServiceEndpoints = {
    [S in ServiceNames]: {
        interface: `I${S}Service`;
        method: MethodsForService<S>;
    }
}[ServiceNames];

export type Endpoint = ServiceEndpoints & {
    version?: string;
};

type ProtoTypeName<S extends string, Method extends string, Suffix extends 'Request' | 'Response'> =
    `C${S}${Method}${Suffix}`;


type UnwrapMessageFns<T> = T extends MessageFns<infer Inner> ? Inner : T;

type GetProtoMessageType<S extends string, Method extends string, Suffix extends 'Request' | 'Response'> =
    ProtoTypeName<S, Method, Suffix> extends ProtoNames
        ? UnwrapMessageFns<ProtoTypes[ProtoTypeName<S, Method, Suffix>]>
        : never;

export type RequestType<
    IFace extends ServiceTypeNames,
    Method extends MethodsForService<ServiceTypeNameToServiceName<IFace>>
> =
    IFace extends `I${infer S}Service`
        ? GetProtoMessageType<S, Method, 'Request'>
        : never;

export type ResponseType<
    IFace extends ServiceTypeNames,
    Method extends MethodsForService<ServiceTypeNameToServiceName<IFace>>
> =
    IFace extends `I${infer S}Service`
        ? GetProtoMessageType<S, Method, 'Response'>
        : never;

export type RequestTypeEP<EP extends Endpoint> =
    EP['interface'] extends `I${infer S}Service`
        ? GetProtoMessageType<S, EP['method'], 'Request'>
        : never;

export type ResponseTypeEP<EP extends Endpoint> =
    EP['interface'] extends `I${infer S}Service`
        ? GetProtoMessageType<S, EP['method'], 'Response'>
        : never;

export type X = RequestType<'IStoreQueryService', 'SearchSuggestions'>;
export type Y = ResponseType<'IStoreQueryService', 'SearchSuggestions'>;
export type XY = ResponseTypeEP<{ interface: 'IStoreQueryService'; method: 'SearchSuggestions' }>;

type ProtoMap = {
    [K in ProtoNames]: ProtoTypes[K];
};

const protos = import.meta.glob('./proto/*pb.ts');
const protoTypes: ProtoMap = (await Promise.all(
    Object.entries(protos).map(
        ([file, importFn]) => {
            if (file.startsWith('./common')) {
                return;
            }
            return importFn() as Promise<Module>;
        }
    )
)).reduce((acc: ProtoMap, module) => {
    if (!module) {
        return acc;
    }
    return Object.assign(acc, module);
}, {} as ProtoMap);

export function CtorForTypeString<T extends ProtoNames>(type: T): ProtoMap[T] {
    const res = protoTypes[type];
    if (!res || typeof res !== 'object' || !('encode' in res) || !('decode' in res)) {
        throw new Error('NoPrototypeError', {
            cause: `No proto type found for string: ${type}`
        });
    }
    return protoTypes[type];
}

export function BuildRequestTypeString<EP extends Endpoint>(ep: EP): RequestTypeNames {
    return `${
        ep.interface.replace(/^I/, 'C')
            .replace(/Service$/, '')
    }${
        ep.method
    }Request` as RequestTypeNames;

}
export function BuildResponseTypeString<EP extends Endpoint>(ep: EP): ResponseTypeNames {
    return `${
        ep.interface.replace(/^I/, 'C')
            .replace(/Service$/, '')
    }${
        ep.method
    }Response` as ResponseTypeNames;
}
