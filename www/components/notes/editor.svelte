<script>
    import { untrack } from 'svelte'
    import notes from '$/stores/notes.svelte.js'
    import Tiptap from '$/features/tiptap/tiptap.svelte'
    import { blur } from 'svelte/transition'

    let setContent = $state()
    let saveTimeout = $state()
    let current = $state()

    $effect(() => {
        if (notes.current === current || !notes.current) return
        current = notes.current
        let map = untrack(() => notes.map)
        let note = map[current]
        let title = note.title || ''
        let content = note.content || ''
        setContent(title, content)
    })

    function onupdate(editor) {
        const markdown = editor.storage.markdown.getMarkdown()
        notes.map[notes.current].title = new DOMParser()
            .parseFromString(markdown, 'text/html')
            .getElementsByTagName('h1')
            .item(0).textContent
        notes.map[notes.current].content = markdown.substring(markdown.indexOf('</h1>') + 5).trim()
        resetSaveTimeout()
    }

    function resetSaveTimeout() {
        clearTimeout(saveTimeout)
        saveTimeout = setTimeout(() => {
            if (!notes.current) return
            let title = notes.map[notes.current].title || ''
            let content = notes.map[notes.current].content || ''
            notes.update(notes.current, title, content)
        }, 1000 * 2)
    }
</script>

<div class="h-full max-w-3xl mb-64 mx-auto prose prose-neutral dark:prose-invert px-4 py-2 w-full">
    <div class={!notes.current && 'hidden'} transition:blur>
        {#key notes.current}
            <Tiptap {onupdate} bind:setContent />
        {/key}
    </div>
</div>
