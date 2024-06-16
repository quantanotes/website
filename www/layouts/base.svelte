<script>
    import { onMount } from 'svelte'
    import auth from '$/stores/auth.svelte.js'
    import { dialog, contextMenu } from '$/stores/portals.svelte'
    import Portal from '$/components/common/portal.svelte'

    let { children } = $props()

    onMount(() => {
        auth.refreshDetails()
    })
</script>

<svelte:head>
    <title>Quanta</title>
</svelte:head>

<Portal portal={contextMenu} />
<Portal portal={dialog} />

<div class="absolute inset-0 h-screen overflow-x-hidden w-screen">
    {@render children()}
</div>
