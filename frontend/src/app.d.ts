import 'unplugin-icons/types/svelte';


// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
    namespace App {
        interface Error {
            message: string;
            [key: string]: unknown;
        }
        // interface Locals {}
        interface PageData {
            steamId?: string;
            theme?: 'light' | 'dark';
        }
        // interface PageState {}
        // interface Platform {}
    }
}

export { };

