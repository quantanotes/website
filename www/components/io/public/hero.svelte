<script>
    import SectionDivider from '$/components/atoms/section-divider.svelte'
    import Hero from '$/components/common/hero.svelte'
    import CtaAtom from './cta-atom.svelte'

    const headline = '10X yourself with Io'
    const tagline =
        'Io is an AI that completes tasks to make progress towards your goals.'
</script>

<Hero {headline} {tagline} />
<CtaAtom />
<SectionDivider />
