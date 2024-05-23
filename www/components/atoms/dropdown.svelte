<script>
    import contextMenu from '$/stores/context-menu.svelte.js'

    let { class: _class, children, menu, placement = 'bottom', style = undefined } = $props()
    let ref

    function handleClick(event) {
        contextMenu.show(ref, menu, placement, style, event)
    }
</script>

<button
    class={_class}
    bind:this={ref}
    onclick={handleClick}
>
    {@render children()}
</button>
