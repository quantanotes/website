<script>
    import TitledSection from '$/components/common/titled-section.svelte'
import CtaAtom from './cta-atom.svelte'
</script>

<TitledSection title="it begins with a question">
    {#snippet content()}
        <CtaAtom />
    {/snippet}
</TitledSection>
