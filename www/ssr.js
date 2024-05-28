import { render } from 'svelte/server'
import createServer from '@inertiajs/svelte/server'
import { createInertiaApp } from '@inertiajs/svelte'

createServer(page => createInertiaApp({
    page,
    resolve(name) {
        const pages = import.meta.glob('./pages/**/*.svelte', { eager: true })
        return pages[`./pages/${name}.svelte`]
    },
    setup: ({ App, props }) => render(App, { props })
}))