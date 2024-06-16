import { SvelteComponent, tick } from 'svelte'

export const contextMenu = createPortalStore()
export const dialog = createPortalStore()

function createPortalStore() {
    /** @type {undefined | { source: any, component: SvelteComponent, placement: string, style: string | undefined }} */
    let portal = $state(undefined)

    return {
        get portal() {
            return portal
        },

        async show(source, component, placement = 'bottom', style = undefined, event = undefined) {
            portal = undefined
            await tick()
            event?.stopPropagation()
            portal = {
                source,
                component,
                placement,
                style,
            }
        },

        hide() {
            portal = undefined
        },
    }

}