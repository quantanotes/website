import { render } from 'svelte/server'
import createServer from '@westacks/inertia-svelte/server'
import { createInertiaApp } from '@westacks/inertia-svelte'
import { resolvePageComponent } from 'laravel-vite-plugin/inertia-helpers'

createServer((page) => createInertiaApp({
    page,
    resolve: (name) => resolvePageComponent(`./pages/${name}.svelte`, import.meta.glob('./pages/**/*.svelte')),
    setup: ({ App, props }) => render(App, { props })
}))