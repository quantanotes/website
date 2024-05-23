<script>
    import contextMenu from '$/stores/context-menu.svelte.js'
    import notes from '$/stores/notes.svelte.js'
    import Actions from '$/components/atoms/actions.svelte'

    let { id, parent, share } = $props() 

    const actions = [
        { name: 'New', action: async () => await notes.create(id) },
        { name: 'Delete', action: async () => await notes.remove(id, parent) },
        {
            name: 'Share',
            action: (event) => contextMenu.show(event.target, share, 'bottom', '', event),
        },
    ]
</script>

<Actions {actions} />
