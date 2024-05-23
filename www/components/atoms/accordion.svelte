<script>
    import { fade } from 'svelte/transition'
    import Icon from '@iconify/svelte'

    /**@type { items: { title: string, content: string }[] } */
    let { items } = $props()
    let active = $state(null)

    function toggle(index) {
        active = active === index ? null : index
    }
</script>

<div class="accordion">
    {#each items as item, index}
        <div>
            <button
                class="w-full text-left p-4 flex justify-between items-center"
                onclick={() => toggle(index)}
            >
                <div>
                    {item.title}
                </div>
                <Icon icon={active === index ? 'mdi-minus' : 'mdi-plus'} />
            </button>
            {#if active === index}
                <div class="p-4 transition-transform" in:fade={{ duration: 300 }}>
                    {item.content}
                </div>
            {/if}
        </div>
    {/each}
</div>
