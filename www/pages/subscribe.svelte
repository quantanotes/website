<script context="module">
    export { default as layout } from '$/layouts/page.svelte'
</script>

<script>
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

<div class="flex flex-row gap-4 items-center justify-center max-w-6xl mx-auto h-[calc(100vh-44px)]">
    {#each prices as price}
        <button class="btn flex flex-col h-64 justify-between text-left w-64">
            <div>{price.name}</div>
        </button>
    {/each}
</div>
