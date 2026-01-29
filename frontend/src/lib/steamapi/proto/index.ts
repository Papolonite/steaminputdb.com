import * as publishedfile from './service_publishedfile.pb';
import * as storequery from './service_storequery.pb';


export type Types = typeof storequery & typeof publishedfile;
export type Keys = keyof Types;
