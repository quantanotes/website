<script>
    import notes from '$/stores/notes.svelte.js'
    import Table from '../dataviews/table/table.svelte'
    import Editor from './editor.svelte'

    let { showSidebar, sidebarWidth } = $props()

    let view = $derived(notes.map[notes.current]?.view || 'note')
</script>

<div class="transition-all py-2 px-8" style={`margin-left: ${showSidebar ? sidebarWidth : 0}px`}>
    {#if view === 'note'}
        <Editor />
    {:else if view == 'table'}
        <Table />
    {/if}
</div>
