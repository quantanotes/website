<script>
    import { fade } from 'svelte/transition'
    import Icon from '@iconify/svelte'
    import Dropdown from '$/components/atoms/dropdown.svelte'

    let { id, state: _state, contextMenu } = $props()
    let expanded = $state(false)
    let children = $derived(_state.map[id]?.children.toSorted(sorter) || [])
    // Counts nested drag enters
    let dragRC = $state(0)
    // Checks if we need to expand due to drag and drop
    let dragExpanded = $derived(dragRC > 0 && !(_state.moving === id || _state.ascendants(id).includes(_state.moving)))

    async function expand() {
        if (!expanded) await _state.children(id)
        expanded = !expanded
    }

    function handleDragStart(event) {
        event.stopPropagation()
        _state.moving = id
        event.dataTransfer.setData('text/plain', id)
    }

    async function handleDragEnter(event) {
        event.preventDefault()
        dragRC++
        if (!expanded && dragExpanded && dragRC == 1) {
            _state.dragover = id
            await _state.children(id)
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
        const moving = _state.moving
        const dragover = _state.dragover
        _state.moving = ''
        _state.dragover = ''
        dragRC = 0
        if (!moving || !dragover) return
        // If a parent is being dropped in to a child
        if (_state.ascendants(id).includes(moving)) return
        // If it's dropping in to the same parent
        if (dragover === _state.map[moving].parent) return
        // Make sure the UI is expanded so we can see where we dropped
        expanded = true
        if (dragover === moving) return
        await _state.move(moving, id)
    }

    function sorter(a, b) {
        return _state.map[a]?.title.localeCompare(_state.map[b]?.title || '') || false
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
            class={`hover:font-bold transition text-left truncate w-full ${(id == _state.current || dragRC > 0) && 'font-bold'}`}
            onclick={() => _state.open(id)}
        >
            {_state.map[id]?.title}
        </button>

        <Dropdown class="h-[24px] icon-btn">
            <Icon icon="mdi:dots-horizontal" />
            {#snippet menu()}
                {@render contextMenu(id, _state.map[id].parent)}
            {/snippet}
        </Dropdown>
    </div>

    {#if expanded || dragExpanded}
        <div
            class="ml-2"
            transition:fade={{ duration: 100 }}
        >
            {#each children as id}
                <svelte:self id={id} state={_state} {contextMenu} />
            {/each}
        </div>
    {/if}
</div>