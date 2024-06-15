<script>
    import Icon from '@iconify/svelte'

    const columns = ['name', 'age', 'gender']

    const data = [
        { name: 'Bob', age: 32, gender: 'male' },
        { name: 'Alice', age: 26, gender: 'female' },
        { name: 'Charlie', age: 18, gender: 'male' },
    ]

    let tableHeight = $state()
    let columnRefs = $state([null, null, null])
    let columnWidths = $state([0, 0, 0])
    let resizingColumn = $state()

    let hoverColumn = $state()
    let hoverRow = $state()

    function handleColumnResize(index) {
        resizingColumn = index

        function handleMouseMove(event) {
            columnWidths[resizingColumn] = event.clientX - columnRefs[resizingColumn].offsetLeft
            event.stopPropogation()
        }

        function handleMouseUp() {
            window.removeEventListener('mousemove', handleMouseMove)
            window.removeEventListener('mouseup', handleMouseUp)
            resizingColumn = undefined
        }

        window.addEventListener('mousemove', handleMouseMove)
        window.addEventListener('mouseup', handleMouseUp)
    }
</script>

<table class="w-fit" bind:clientHeight={tableHeight}>
    <thead>
        <tr class="border-secondary border-b">
            {#each columns as col, i}
                <th
                    class="font-normal py-1 px-2 relative text-left"
                    style={`width: ${columnWidths[i]}px`}
                    bind:clientWidth={columnWidths[i]}
                    bind:this={columnRefs[i]}
                >
                    <span>
                        {col}
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
                <button class="icon-btn">
                    <Icon icon="mdi-plus" />
                </button>
            </th>
        </tr>
    </thead>
    <tbody>
        {#each data as row, i}
            <tr>
                {#each columns as col}
                    <td class="py-1 px-2">
                        {row[col]}
                    </td>
                {/each}
            </tr>
        {/each}
        <tr>
            <td class="p-1 w-4">
                <button class="icon-btn">
                    <Icon icon="mdi-plus" />
                </button>
            </td>
        </tr>
    </tbody>
</table>
