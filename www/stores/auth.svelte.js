import { router } from '@inertiajs/svelte'

let user = $state(undefined)

export default {
    get user() {
        return user
    },

    get modal() {
        return modal
    },

    async signout() {
        router.post('/api/auth/signout')
        user = undefined
    },

    async refreshDetails() {
        try {
            const response = await fetch('/api/auth/details', { method: 'POST' })
            if (!response.ok) {
                user = undefined
            } else {
                user = await response.json()
            }
        } catch {}
    },
}
