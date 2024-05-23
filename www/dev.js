import { mount } from 'svelte'
import { createInertiaApp } from '@westacks/inertia-svelte'
import { resolvePageComponent } from 'laravel-vite-plugin/inertia-helpers'

createInertiaApp({
    resolve: name => resolvePageComponent(`./pages/${name}.svelte`, import.meta.glob('./pages/**/*.svelte')),
    setup({ el, App, props }) {
        mount(App, { target: el, props })
    }
})
