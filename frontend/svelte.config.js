// import adapter from '@sveltejs/adapter-static';
import adapter from '@sveltejs/adapter-node';
import { mdsvex } from "mdsvex";
import remarkGfm from "remark-gfm";

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
    extensions: [".svelte", ".svx"],
    preprocess: mdsvex({
        remarkPlugins: [
            [remarkGfm, {

            }]
        ],
    })
};

export default config;
