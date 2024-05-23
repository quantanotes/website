<script>
    import { fade } from 'svelte/transition'
    import Note from './note.svelte'

    let { nodes } = $props()
</script>

<div class="absolute backdrop-blur-lg h-screen overflow-y-auto w-screen" transition:fade>
    <div class="max-w-7xl mb-64 mt-16 mx-auto p-4 w-full">
        {#each nodes as node}
            <Note node={node} />
        {/each}
    </div>
</div>
