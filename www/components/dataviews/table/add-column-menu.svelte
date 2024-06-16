<script>
    import Input from '$/components/atoms/input.svelte'
    import contextMenu from '$/stores/context-menu.svelte.js'

    let { addColumn } = $props()

    let name = $state()
    let type = $state()

    function handleOnClick() {
        addColumn(name, type)
        contextMenu.hide()
    }
</script>

<div
    class="bg-primary border border-secondary center-screen flex flex-col gap-4 p-4 shadow-md shadow-secondary w-96"
>
    <Input label="Name" bind:value={name} />
    <select class="input" bind:value={type}>
        <option value="bool">True/False</option>
        <option value="number">Number</option>
        <option value="text">Text</option>
        <option value="note">Note</option>
    </select>
    <button class="btn w-full" onclick={handleOnClick}>Add Column</button>
</div>
