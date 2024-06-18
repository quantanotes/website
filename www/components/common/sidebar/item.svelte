<script>
    import Icon from '@iconify/svelte'
    import DropdownMenu from '$/components/atoms/dropdown-menu.svelte'

    let { id, state, contextMenu } = $props()
    let expanded = $state(false)
    let children = $derived(state.map[id]?.children.toSorted(sorter).filter(_filter) || [])
    // Counts nested drag enters
    let dragEnterCount = $state(0)
    // Checks if we need to expand due to drag and drop
    let dragExpanded = $derived(
        dragEnterCount > 0 && !(state.moving === id || state.ascendants(id).includes(state.moving)),
    )

    function expand() {
        if (!expanded) state.children(id)
        expanded = !expanded
    }

    function handleDragStart(event) {
        event.stopPropagation()
        state.moving = id
        event.dataTransfer.setData('text/plain', id)
    }

    function handleDragEnter(event) {
        event.preventDefault()
        dragEnterCount++
        if (!expanded && dragExpanded && dragEnterCount == 1) {
            state.dragover = id
            state.children(id)
        }
    }

    function handleDragLeave(event) {
        event.preventDefault()
        dragEnterCount--
    }

    function handleDragOver(event) {
        event.preventDefault()
    }

    async function handleDrop(event) {
        event.preventDefault()
        // Copy so we can reset (avoiding a long if statement below)
        const moving = state.moving
        const dragover = state.dragover
        state.moving = ''
        state.dragover = ''
        dragEnterCount = 0
        if (!moving || !dragover) return
        // If a parent is being dropped in to a child
        if (state.ascendants(id).includes(moving)) return
        // If it's dropping in to the same parent
        if (dragover === state.map[moving].parent) return
        // Make sure the UI is expanded so we can see where we dropped
        expanded = true
        if (dragover === moving) return
        await state.move(moving, id)
    }

    function sorter(a, b) {
        return state.map[a]?.title.localeCompare(state.map[b]?.title || '') || false
    }

    // We need to filter the child nodes for certain node categories as they are not to be displayed
    function _filter(node) {
        return !['message', 'link', 'record'].includes(state.map[node].category)
    }
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
    ondragstart={handleDragStart}
    ondragenter={handleDragEnter}
    ondragleave={handleDragLeave}
    ondragover={handleDragOver}
    ondrop={handleDrop}
    draggable="true"
>
    <div class="flex gap-2">
        <button class="h-[24px] icon-btn" onclick={expand}>
            {#if expanded || dragExpanded}
                <Icon icon="mdi:minus" />
            {:else}
                <Icon icon="mdi:add" />
            {/if}
        </button>

        <button
            class={`hover:font-bold transition text-left truncate w-full ${(id == state.current || dragEnterCount > 0) && 'font-bold'}`}
            onclick={() => state.open(id)}
        >
            {state.map[id]?.title}
        </button>

        <DropdownMenu class="h-[24px] icon-btn">
            <Icon icon="mdi:dots-horizontal" />
            {#snippet menu()}
                {@render contextMenu(id, state.map[id].parent)}
            {/snippet}
        </DropdownMenu>
    </div>

    {#if expanded || dragExpanded}
        <div class="ml-2">
            {#each children as id}
                <svelte:self {id} {state} {contextMenu} />
            {/each}
        </div>
    {/if}
</div>
