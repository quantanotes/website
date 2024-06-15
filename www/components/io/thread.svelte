<script>
    import { fade } from 'svelte/transition'
    import { marked } from 'marked'

    export let thread
    export let container
    export let disableTransitions = false

    let animation = disableTransitions ? fade : undefined

    function parse(content, citations) {
        if (!content) return ''

        if (!citations || citations.length === 0) return marked.parse(content)

        const cited = content.replace(
            /\[@(\d+)\]/g,
            (_, n) => `<a href="${citations[n - 1]}">[${n}]</a>`,
        )

        return marked.parse(cited)
    }
</script>

<div class="flex flex-col no-scrollbar overflow-y-auto pb-64 w-full" bind:this={container}>
    {#each thread as msg}
        <div
            class="border-b first:border-t border-x border-contrast flex flex-col gap-2 p-4 text-left w-full"
            transition:fade={disableTransitions && { duration: 0 }}
        >
            <p class="font-bold">
                {msg.role === 'assistant' ? 'Io' : 'You'}
            </p>

            <p class="max-w-none prose dark:prose-invert prose-neutral">
                {@html parse(msg.content, msg.citations)}
            </p>

            {#if msg.citations}
                <div class="max-w-none prose prose-neutral dark:prose-invert">
                    <ol class="mx-8">
                        {#each msg.citations as citation}
                            <li>
                                <a class="text-ellipsis" href={citation}>
                                    {citation}
                                </a>
                            </li>
                        {/each}
                    </ol>
                </div>
            {/if}
        </div>
    {/each}
</div>
