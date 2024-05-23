<script>
    import SectionDivider from '$/components/atoms/section-divider.svelte'

    const usecases = [
        'life coach',
        'software engineer',
        'personal trainer',
        'medical advisor',
        'sales automation',
        'public relations specialist',
        'researcher',
        'virtual tutor',
        'market analyst',
        'customer service representative',
        'event planner',
        'roleplaying',
        'legal advisor',
        'personal assistant',
        'internet search engine',
        'web scraping',
        'cold emailer',
        'content planning',
        'accountability partner',
        'responsive journal',
        'meditation guru',
        'story telling',
        'copywriting',
    ]
</script>

<div class="p-4">
    <h1 class="flex font-bold justify-between mx-auto text-4xl text-center w-fit">use cases</h1>

    <br />
    <br />
    <br />
    <br />

    <div class="flex flex-col flex-wrap h-96 max-w-4xl mx-auto">
        {#each usecases as usecase}
            <div>
                {usecase}
            </div>
        {/each}
    </div>
</div>

<SectionDivider />
