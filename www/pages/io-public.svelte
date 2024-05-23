<script context="module">
    export { default as layout } from '$/layouts/page.svelte'
</script>

<script>
    import Cta from '$/components/io/public/cta.svelte'
    import Faq from '$/components/io/public/faq.svelte'
    import Goals from '$/components/io/public/goals.svelte'
    import Hero from '$/components/io/public/hero.svelte'
    import Pricing from '$/components/io/public/pricing.svelte'
    import Speech from '$/components/io/public/speech.svelte'
    import UseCases from '$/components/io/public/use-cases.svelte'
</script>

<Hero />
<Goals />
<UseCases />
<Speech />
<Pricing />
<Faq />
<Cta />
