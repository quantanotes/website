<script>
    import { inertia, router } from '@inertiajs/svelte'
    import auth from '$/stores/auth.svelte.js'
    import {
        emailSchema,
        matchingSchema,
        nameSchema,
        passwordSchema,
        usernameSchema,
    } from '$/utils/schema.js'
    import Form from '$/components/atoms/form.svelte'

    let data = $state({})
    let message = $state('')
    let passwordConfirmSchema = $derived(matchingSchema(data['password'], 'Passwords'))

    async function action(data) {
        const response = await fetch('/api/auth/signup', {
            method: 'POST',
            body: JSON.stringify(data),
        })

        if (response.ok) {
            router.get('/verify')
            auth.refreshDetails()
        }

        const message = await response.text()
        return message || 'An unexpected error has occurred, try again later.'
    }
</script>

<Form
    inputs={[
        { name: 'email', label: 'Email', type: 'text', schema: emailSchema },
        { name: 'fullName', label: 'Full name', type: 'text', schema: nameSchema },
        { name: 'preferredName', label: 'Preferred name', type: 'text', schema: nameSchema },
        {
            name: 'username',
            label: 'Username',
            type: 'text',
            schema: usernameSchema,
        },
        {
            name: 'password',
            label: 'Password',
            type: 'password',
            schema: passwordSchema,
        },
        {
            name: 'confirmPassword',
            label: 'Confirm password',
            type: 'password',
            schema: passwordConfirmSchema,
        },
    ]}
    submitLabel="Sign up"
    {action}
    bind:message
    bind:data
/>
<p class="text-center">
    Have an account? Sign in
    <a use:inertia class="anchor" href="/signin">here</a>
</p>
<br />
