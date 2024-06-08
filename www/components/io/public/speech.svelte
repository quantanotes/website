<script>
    import { blur } from 'svelte/transition'
    import { inview } from 'svelte-inview'
    import SectionDivider from '$/components/atoms/section-divider.svelte'

    const sentences = [
        'What you take in, is what you create; we are in I/O with the universe.',
        'We calculate the input, to unlock your maximum output.',
        'Io understands your goals and behaviour to build action and insight.',
        'Using our LARGE ACTION MODEL technology, we connect to your documents and external tools to solve complex tasks with AI.',
    ]

    let inView = $state(sentences.map(() => false))
</script>

{#each sentences as sentence, i}
    <div class="font-bold h-screen p-4 relative text-5xl md:text-7xl">
        <div
            use:inview={{ threshold: 0.5 }}
            oninview_change={(event) => {
                inView[i] = event.detail.inView
            }}
        >
            {#if inView[i]}
                <div transition:blur={{ duration: 800 }}>
                    {sentence}
                </div>
            {/if}
        </div>
    </div>
{/each}

<SectionDivider />
