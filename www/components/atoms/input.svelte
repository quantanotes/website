<script>
    import { scale } from 'svelte/transition'

    export let label = ''
    export let value = ''
    export let name = ''
    export let type = 'text'
    export let schema
    export let shake = false

    let active = false
    let input
    let issues = []

    $: if (input) input.type = type
    $: if (schema != undefined) {
        const result = schema.safeParse(value)
        if (!result.success) {
            issues = result.error.issues
        } else {
            issues = []
        }
    }
    $: if (shake) setTimeout(() => (shake = false), 500)
</script>

<div class={`flex flex-col relative ${shake && 'animate-shake'}`}>
    {#if value.length > 0}
        <label
            class={`absolute bg-primary px-1 text-sm transition-colors translate-x-2 -translate-y-2 
                    ${shake ? 'text-red-500' : !active ? 'text-contrast' : issues.length ? 'text-red-500' : ''}`}
            for={name}
            transition:scale={{ duration: 200 }}
        >
            {label}
        </label>
    {/if}
    <input
        bind:this={input}
        class={`
            input 
            ${issues.length && 'focus:border-red-500'}
            ${shake && 'border-red-500'}`}
        placeholder={label}
        {name}
        bind:value
        on:focus={() => (active = true)}
        on:blur={() => (active = false)}
    />
</div>
{#if active && issues.length}
    <ul class="flex flex-col flex-wrap gap-1 text-left text-sm">
        {#each issues as issue}
            <li>{issue.message}</li>
        {/each}
    </ul>
{/if}
