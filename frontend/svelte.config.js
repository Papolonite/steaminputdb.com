import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    compilerOptions: {
        runes: true,
    },
    kit: {
        adapter: adapter({
            pages: '../backend/frontend/dist',
            assets: '../backend/frontend/dist',
            fallback: 'index.html'
        }),
    },
};

export default config;
