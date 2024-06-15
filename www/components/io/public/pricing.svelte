<script>
    import SectionDivider from '$/components/atoms/section-divider.svelte'

    const prices = [
        {
            name: 'Basic',
            price: '£5 / month',
            description:
                'An empathetic life coach that helps you overcome your problems with internet access',
        },
        {
            name: 'Standard',
            price: '£20 / month',
            description:
                'Connects with your data and the outside world for powerful task solving capabilities',
        },
        {
            name: 'Advanced',
            price: '£100 / month',
            description:
                'Orchestrates multiple agents that can perform tasks that require higher levels of planning',
        },
        {
            name: 'Enterprise',
            price: 'Contact us',
            description: 'For advanced use cases to power entire organisation',
        },
    ]
</script>

<div class="p-4">
    <h1 class="flex font-bold justify-between mx-auto text-4xl text-center w-fit">Pricing</h1>

    <br />
    <br />
    <br />
    <br />

    <div class="flex flex-col gap-8 max-w-7xl mx-auto">
        {#each prices as price}
            <div>
                <div class="gap-4 grid grid-cols-3">
                    <h3 class="font-bold text-xl">{price.name}</h3>
                    <div>
                        {price.price}
                    </div>
                    <div class="max-w-sm">
                        {price.description}
                    </div>
                </div>
                <br />
                <br />
                <hr class="border-contrast" />
            </div>
        {/each}
    </div>
</div>

<SectionDivider />
