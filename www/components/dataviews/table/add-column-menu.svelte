<script>
    import { dialog } from '$/stores/portals.svelte.js'
    import Dialog from '$/components/atoms/dialog.svelte'
    import Input from '$/components/atoms/input.svelte'

    let { state } = $props()

    let name = $state()
    let type = $state()

    function handleOnClick() {
        state.addColumn(name, type)
        dialog.hide()
    }
</script>

<Dialog>
    <Input label="Name" bind:value={name} />
    <select class="input" bind:value={type}>
        <option value="bool">True/False</option>
        <option value="number">Number</option>
        <option value="text">Text</option>
        <option value="note">Note</option>
    </select>
    <button class="btn w-full" onclick={handleOnClick}>Add Column</button>
</Dialog>
