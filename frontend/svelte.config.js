// import adapter from '@sveltejs/adapter-static';
import adapter from '@sveltejs/adapter-node';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    compilerOptions: {
        runes: true,
        experimental: {
            async: true
        }
    },
    kit: {
        adapter: adapter({
            // pages: '../backend/frontend/dist',
            // assets: '../backend/frontend/dist',
            // fallback: 'index.html'
        })
    },
};

export default config;
