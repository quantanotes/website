import { SvelteComponent, tick } from 'svelte'

/** @type {undefined | { source: any, component: SvelteComponent, placement: string, style: string | undefined }} */
let contextMenu = $state(undefined)

export default {
    get contextMenu() {
        return contextMenu
    },

    async show(source, component, placement = 'bottom', style = undefined, event = undefined) {
        contextMenu = undefined
        await tick()
        event?.stopPropagation()
        contextMenu = {
            source,
            component,
            placement,
            style,
        }
    },

    hide() {
        contextMenu = undefined
    }
}