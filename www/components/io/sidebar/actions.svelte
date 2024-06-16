<script>
    import io from '$/stores/io.svelte.js'
    import { contextMenu } from '$/stores/portals.svelte.js'
    import ContextMenu from '$/components/atoms/context-menu.svelte'

    let { id, parent, share, settings } = $props()

    const actions = [
        { name: 'New', action: async () => await io.create(id, 'thread') },
        { name: 'Delete', action: async () => await io.remove(id, parent) },
        {
            name: 'Share',
            action: (event) => contextMenu.show(event.target, share, 'bottom', '', event),
        },
        {
            name: 'Settings',
            action: (event) => contextMenu.show(event.target, settings, 'bottom', '', event),
        },
    ]
</script>

<ContextMenu {actions} />
