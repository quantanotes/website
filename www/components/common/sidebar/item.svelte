<script>
    import { fade } from 'svelte/transition'
    import Icon from '@iconify/svelte'
    import Dropdown from '$/components/atoms/dropdown.svelte'

    let { id, state, contextMenu } = $props()
    let expanded = $state(false)
    let children = $derived(state.map[id]?.children.toSorted(sorter) || [])
    // Counts nested drag enters
    let dragRC = $state(0)
    // Checks if we need to expand due to drag and drop
    let dragExpanded = $derived(dragRC > 0 && !(state.moving === id || state.ascendants(id).includes(state.moving)))

    async function expand() {
        if (!expanded) await state.children(id)
        expanded = !expanded
    }

    function handleDragStart(event) {
        event.stopPropagation()
        state.moving = id
        event.dataTransfer.setData('text/plain', id)
    }

    async function handleDragEnter(event) {
        event.preventDefault()
        dragRC++
        if (!expanded && dragExpanded && dragRC == 1) {
            state.dragover = id
            await state.children(id)
        }
    }

    function handleDragLeave(event) {
        event.preventDefault()
        dragRC--
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
        dragRC = 0
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
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
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

        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <button
            class={`hover:font-bold transition text-left truncate w-full ${(id == state.current || dragRC > 0) && 'font-bold'}`}
            onclick={() => state.open(id)}
        >
            {state.map[id]?.title}
        </button>

        <Dropdown class="h-[24px] icon-btn">
            <Icon icon="mdi:dots-horizontal" />
            {#snippet menu()}
                {@render contextMenu(id, state.map[id].parent)}
            {/snippet}
        </Dropdown>
    </div>

    {#if expanded || dragExpanded}
        <div
            class="ml-2"
            transition:fade={{ duration: 100 }}
        >
            {#each children as id}
                <svelte:self id={id} {state} {contextMenu} />
            {/each}
        </div>
    {/if}
</div>