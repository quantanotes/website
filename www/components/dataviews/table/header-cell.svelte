<script>
    import Icon from '@iconify/svelte'
    import DropdownMenu from '$/components/atoms/dropdown-menu.svelte'
    import ContextMenu from '$/components/atoms/context-menu.svelte'

    let { state, pos, col } = $props()
</script>

<!-- svelte-ignore a11y_mouse_events_have_key_events -->
<th
    class="py-1 px-2 relative text-left"
    style={`width: ${state.columnWidths[pos]}px`}
    bind:clientWidth={state.columnWidths[pos]}
    bind:this={state.columnRefs[pos]}
>
    <div class="flex">
        <span class="flex-grow truncate">
            {col.name}
        </span>
        <DropdownMenu class="h-[24px] icon-btn">
            <Icon icon="mdi-dots-horizontal" />
            {#snippet menu()}
                <ContextMenu
                    actions={[
                        { name: 'Delete', action: () => state.removeColumn(pos) },
                        { name: 'Rename', action: () => {} },
                    ]}
                />
            {/snippet}
        </DropdownMenu>

        <!-- svelte-ignore a11y_no_static_element_interactions -->
        <div
            class="absolute block cursor-col-resize top-0 right-0 w-2 z-10"
            onmousedown={() => state.handleColumnResize(pos)}
            style={`height: ${state.tableHeight}px`}
        ></div>
    </div>
</th>
