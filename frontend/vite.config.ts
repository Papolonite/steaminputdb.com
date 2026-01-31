import { sveltekit } from '@sveltejs/kit/vite';
import { playwright } from '@vitest/browser-playwright';
import { existsSync } from 'fs';
import Icons from 'unplugin-icons/vite';
import { defineConfig } from 'vitest/config';

const chromiumPath = existsSync('/usr/bin/chromium') ? '/usr/bin/chromium' : undefined;

export default defineConfig({
    plugins: [
        sveltekit(),
        Icons({
            compiler: 'svelte',
            autoInstall: true
        })
    ],
    test: {
        expect: { requireAssertions: true },
        projects: [
            {
                extends: './vite.config.ts',
                test: {
                    name: 'client',
                    browser: {
                        enabled: true,
                        provider: playwright({
                            launchOptions: chromiumPath ? { executablePath: chromiumPath } : {}
                        }),
                        instances: [{ browser: 'chromium', headless: process.env.TEST_SHOW_BROWSER ? false : true }]
                    },
                    include: ['src/**/*.svelte.{test,spec}.{js,ts}'],
                    exclude: ['src/lib/server/**']
                }
            },

            {
                extends: './vite.config.ts',
                test: {
                    name: 'server',
                    environment: 'node',
                    include: ['src/**/*.{test,spec}.{js,ts}'],
                    exclude: ['src/**/*.svelte.{test,spec}.{js,ts}']
                }
            }
        ]
    }
});
