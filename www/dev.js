import { mount } from 'svelte'
import { createInertiaApp } from '@inertiajs/svelte'

const pages = import.meta.glob('./pages/**/*.svelte')

createInertiaApp({
    resolve: async (name) => pages[`./pages/${name}.svelte`](),
    setup: ({ el, App, props }) => { mount(App, { target: el, props }) }
})
