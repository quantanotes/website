<script>
    import Footer from '$/components/common/footer.svelte'
    import Navbar from '$/components/index/navbar.svelte'
    import Base from '$/layouts/base.svelte'

    let { children } = $props()
</script>

<Base>
    <Navbar />
    {@render children()}
    <Footer />
</Base>