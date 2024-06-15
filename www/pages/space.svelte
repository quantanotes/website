<script context="module">
    export { default as layout } from '$/layouts/base.svelte'
</script>

<script>
    import { onMount } from 'svelte'
    import Space from '$/components/space/space.svelte'
    import Thread from '$/components/space/thread.svelte'
    import Input from '$/components/space/input.svelte'
    import Navbar from '$/components/space/navbar.svelte'

    let data = $state({ nodes: [], links: [] })
    let current = $state([])

    async function children(id) {
        const response = await fetch('/api/space/children', {
            method: 'POST',
            body: JSON.stringify({
                parent: id,
            }),
        })
        const nodes = (await response.json()) || []
        await addNodes(nodes)
    }

    async function roots() {
        const response = await fetch('/api/space/roots', {
            method: 'POST',
        })
        const nodes = (await response.json()) || []
        await addNodes(nodes)
        getEdges()
    }

    async function addNodes(nodes) {
        nodes = nodes.filter((node) => data.nodes.find((n) => n.id === node.id) === undefined)
        nodes.forEach((node) => data.nodes.push(node))
        await Promise.all(nodes.map((node) => children(node.id)))
    }

    function getEdges() {
        data.links = []
        data.nodes.forEach((node) => {
            if (data.nodes.find((n) => n.id === node.parent)) {
                data.links.push({ source: node.parent, target: node.id })
            }
        })
    }

    onMount(() => {
        roots()
    })

    function onNodeClick(node) {
        current = [node]
    }
</script>

<Space bind:data {onNodeClick} />
<Navbar bind:current />
<div class="absolute pt-11 w-screen">
    <div class="h-fit max-w-3xl mx-auto p-4">
        <Input />
    </div>
</div>
{#if current.length}
    <Thread bind:nodes={current} />
{/if}
