let user = $state(undefined)

export default {
    get user() {
        return user
    },

    get modal() {
        return modal 
    },

    async signout() {
        await fetch('/api/auth/signout', { method: 'POST' })
    },

    async refreshDetails() {
        const response = await fetch('/api/auth/details', { method: 'POST' })
        const details = await response.json()
        user = details || undefined
    },
}