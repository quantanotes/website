<script>
    import notes from '$/stores/notes.svelte.js'
    import Toggle from '$/components/atoms/toggle.svelte'

    let { id } = $props()
    let published = $state(notes.map[id].public > 0)

    async function onToggle(active) {
        if (active) {
            await notes.publish(id)
        } else if (notes.map[id].public > 0) {
            await notes.unpublish(id)
        }

        published = notes.map[id].public > 0
    }
</script>

<div class="flex justify-between">
    <div>Publish to Space</div>
    <Toggle bind:active={published} {onToggle} />
</div>
