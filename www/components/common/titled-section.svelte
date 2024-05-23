<script>
    import SectionDivider from '../atoms/section-divider.svelte'

    let { title, content } = $props()
</script>

<div class="p-4">
    <h1 class="flex font-bold justify-between mx-auto text-4xl text-center w-fit">{title}</h1>

    <br />
    <br />
    <br />
    <br />

    {@render content()}
</div>

<SectionDivider />
