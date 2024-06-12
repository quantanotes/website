import { hydrate } from 'svelte'
import { createInertiaApp } from '@westacks/inertia-svelte'

const pages = import.meta.glob('./pages/**/*.svelte', { eager: true })

createInertiaApp({
    resolve: (name) => pages[`./pages/${name}.svelte`],
    setup: ({ el, App, props }) => { hydrate(App, { target: el, props }) }
})