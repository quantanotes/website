<script>
    let { headline, tagline } = $props()
</script>

<div class="flex h-[calc(100vh-144px)] items-center justify-center px-4 w-full">
    <div class="max-w-4xl mx-auto text-center">
        <h1 class="font-bold mb-8 text-7xl">{headline}</h1>
        <h3 class="text-xl">{tagline}</h3>
    </div>
</div>
