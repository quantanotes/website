<script>
    import Item from './item.svelte'

    let { show, width = 288, state, contextMenu } = $props()
</script>

<div
    class={`border-contrast border-r fixed h-full pb-64 no-scrollbar overflow-y-auto p-2 transition-all ${!show && '-translate-x-full'}`}
    style={`width: ${width}px`}
>
    {#if state.root}
        <Item bind:id={state.root} {state} {contextMenu} />
    {/if}
</div>
