<script>
    import SectionDivider from '$/components/atoms/section-divider.svelte'

    const goals = [
        {
            goal: 'problem solving',
            description:
                'io is capable of connecting to the internet, your notes and your databases to perform complex tasks and automate your workflows',
        },
        {
            goal: 'human solving',
            description:
                'we designed io to be as human oriented as possible. io provides you the emotional and cognitive tools to help you overcome your problems whether personal or practical',
        },
    ]
</script>

<div class="p-4">
    <h1 class="flex font-bold justify-between mx-auto text-4xl text-center w-fit">goals</h1>

    <br />
    <br />
    <br />
    <br />

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
</div>

<SectionDivider />
