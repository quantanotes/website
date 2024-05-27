<script context="module">
    export { default as layout } from '$/layouts/page.svelte'
</script>

<script>
    import { untrack } from 'svelte'
    import { router } from '@inertiajs/svelte'
    import SectionDivider from '$/components/atoms/section-divider.svelte'

    const intervals = [
        {
            name: 'Onetime',
            subtext: '',
            discount: 1.0,
            min: 100,
            max: 50000,
            step: 100,
        },
        {
            name: 'Monthly',
            subtext: 'Save 10%',
            discount: 0.9,
            min: 100,
            max: 50000,
            step: 100,
        },
        {
            name: 'Annual',
            subtext: 'Save 20%',
            discount: 0.8,
            min: 1000,
            max: 50000,
            step: 200,
        },
    ]

    let quantity = $state(0)
    let interval = $state('Monthly')
    let currInterval = $derived(intervals.find((i) => i.name == interval))

    $effect(() => {
        quantity = currInterval.min
        untrack(() => quantity)
    })

    async function checkout() {
        router.post('/api/billing/checkout', {
            quantity,
            interval: interval.toLowerCase(),
        })
    }
</script>

<div class="h-12"></div>
<div class="max-w-7xl mx-auto p-4 text-center">
    <div class="flex justify-evenly max-w-4xl mx-auto">
        {#each intervals as i}
            <button
                onclick={() => (interval = i.name)}
                class={`hover:font-bold flex flex-col h-12 justify-between ${interval == i.name && 'font-bold'}`}
            >
                <div class="text-lg uppercase">{i.name}</div>
                <div>{i.subtext}</div>
            </button>
        {/each}
    </div>
    <div class="h-48"></div>
    <div class="mx-auto w-fit">
        {#key interval}
            <input
                class="range w-96"
                type="range"
                bind:value={quantity}
                min={currInterval.min}
                max={currInterval.max}
                step={currInterval.step}
            />
        {/key}
    </div>
    <div class="h-12"></div>
    <div class="mx-auto w-fit">
        <input
            class="border-0 input p-0 w-24"
            type="number"
            bind:value={quantity}
            min={currInterval.min}
            max={currInterval.max}
            step={currInterval.step}
            onblur={() => {
                quantity = Math.min(
                    Math.max(
                        currInterval.min,
                        Math.ceil(quantity / currInterval.step) * currInterval.step,
                    ),
                    currInterval.max,
                )
            }}
        />
        credits / £{(quantity * 0.05 * currInterval.discount).toFixed(2)}
        + {Math.floor(quantity / 500) * 100} bonus credits
    </div>
    <div class="h-12"></div>
    <p>100 bonus credits for every 500 credits purchased</p>
    <div class="h-24"></div>
    <button class="btn" onclick={checkout}>Purchase Credits</button>
    <div class="h-48"></div>
    <p>Looking for more credits? Contact us.</p>
    <div class="h-12"></div>
    <p>Credits allow you to use Quanta's products. 1 credit is worth 1000 LLM tokens.</p>
</div>

<SectionDivider />
