<script>
    import Icon from '@iconify/svelte'
    import Hero from '$/components/common/hero.svelte'
    import SectionDivider from '$/components/atoms/section-divider.svelte'

    const headline = 'machine intelligence for human intelligence'
    const tagline = 'we build AI to enhance productivity and mindfulness.'
    const suggestions = [
        'can you keep me accountable for me meditation habit?',
        'build an automation for my catering business',
        'scale and automate my email outreach campaign',
    ]
</script>

<Hero {headline} {tagline} />

<ul class="flex flex-col gap-4 max-w-4xl mx-auto px-4 text-left w-full">
    {#each suggestions as suggestion}
        <li>
            <a class="flex hover:font-bold justify-between transition-all" href="/">
                {suggestion}
                <Icon icon="mdi:arrow-top-right" />
            </a>
        </li>
    {/each}
</ul>

<SectionDivider />
