<script>
    import Accordion from '$/components/atoms/accordion.svelte'
    import SectionDivider from '$/components/atoms/section-divider.svelte'

    const faqs = [
        {
            q: 'what is io?',
            a: 'io is an ai agent designed to solve human problems. it uses real world tools available to it from apis and the internet to construct a solution. it is also designed to solve human problems by being empathetic and providing support according to your needs',
        },
        {
            q: 'what tasks can io solve?',
            a: 'anything you can do with a computer, you can do with io. keep in mind there are limitations with the apis io connects to',
        },
        {
            q: "what is io's personality?",
            a: 'io is designed to be empathetic yet sincere. we wanted to build an ai that can connect with the human soul',
        },
    ]
</script>

<div class="p-4">
    <h1 class="flex font-bold justify-between mx-auto text-4xl text-center w-fit">faqs</h1>

    <br />
    <br />
    <br />
    <br />

    <div class="max-w-4xl mx-auto">
        <Accordion items={faqs.map((faq) => ({ title: faq.q, content: faq.a }))} />
    </div>
</div>

<SectionDivider />
