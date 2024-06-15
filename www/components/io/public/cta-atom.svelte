<script>
    import { router } from '@inertiajs/svelte'
    import Input from '../input.svelte'

    const suggestions = [
        'Can you keep me accountable for me meditation habit?',
        'Build an automation for my catering business',
        'Scale and automate my email outreach campaign',
    ]

    let value = $state()

    function send() {
        const redirect = encodeURIComponent(`/maxwell?q=${encodeURIComponent(value)}`)
        router.get(`/signup?redirect=${redirect}`)
    }
</script>

<div class="max-w-4xl mx-auto px-4">
    <Input bind:value placeholder={'Ask Io anything...'} {send} {suggestions} />
</div>
