import { render } from 'svelte/server'
import createServer from '@westacks/inertia-svelte/server'
import { createInertiaApp } from '@westacks/inertia-svelte'

const pages = import.meta.glob('./pages/**/*.svelte', { eager: true })

createServer((page) => createInertiaApp({
    page,
    resolve: (name) => pages[`./pages/${name}.svelte`],
    setup: ({ App, props }) => render(App, { props })
}))