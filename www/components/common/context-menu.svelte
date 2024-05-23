<script>
    import { computePosition } from '@floating-ui/dom'
    import contextMenu from '$/stores/context-menu.svelte.js'

    let ref = $state()
    let style = $state('')

    function handleClick(event) {
        if (!contextMenu.contextMenu || (ref && ref.contains(event.target))) {
            return
        }
        contextMenu.hide()
    }

    $effect(async () => {
        const cm = contextMenu.contextMenu
        if (cm === undefined) return

        const strategy = { strategy: 'fixed', placement: cm.placement || 'bottom' }
        const styleFn = ({x, y}) => {
            style = cm.style || `position: fixed; left: ${x}px; top: ${y}px; z-index: 100;`
        } 

        // This has to be done twice - perhaps something to do with ref not being populated?
        await computePosition(cm.source, ref, strategy).then(styleFn)
        await computePosition(cm.source, ref, strategy).then(styleFn)
    })
</script>

<div class={contextMenu.contextMenu === undefined ? 'invisible' : ''} bind:this={ref} {style}>
    {#if contextMenu.contextMenu}
        {@render contextMenu.contextMenu.component()}
    {/if}
</div>

<svelte:window onclick={handleClick} />
