<script context="module">
    export { default as layout } from '$/layouts/page.svelte'
</script>

<script>
    import Hero from '$/components/index/hero.svelte'
    import Products from '$/components/index/products.svelte'
    import Mission from '$/components/index/mission.svelte'
</script>

<Hero />
<Products />
<Mission />
