<script>
    import Accordion from '$/components/atoms/accordion.svelte'
    import TitledSection from '$/components/common/titled-section.svelte'

    const faqs = [
        {
            q: 'What is Io?',
            a: 'Io is an AI agent designed to solve human problems. It uses real world tools available from the internet to construct solutions. Is is designed with humanity in mind, being empathetic and providing support according to your needs',
        },
        {
            q: 'What tasks can Io solve?',
            a: 'Anything you can do with a computer, you can do with Io. keep in mind there are limitations with the apis io connects to',
        },
        {
            q: "What is Io's personality?",
            a: 'Io is designed to be empathetic yet sincere. We wanted to build an AI that can connect with each individual by reflecting their personalities.',
        },
    ]
</script>

<TitledSection title={'Frequently asked questions'}>
    <Accordion items={faqs.map((faq) => ({ title: faq.q, content: faq.a }))} />
</TitledSection>
