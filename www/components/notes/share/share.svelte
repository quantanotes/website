<script>
    import Links from './links.svelte'
    import Publish from './publish.svelte'

    let { id } = $props()
</script>

<div
    class="bg-primary border border-secondary center-screen flex flex-col gap-4 p-4 shadow-md shadow-secondary w-96"
>
    <Publish {id} />
    <hr class="border-contrast" />
    <Links {id} />
</div>
