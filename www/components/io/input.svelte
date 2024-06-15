<script>
    import { fade } from 'svelte/transition'
    import Icon from '@iconify/svelte'
    import Textarea from '$/components/atoms/textarea.svelte'

    let { value, send, placeholder = 'Ask anything...', suggestions = [] } = $props()
    let focus = $state()

    function handleKeydown(event) {
        if (event.key === 'Enter' && !event.shiftKey) {
            event.preventDefault()
            send()
        }
    }
</script>

<div
    class="bg-primary border border-contrast focus-within:border-secondary px-8 py-4 focus-within:shadow-md transition w-full"
    onfocusin={() => (focus = true)}
    onfocusout={() => (focus = false)}
>
    <div class="flex flex-row gap-2">
        <Textarea
            class="bg-primary focus:outline-none resize-none text-secondary placeholder:text-contrast-light w-full"
            onkeydown={handleKeydown}
            {placeholder}
            bind:value
        />
        <button class="h-[24px] icon-btn" onclick={send}>
            <Icon icon="mdi:send" />
        </button>
    </div>

    {#if focus && suggestions.length}
        <br />
        <ul class="flex flex-col gap-4 text-left w-full" transition:fade={{ duration: 100 }}>
            {#each suggestions as suggestion}
                <li>
                    <button
                        class="flex hover:font-bold justify-between transition-all w-full"
                        onclick={() => (value = suggestion)}
                    >
                        <div class="flex-grow text-left">
                            {suggestion}
                        </div>
                        <Icon icon="mdi:arrow-top-right" />
                    </button>
                </li>
            {/each}
        </ul>
    {/if}
</div>
