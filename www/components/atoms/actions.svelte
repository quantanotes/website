<script>
    import contextMenu from '$/stores/context-menu.svelte.js'

    /** @type {{ actions: { name: string, action: (event: Event) => void }[] }} */
    let { actions } = $props()

    function handleAction(event, action) {
        contextMenu.hide()
        action(event)
    }
</script>

<div class="bg-primary border border-secondary flex flex-col p-1 shadow-md w-32">
    {#each actions as action}
        <button
            class="flex flex-row hover:font-bold gap-6 px-4 text-left w-full"
            onclick={(event) => handleAction(event, action.action)}
        >
            {action.name}
        </button>
    {/each}
</div>
