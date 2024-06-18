<script>
    import createTableStore from './store.svelte'
    import AddColumnMenu from './add-column-menu.svelte'
    import Body from './body.svelte'
    import Header from './header.svelte'

    // The crud controller state
    let { state: _state } = $props()

    const state = createTableStore(_state, addColumnMenu)
</script>

{#snippet addColumnMenu()}
    <AddColumnMenu {state} />
{/snippet}

<table class="w-fit" bind:clientHeight={state.tableHeight}>
    <Header {state} />
    <Body {state} />
</table>
