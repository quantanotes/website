<script>
    import { fade } from 'svelte/transition'

    let { children } = $props()
</script>

<div
    class="backdrop-blur fixed h-screen left-0 pointer-events-none overflow-clip top-0 w-screen z-10"
    transition:fade={{ duration: 100 }}
></div>
<div
    class="bg-primary border border-contrast center-screen flex flex-col gap-4 p-4 relative w-96 z-20"
>
    {@render children()}
</div>
