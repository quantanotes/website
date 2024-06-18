<script>
    import Icon from '@iconify/svelte'
    import HeaderCell from './header-cell.svelte'

    let { state } = $props()
</script>

<thead>
    <tr>
        {#each state.columns as col, pos}
            <HeaderCell {state} {col} {pos} />
        {/each}
        <th class="font-normal p-2">
            <button class="flex flex-row gap-2 items-center icon-btn" onclick={state.showAddColumnMenu}>
                <div class="h-fit">
                    <Icon icon="mdi-plus" />
                </div>
                <div class="leading-normal">Add column</div>
            </button>
        </th>
    </tr>
</thead>
