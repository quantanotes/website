<script context="module">
    export { default as layout } from '$/layouts/base.svelte'
</script>

<script>
    import { onMount } from 'svelte'
    import io from '$/stores/io.svelte.js'
    import Sidebar from '$/components/common/sidebar/sidebar.svelte'
    import Input from '$/components/io/input.svelte'
    import Thread from '$/components/io/thread.svelte'
    import Navbar from '$/components/io/navbar.svelte'
    import Actions from '$/components/io/sidebar/actions.svelte'
    import Settings from '$/components/io/settings/settings.svelte'

    let input = $state('')
    let thread = $state([])
    let container = $state()
    let abort = $state()
    let generating = $state(false)
    let showSidebar = $state(false)
    let sidebarWidth = $state(288)

    onMount(async () => {
        await io.mount()
    })

    async function send() {
        if (input.length <= 0 || generating) {
            return
        }

        thread = [...thread, { role: 'user', content: input, citations: [] }]
        input = ''
        abort = new AbortController()
        await scroll()

        try {
            const stream = io.chat(thread, abort)
            const reply = { role: 'assistant', content: '', citations: [] }
            let first = true
            for await (const part of stream) {
                if (first) {
                    thread.push(reply)
                    first = false
                }
                thread[thread.length - 1].content += part.content
                thread[thread.length - 1].citations.push(...part.citations)
            }
        } catch (error) {
            handleError(error)
        } finally {
            generating = false
        }
    }

    function handleError(error) {
        const content = thread.at(-1).content
        input = content
        thread = [...thread.slice(0, thread.length - 1)]
        if (error.status == 401) {
            
        } else {
            console.error(error)
        }
    }

    async function scroll() {
        if (!container) return
        await container.scroll({
            top: container.scrollHeight,
            behavior: 'smooth',
        })
    }

    function toggleSidebar() {
        showSidebar = !showSidebar
    }
</script>

<div class="absolute h-screen w-screen">
    <Navbar {toggleSidebar} />
    <div class="pt-11 h-full w-full">
        <Sidebar bind:show={showSidebar} bind:width={sidebarWidth} state={io} >
            {#snippet contextMenu(id, parent)}
                <Actions {id} {parent} >
                    {#snippet settings()}
                        <Settings {id} />
                    {/snippet}
                </Actions>
            {/snippet}
        </Sidebar>
        
        <div
            class="h-full transition-all w-full"
            style={`padding-left: ${showSidebar ? sidebarWidth : 0}px`}
        >
            <div class="flex flex-col gap-4 h-full items-center justify-center max-w-7xl mx-auto p-4 w-full">
                <Thread bind:thread bind:container />
                <div class="flex-grow w-full">
                    <Input bind:value={input} {send} />
                </div>
            </div>
        </div>
    </div>
</div>
