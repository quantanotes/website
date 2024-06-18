<script>
    let { state, row } = $props()
</script>

<tr>
    {#each state.columns as col}
        <td class="p-2">
            {#if col.type === 'bool'}
                <input
                    type="checkbox"
                    checked={row.data[col.name]}
                    onchange={(event) => state.updateColumn(row.id, col, event.target.checked)}
                />
            {:else if col.type === 'number'}
                <input
                    class="bg-primary focus:outline-none text-secondary w-fit"
                    type="number"
                    value={row.data[col.name]}
                    oninput={(event) =>
                        state.updateColumn(row.id, col, parseFloat(event.target.value))}
                />
            {:else if col.type === 'text'}
                <input
                    class="bg-primary focus:outline-none text-secondary w-fit"
                    type="text"
                    value={row.data[col.name]}
                    oninput={(event) => state.updateColumn(row.id, col, event.target.value)}
                />
            {:else if col.type === 'note'}
                {row.data[col.name] || '\xa0'}
            {/if}
        </td>
    {/each}
</tr>
