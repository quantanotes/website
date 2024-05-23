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
    <div transition:blur>
        {@render children()}
    </div>
{/key}
