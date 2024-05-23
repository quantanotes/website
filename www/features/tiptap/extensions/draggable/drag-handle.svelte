<script>
    import Icon from '@iconify/svelte'
</script>

<div class="flex flex-row justify-between text-2xl w-10">
    <button class="cursor-grab icon-btn">
        <Icon icon="mdi-drag" />
    </button>
    <button class="icon-btn">
        <Icon icon="mdi-add" />
    </button>
</div>

<style lang="postcss">
    :global(.drag-handle) {
        @apply fixed transition-opacity z-50;
    }
</style>
