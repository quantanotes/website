<script context="module">
    export { default as layout } from '$/layouts/base.svelte'
</script>

<script>
    import { onMount } from 'svelte'
    import notes from '$/stores/notes.svelte.js'
    import Sidebar from '$/components/common/sidebar/sidebar.svelte'
    import Navbar from '$/components/notes/navbar.svelte'
    import Actions from '$/components/notes/sidebar/actions.svelte'
    import Share from '$/components/notes/share/share.svelte'
    import View from '$/components/notes/view.svelte'

    let showSidebar = $state(false)
    let sidebarWidth = $state(288)
    let input = $state('')

    onMount(async () => {
        await notes.mount()
    })

    function send() {
        if (input[0] === '/') {
            handleCommands(input)
        }
        input = ''
    }

    function toggleSidebar() {
        showSidebar = !showSidebar
    }

    // I'm sure we'll refactor this to another place

    function parseCommand(input) {
        const parts = input.split(' ')
        const command = parts[0].toLowerCase().slice(1)
        const args = parts.slice(1).map((arg) => arg.toLowerCase())
        return { command, args }
    }

    function handleCommands(input) {
        const { command, args } = parseCommand(input)
        switch (command) {
            case 'view':
                handleViewCommand(args[0])
                break
        }
    }

    function handleViewCommand(view) {
        if (!['note', 'table'].includes(view)) {
            return
        }
        notes.map[notes.current].view = view
    }
</script>

<div class="absolute h-screen overflow-auto w-screen">
    <Navbar {toggleSidebar} bind:inputValue={input} inputSend={send} />
    <div class="pt-11 h-full">
        <Sidebar bind:show={showSidebar} bind:width={sidebarWidth} state={notes}>
            {#snippet contextMenu(id, parent)}
                <Actions {id} {parent}>
                    {#snippet share()}
                        <Share {id} />
                    {/snippet}
                </Actions>
            {/snippet}
        </Sidebar>
        <View bind:showSidebar bind:sidebarWidth />
    </div>
</div>
