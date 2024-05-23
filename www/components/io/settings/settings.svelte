<script>
    import Textarea from '$/components/atoms/textarea.svelte'
    import Resources from './resources.svelte'

    let { id } = $props()
</script>

<div
    class="bg-primary border border-secondary center-screen flex flex-col gap-4 p-4 shadow-md shadow-secondary w-96"
>
    <Textarea
        class="input resize-none no-scrollbar"
        label="Instruction"
        placeholder="Instruction"
    />
    <Resources />
    <button class="btn">Save</button>
</div>
