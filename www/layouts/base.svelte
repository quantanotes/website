<script>
    import { onMount } from 'svelte'
    import { blur } from 'svelte/transition'
    import { page } from '@inertiajs/svelte'
    import auth from '$/stores/auth.svelte.js'
    import ContextMenu from '$/components/common/context-menu.svelte'

    let { children } = $props()

    onMount(async () => {
        await auth.refreshDetails()
    })
</script>

<svelte:head>
    <title>Quanta</title>
</svelte:head>

<ContextMenu />

{#key $page.url}
    <div class="absolute inset-0 h-screen overflow-x-hidden w-screen" transition:blur={{ duration: 300 }}>
        {@render children()}
    </div>
{/key}
