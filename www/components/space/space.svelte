<script>
    import { onMount } from 'svelte'
    import ForceGraph from 'force-graph'
    import darkMode from '$/stores/dark-mode.svelte.js'

    /** @type {{ data: any, onNodeClick: (node: NodeObject, event: MouseEvent) => void }} */
    let { data, onNodeClick } = $props()

    let ref
    let width = $state(0)
    let height = $state(0)

    const graph = ForceGraph()

    onMount(async () => {
        // if (typeof window === 'undefined') return

        // let ForceGraph = (await import('force-graph')).default
        // graph = ForceGraph()

        graph(ref)
            .graphData(data)
            .nodeCanvasObject((node, ctx, globalScale) => {
                const label = node.title
                const fontSize = 18 / globalScale
                ctx.font = `${fontSize}px Roboto`
                ctx.textAlign = 'center'
                ctx.textBaseline = 'middle'
                ctx.fillStyle = darkMode.darkMode ? 'hsl(0 0% 90%)' : 'hsl(0 0% 10%)'
                ctx.fillText(label, node.x, node.y)
            })
            .linkCanvasObject((link, ctx, globalScale) => {
                ctx.lineWidth = 1 / globalScale
                ctx.strokeStyle = darkMode.darkMode ? 'hsl(0 0% 24%)' : 'hsl(0 0% 76%)'
                ctx.beginPath()
                ctx.moveTo(link.source.x, link.source.y)
                ctx.lineTo(link.target.x, link.target.y)
                ctx.stroke()
            })
            .onNodeClick(onNodeClick)
            .minZoom(Number.NEGATIVE_INFINITY)
            .maxZoom(Number.POSITIVE_INFINITY)
            .nodeRelSize(4)
    })

    $effect(() => {
        graph.graphData(data)
    })

    $effect(() => {
        graph.width(width)
        graph.height(height)
    })
</script>

<div
    class="absolute h-screen w-screen"
    bind:this={ref}
    bind:clientWidth={width}
    bind:clientHeight={height}
></div>
