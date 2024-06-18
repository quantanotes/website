<script>
    import Icon from '@iconify/svelte'
    import notes from '$/stores/notes.svelte.js'
    import Navbar from '../common/navbar/navbar.svelte'
    import Input from './input.svelte'

    let { toggleSidebar, inputValue, inputSend } = $props()
    let current = $derived(notes.map[notes.current])
</script>

<div class="fixed w-full z-50">
    <Navbar>
        {#snippet start()}
            <button class="icon-btn" onclick={toggleSidebar}>
                <Icon icon="mdi:menu" />
            </button>
            <div class="h-6 text-contrast truncate">
                {current && current.title}
            </div>
        {/snippet}

        {#snippet middle()}
            <div class="flex min-w-0 w-96">
                <Input bind:value={inputValue} bind:send={inputSend} />
            </div>
        {/snippet}
    </Navbar>
</div>
