<script>
    import { onMount, onDestroy } from 'svelte'
    import { Editor } from '@tiptap/core'
    import Document from '@tiptap/extension-document'
    import Link from '@tiptap/extension-link'
    import Placeholder from '@tiptap/extension-placeholder'
    import StarterKit from '@tiptap/starter-kit'
    import { Markdown } from 'tiptap-markdown'
    import Title from './extensions/title'
    import Draggable from './extensions/draggable'
    import { EditorState } from '@tiptap/pm/state'
    import { hideDragHandle } from './extensions/draggable/draggable'

    /** @type {{ onupdate: ((editor: Editor) => void), setContent: (title: string, content: string) => void  }} */
    let { onupdate, setContent } = $props()
    /** @type {Editor} */
    let editor = $state()
    let element = $state()

    const extensions = [
        Document.extend({ content: 'title block*' }),
        Draggable,
        Link,
        Markdown,
        Placeholder.configure({
            placeholder: ({ node }) =>
                node.type.name === 'title' ? 'Title...' : 'Create anything...',
            emptyNodeClass:
                'cursor-text before:absolute before:content-[attr(data-placeholder)] before:text-contrast before-pointer-events-none',
        }),
        StarterKit.configure({ document: false }),
        Title,
    ]

    onMount(() => {
        const options = {
            editorProps: { attributes: { class: 'focus:outline-none' } },
            element,
            extensions,
            onTransaction: () => (editor = editor),
            onUpdate: ({ editor }) => onupdate(editor),
        }

        editor = new Editor(options)

        setContent = function (title, content) {
            editor.commands.setContent(`<h1>${title}</h1>\n\n${content}`)
            editor.view.updateState(
                EditorState.create({
                    doc: editor.state.doc,
                    plugins: editor.state.plugins,
                    schema: editor.state.schema,
                }),
            )
        }
    })

    onDestroy(() => {
        if (editor) {
            editor.destroy()
        }
    })
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div bind:this={element} onmouseleave={hideDragHandle}></div>

<style lang="postcss">
    :global(.tiptap) {
        @apply px-10;
    }

    :global(.ProseMirror p.is-editor-empty:first-child::before) {
        @apply text-contrast;
        content: attr(data-placeholder);
        float: left;
        height: 0;
        pointer-events: none;
    }
</style>
