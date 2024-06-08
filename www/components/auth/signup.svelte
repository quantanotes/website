<script>
    import { inertia, page, router } from '@inertiajs/svelte'
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
    let verify = $state(false)
    let passwordConfirmSchema = $derived(matchingSchema(data['password'], 'Passwords'))
    let redirect = $derived($page.props?.redirect || '/')

    async function action(data) {
        const response = await fetch(`/api/auth/signup?redirect=${encodeURIComponent(redirect)}`, {
            method: 'POST',
            body: JSON.stringify(data),
        })

        if (response.ok) {
            verify = true
        }

        const message = await response.text()
        return message || 'An unexpected error has occurred, try again later.'
    }
</script>

<br />
<br />
<h1 class="mx-auto text-lg w-fit">Get started with Quanta</h1>
<br />
<br />

{#if !verify}
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
{:else}
    <p class="text-center">A verification link was sent to your email.</p>
{/if}
