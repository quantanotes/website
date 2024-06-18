<script>
    import { dialog } from '$/stores/portals.svelte.js'
    import notes from '$/stores/notes.svelte.js'
    import ContextMenu from '$/components/atoms/context-menu.svelte'

    let { id, parent, share } = $props()

    const actions = [
        { name: 'New', action: async () => await notes.create(id, 'note') },
        { name: 'Delete', action: async () => await notes.remove(id, parent) },
        {
            name: 'Share',
            action: (event) => dialog.show(event.target, share, 'bottom', '', event),
        },
    ]
</script>

<ContextMenu {actions} />
