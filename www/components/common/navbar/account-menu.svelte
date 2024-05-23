<script>
    import { router } from '@inertiajs/svelte'
    import auth from '$/stores/auth.svelte.js'
    import Actions from '$/components/atoms/actions.svelte'

    let actions = $derived(
        auth.user
            ? [
                  { name: 'Account', action: () => {} },
                  { name: 'Billing', action: () => {} },
                  { name: 'Sign out', action: auth.signout },
              ]
            : [
                  { name: 'Sign in', action: () => router.get('/signin') },
                  { name: 'Sign up', action: () => router.get('/signup') },
              ],
    )
</script>

<Actions {actions} />
