<script>
    import Icon from '@iconify/svelte'
    import SectionDivider from '$/components/atoms/section-divider.svelte'

    const products = [
        {
            name: 'notes',
            description: 'a minimalistic yet powerful AI note taking app for your knowledge',
        },
        {
            name: 'io',
            description: 'an AI agent that connects with your tools and data',
        },
        {
            name: 'space',
            description: 'an open community to share knowledge',
        },
        {
            name: 'heisenberg',
            description: 'a powerful AI native database for semantic search',
        },
    ]

    let ref = $state()
    let refs = $state(Array(products.length))
    let index = $state(0)

    function next() {
        const newIndex = index >= products.length - 1 ? 0 : index + 1
        scrollTo(newIndex)
    }

    function previous() {
        const newIndex = index <= 0 ? products.length - 1 : index - 1
        scrollTo(newIndex)
    }

    function scrollTo(index) {
        refs[index]?.scrollIntoView({ behavior: 'smooth', block: 'nearest', inline: 'start' })
    }

    function handleScroll() {
        const width = ref.clientWidth
        const centerX = width / 2

        const distances = refs.map((item) => {
            const rect = item.getBoundingClientRect()
            const itemCenterX = rect.left + rect.width / 2
            return Math.abs(itemCenterX - centerX)
        })

        index = distances.indexOf(Math.min(...distances))
    }
</script>

<div class="p-4">
    <div class="flex justify-between max-w-4xl mx-auto">
        <h1 class="text-xl">Products</h1>
        <div class="flex gap-4">
            <button onclick={previous} class="icon-btn"><Icon icon="mdi:chevron-left" /></button>
            <button onclick={next} class="icon-btn"><Icon icon="mdi:chevron-right" /></button>
        </div>
    </div>

    <br />
    <br />

    <div
        bind:this={ref}
        class="flex gap-8 no-scrollbar overflow-x-auto snap-x snap-mandatory w-full"
        onscroll={handleScroll}
    >
        <div class="snap-center shrink-0 w-[calc(50%-160px)]"></div>
        {#each products as product, i}
            <div
                bind:this={refs[i]}
                class="border border-secondary flex flex-col h-96 justify-between my-8 p-4 shrink-0 snap-always snap-center hover:shadow-lg transition-shadow w-80"
            >
                <div>
                    {product.name}
                </div>
                <div>
                    {product.description}
                </div>
            </div>
        {/each}
        <div class="snap-center shrink-0 w-[calc(100%-160px)]"></div>
    </div>
</div>

<SectionDivider />
