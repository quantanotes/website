<script>
    import Toggle from '$/components/atoms/toggle.svelte'
    import Icon from '@iconify/svelte'
    import { onMount } from 'svelte'

    let { id } = $props()
    let links = $state()

    onMount(async () => {
        await getLinks()
    })

    async function getLinks() {
        const res = await fetch('/api/permissions/links/retrieve', {
            method: 'POST',
            body: JSON.stringify({ id }),
        })
        links = await res.json()
    }

    async function newLink() {
        await fetch('/api/permissions/links/create', {
            method: 'POST',
            body: JSON.stringify({ id }),
        })
        await getLinks()
    }

    async function deleteLink(id) {
        const res = await fetch('/api/permissions/links/delete', {
            method: 'POST',
            body: JSON.stringify({ id }),
        })
        if (res.ok) links = links.filter((link) => link.id !== id)
    }
</script>

<div class="flex justify-between">
    <div>Create link</div>
    <button class="icon-btn" onclick={newLink}>
        <Icon icon="mdi:add" />
    </button>
</div>
<div class="flex flex-col gap-4">
    {#each links as link}
        <div class="border border-contrast flex flex-col gap-4 p-4">
            <div class="flex justify-between">
                <a class="truncate" href="https://quanta.uno/notes/{link.id}"
                    >quanta.uno/notes/{link.id}</a
                >
                <button class="icon-btn" onclick={() => deleteLink(link.id)}>
                    <Icon icon="mdi:delete" />
                </button>
            </div>
            <div class="flex justify-between">
                <div>Allow modifications</div>
                <Toggle />
            </div>
        </div>
    {/each}
</div>
