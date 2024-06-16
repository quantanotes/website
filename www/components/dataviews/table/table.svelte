<script>
    import Icon from '@iconify/svelte'
    import contextMenu from '$/stores/context-menu.svelte.js'
    import AddColumnMenu from './add-column-menu.svelte'
    import { onMount } from 'svelte'

    let { state } = $props()

    let root = $derived(state.current)
    let columns = $derived(state.map[state.current].metadata?.columns || [])
    let data = $derived(
        state.map[state.current].children
            .map((id) => state.map[id])
            .filter((node) => node.category === 'record')
            .map((node) => ({ id: node.id, data: deserialiseRow(node.id) })),
    )
    let tableHeight = $state()
    let columnRefs = $state([null, null, null])
    let columnWidths = $state([0, 0, 0])
    let resizingColumn = $state()

    onMount(() => {
        state.children(root)
    })

    function handleColumnResize(index) {
        resizingColumn = index

        function handleMouseMove(event) {
            columnWidths[resizingColumn] = event.clientX - columnRefs[resizingColumn].offsetLeft
        }

        function handleMouseUp() {
            window.removeEventListener('mousemove', handleMouseMove)
            window.removeEventListener('mouseup', handleMouseUp)
            resizingColumn = undefined
        }

        window.addEventListener('mousemove', handleMouseMove)
        window.addEventListener('mouseup', handleMouseUp)
    }

    // TODO: add server side validation

    function showAddColumnMenu(event) {
        contextMenu.show(event.target, addColumnMenu, 'bottom', '', event)
    }

    function addColumn(name, type) {
        if (!['bool', 'number', 'text', 'note'].includes(type)) {
            throw new Error('Column is invalid type')
        }
        let metadata = state.map[root].metadata || {}
        let columns = metadata.columns || []
        if (columns.find((column) => column.name === name)) {
            throw new Error('Column already exists')
        }
        state.map[root].metadata = { ...metadata, columns: [...columns, { name, type }] }
        state.update(root)
    }

    function addRow() {
        state.create(root, 'record')
    }

    function updateColumn(id, column, value) {
        let row = deserialiseRow(id)
        switch (column.type) {
            case 'bool':
                if (typeof value !== 'boolean')
                    throw new Error('Value is not correct type for column')
                break
            case 'number':
                if (typeof value !== 'number')
                    throw new Error('Value is not correct type for column')
            case 'text':
                if (typeof value !== 'string')
                    throw new Error('Value is not correct type for column')
            case 'note':
                if (typeof value !== 'string')
                    throw new Error('Value is not correct type for column')
        }
        row[column.name] = value
        serialiseRow(id, row)
        state.update(id)
    }

    // Using JSON serialisation in to content for row
    // Not the best idea but fits nicely in to the data model
    function deserialiseRow(id) {
        let node = state.map[id]
        try {
            let data = JSON.parse(node.content)
            return data
        } catch {
            return {}
        }
    }

    function serialiseRow(id, data) {
        state.map[id].content = JSON.stringify(data)
    }
</script>

{#snippet addColumnMenu()}
    <AddColumnMenu {addColumn} />
{/snippet}

<table class="w-fit" bind:clientHeight={tableHeight}>
    <thead>
        <tr>
            {#each columns as col, i}
                <th
                    class="border-secondary border-b font-normal py-1 px-2 relative text-left"
                    style={`width: ${columnWidths[i]}px`}
                    bind:clientWidth={columnWidths[i]}
                    bind:this={columnRefs[i]}
                >
                    <span>
                        {col.name}
                    </span>
                    <!-- svelte-ignore a11y_no_static_element_interactions -->
                    <div
                        class="absolute block cursor-col-resize top-0 right-0 w-2 z-10"
                        onmousedown={() => handleColumnResize(i)}
                        style={`height: ${tableHeight}px`}
                    ></div>
                </th>
            {/each}
            <th class="p-1 w-4">
                <button class="icon-btn" onclick={showAddColumnMenu}>
                    <Icon icon="mdi-plus" />
                </button>
            </th>
        </tr>
    </thead>
    <tbody>
        {#each data as row, i}
            <tr>
                {#each columns as col}
                    <td class={`py-1 px-2 ${i % 2 !== 0 && 'bg-contrast-dark'}`}>
                        {#if col.type === 'bool'}
                            <input
                                type="checkbox"
                                checked={row.data[col.name]}
                                onchange={(event) => updateColumn(row.id, col, event.target.checked)}
                            />
                        {:else if col.type === 'number'}
                            <input
                                type="number"
                                value={row.data[col.name]}
                                oninput={(event) =>
                                    updateColumn(row.id, col, parseFloat(event.target.value))}
                            />
                        {:else if col.type === 'text'}
                            <input
                                type="text"
                                value={row.data[col.name]}
                                oninput={(event) => updateColumn(row.id, col, event.target.value)}
                            />
                        {:else if col.type === 'note'}
                            {row.data[col.name] || '\xa0'}
                        {/if}
                    </td>
                {/each}
            </tr>
        {/each}
        <tr>
            <td class="p-1 w-4">
                <button class="icon-btn" onclick={addRow}>
                    <Icon icon="mdi-plus" />
                </button>
            </td>
        </tr>
    </tbody>
</table>
