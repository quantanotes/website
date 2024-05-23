<script lang="ts">
    import { blur } from 'svelte/transition'
    import Icon from '@iconify/svelte'
    import { z } from 'zod'
    import Input from '$/components/atoms/input.svelte'

    export let inputs: {
        name: string
        label: string
        type: string
        schema?: z.Schema
        shake?: boolean
    }[]
    export let data: { [key: string]: string } = {}
    export let submitLabel: string
    export let action: (data: { [key: string]: string }) => Promise<string>

    let message: string
    let loading = false

    async function handleSubmit(event: Event) {
        event.preventDefault()
        if (!check()) return

        message = ''
        loading = true

        try {
            message = await action(data)
        } finally {
            loading = false
        }
    }

    function check() {
        let valid = true
        for (let i = 0; i < inputs.length; i++) {
            const input = inputs[i]
            if (
                input.schema &&
                !input.schema.safeParse(data[input.name]).success
            ) {
                valid = false
                inputs[i] = { ...input, shake: true }
            }
        }
        return valid
    }
</script>

{#key loading}
    <div transition:blur>
        {#if loading}
            <Icon class="animate-spin mx-auto size-5 " icon="mdi:loading" />
            <br />
            <br />
        {:else}
            <form
                class="flex flex-col gap-4 text-center w-96"
                on:submit={handleSubmit}
            >
                {#each inputs as input}
                    <Input
                        name={input.name}
                        label={input.label}
                        type={input.type}
                        schema={input.schema}
                        bind:shake={input.shake}
                        bind:value={data[input.name]}
                    />
                {/each}
                <button class="btn" type="submit">{submitLabel}</button>
            </form>
            <br />
            {#if message}
                <p>{message}</p>
                <br />
            {/if}
        {/if}
    </div>
{/key}
