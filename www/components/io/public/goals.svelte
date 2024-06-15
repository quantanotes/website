<script>
    import TitledSection from '$/components/common/titled-section.svelte'

    const goals = [
        {
            goal: 'Problem solving',
            description:
                'Io connects to the internet, your documents and external tools to complete tasks and automate your workflow.',
        },
        {
            goal: 'Human solving',
            description:
                "We designed Io with humanity at it's core. Io is like a journal or that helps you process your thoughts and feelings.",
        },
    ]
</script>

<TitledSection title={'Goals'}>
    <div class="flex flex-row flex-wrap justify-evenly gap-8 max-w-4xl mx-auto">
        {#each goals as goal}
            <div class="border border-secondary flex flex-col justify-between h-96 p-4 w-80">
                <div>
                    {goal.goal}
                </div>
                <div>
                    {goal.description}
                </div>
            </div>
        {/each}
    </div>
</TitledSection>
