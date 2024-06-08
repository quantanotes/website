<script>
    import contextMenu from '$/stores/context-menu.svelte.js'
    import io from '$/stores/io.svelte.js'
    import Actions from '$/components/atoms/actions.svelte'

    let { id, parent, share, settings } = $props()

    const actions = [
        { name: 'New', action: async () => await io.create(id) },
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

<Actions {actions} />
