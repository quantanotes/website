<script>
    import { fade } from 'svelte/transition'
    import Icon from '@iconify/svelte'
    import Textarea from '$/components/atoms/textarea.svelte'

    export let value
    export let send

    let focus = false

    function handleKeydown(event) {
        if (event.key === 'Enter' && !event.shiftKey) {
            event.preventDefault()
            send()
        }
    }
</script>

<div
    class="bg-primary border border-contrast focus-within:border-secondary px-8 py-4 focus-within:shadow-md transition w-full"
    on:focusin={() => (focus = true)}
    on:focusout={() => (focus = false)}
>
    <div class="flex flex-row gap-2">
        <Textarea
            class="bg-primary focus:outline-none resize-none text-secondary placeholder:text-contrast-light w-full"
            placeholder="Search anything..."
            onkeydown={handleKeydown}
            bind:value
        />
        <button
            class="h-[24px] icon-btn"
            on:click={send}
        >
            <Icon icon="mdi:search" />
        </button>
    </div>

    {#if focus}
        <div transition:fade={{ duration: 100 }}></div>
    {/if}
</div>
