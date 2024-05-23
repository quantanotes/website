<script context="module">
    export { default as layout } from '$/layouts/base.svelte'
</script>

<script>
    import { onMount } from 'svelte'
    import notes from '$/stores/notes.svelte.js'
    import Sidebar from '$/components/common/sidebar/sidebar.svelte'
    import Editor from '$/components/notes/editor.svelte'
    import Navbar from '$/components/notes/navbar.svelte'
    import Actions from '$/components/notes/sidebar/actions.svelte'
    import Share from '$/components/notes/share/share.svelte'

    let showSidebar = $state(false)
    let sidebarWidth = $state(288)

    onMount(async () => {
        await notes.mount()
    })

    function toggleSidebar() {
        showSidebar = !showSidebar
    }
</script>

<div class="absolute h-screen overflow-x-hidden w-screen">
    <Navbar {toggleSidebar} />
    <div class="pt-11 h-full">
        <Sidebar bind:show={showSidebar} bind:width={sidebarWidth} state={notes} >
            {#snippet contextMenu(id, parent)}
                <Actions {id} {parent}>
                    {#snippet share()}
                        <Share {id} />
                    {/snippet}
                </Actions>
            {/snippet}
        </Sidebar>
        <div
            class="transition-all py-2 px-8"
            style={`margin-left: ${showSidebar ? sidebarWidth : 0}px`}
        >
            <Editor />
        </div>
    </div>
</div>
