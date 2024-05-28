<script>
    import { inertia, page, router } from '@inertiajs/svelte'
    import auth from '$/stores/auth.svelte.js'
    import Form from '$/components/atoms/form.svelte'

    let redirect = $derived($page.props?.redirect || '/')

    async function action(data) {
        const response = await fetch('/api/auth/signin', {
            method: 'POST',
            body: JSON.stringify(data),
        })

        if (response.ok) {
            router.get(redirect)
            auth.refreshDetails()
        }

        const message = await response.text()
        return message || 'An unexpected error has occurred, try again later.'
    }
</script>

<Form
    inputs={[
        { name: 'email', label: 'Email', type: 'text' },
        { name: 'password', label: 'Password', type: 'password' },
    ]}
    submitLabel="Sign in"
    {action}
/>
<br />
<p class="text-center">
    New to Quanta? Sign up
    <a use:inertia class="anchor" href="/signup">here</a>
</p>
<br />
