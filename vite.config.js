import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import laravel from 'laravel-vite-plugin'

const dev = process.env.APP_ENV === 'DEV'

export default defineConfig({
    build: {
        bundle: true,
        minify: true,
        emptyOutDir: true,
    },
    plugins: [
        svelte(),
        laravel({
            input: [`www/${dev ? 'dev' : 'app'}.js`, 'www/app.css'],
            ssr: 'www/ssr.js',
            buildDirectory: 'build',
            publicDirectory: 'static',
            ssrOutputDirectory: 'ssr',
            refresh: ['www/**'],
        }),
    ],
    resolve: { alias: { '$': './www', }, },
})
