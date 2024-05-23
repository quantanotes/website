<script>
    import { inertia, router } from '@inertiajs/svelte'
    import Form from '$/components/atoms/form.svelte'

    async function action(data) {
        const response = await fetch('/api/auth/verify', {
            method: 'POST',
            body: JSON.stringify(data),
        })

        if (response.ok) {
            router.get('/')
        }

        const message = await response.text()
        return message || 'An unexpected error has occurred, try again later.'
    }
</script>

<Form
    inputs={[{ name: 'code', label: 'Verification code', type: 'text' }]}
    submitLabel="Verify"
    {action}
/>
<br />
<p class="text-center">
    New to Quanta? Sign up
    <a use:inertia class="anchor" href="/signin">here</a>
</p>
<br />
