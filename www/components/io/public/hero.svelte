<script>
    import SectionDivider from '$/components/atoms/section-divider.svelte'
    import Hero from '$/components/common/hero.svelte'
    import CtaAtom from './cta-atom.svelte'

    const headline = 'io saids hello'
    const tagline =
        'a personalised ai agent for your mind, thoughts and feelings, connected to the outside world to complete tasks'
</script>

<Hero {headline} {tagline} />
<CtaAtom />
<SectionDivider />
