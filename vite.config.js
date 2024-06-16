import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import laravel from 'laravel-vite-plugin'

const dev = process.env.APP_ENV === 'DEV'

export default defineConfig({
    base: '/static/build',
    build: {
        bundle: true,
        emptyOutDir: true,
        minify: true,
        modulePreload: false,
        sourcemap: dev,
    },
    plugins: [
        svelte({}),
        laravel({
            input: [`www/${dev ? 'dev' : 'app'}.js`, 'www/app.css'],
            publicDirectory: 'static',
            refresh: ['www/**'],
            ssr: 'www/ssr.js',
            ssrOutputDirectory: 'ssr',
        }),
    ],
    resolve: {
        alias: { $: './www' },
    },
})
