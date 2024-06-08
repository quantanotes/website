<script>
    import { onMount } from 'svelte'
    import { scale } from 'svelte/transition'

    export let value = ''
    export let placeholder = ''
    export let label = ''
    export let onkeydown

    let textarea
    let rows = 1

    onMount(() => {
        textarea.style.height = '0px'
        textarea.style.height = textarea.scrollHeight + 'px'
    })

    $: {
        rows = (value.match(/\n/g) || []).length + 1 || 1
        if (textarea) {
            textarea.style.height = '0px'
            textarea.style.height = textarea.scrollHeight + 'px'
        }
    }
</script>

{#if label}
    {#if value.length > 0}
        <label
            class="absolute bg-primary px-1 text-sm transition-colors translate-x-2 -translate-y-2"
            transition:scale={{ duration: 200 }}
            for=""
        >
            {label}
        </label>
    {/if}
{/if}
<textarea
    class={$$props.class || ''}
    {placeholder}
    {rows}
    bind:this={textarea}
    bind:value
    on:keydown={onkeydown}
></textarea>
