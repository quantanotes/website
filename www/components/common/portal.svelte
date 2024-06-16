<script>
    import { computePosition } from '@floating-ui/dom'

    let { portal } = $props()
    
    let ref = $state()
    let style = $state('')

    function handleClick(event) {
        if (!portal.state || (ref && ref.contains(event.target))) {
            return
        }
        portal.hide()
    }

    $effect(async () => {
        const p = portal.state
        if (p === undefined) return

        const strategy = { strategy: 'fixed', placement: p.placement || 'bottom' }
        const styleFn = ({ x, y }) => {
            style = p.style || `position: fixed; left: ${x}px; top: ${y}px; z-index: 100;`
        }

        // This has to be done twice - perhaps something to do with ref not being populated?
        await computePosition(p.source, ref, strategy).then(styleFn)
        await computePosition(p.source, ref, strategy).then(styleFn)
    })
</script>

<div class={portal.contextMenu === undefined ? 'invisible' : ''} bind:this={ref} {style}>
    {#if portal.state}
        {@render portal.state.component()}
    {/if}
</div>

<svelte:window onclick={handleClick} />
