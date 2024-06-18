<script>
    import Icon from '@iconify/svelte'
    import Cell from './cell.svelte'

    let { state } = $props()
</script>

<tbody>
    {#each state.data as row}
        <Cell {state} {row} />
    {/each}
    <tr>
        <td class="p-2">
            <button class="flex flex-row gap-2 items-center icon-btn" onclick={state.addRow}>
                <div class="h-fit">
                    <Icon icon="mdi-plus" />
                </div>
                <div class="leading-normal">Add row</div>
            </button>
        </td>
    </tr>
</tbody>
