<script>
    import Icon from '@iconify/svelte'
    import Navbar from '$/components/common/navbar/navbar.svelte'

    let { current } = $props()

    function back() {
        current = []
    }
</script>

<div class="fixed w-screen z-50">
    <Navbar>
        {#snippet start()}
            {#if current.length}
                <button class="icon-btn" onclick={back}>
                    <Icon icon="mdi:arrow-left" />
                </button>
                <div class="h-6 text-contrast">
                    {current[0].title}
                </div>
            {/if}
        {/snippet}
    </Navbar>
</div>