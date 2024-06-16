import { mergeAttributes, Node } from '@tiptap/core'

export default Node.create({
    name: 'title',
    content: 'inline*',
    group: 'title',

    addOptions() {
        return {
            HTMLAttributes: {
                class: 'title text-5xl',
            },
        }
    },

    parseHTML() {
        return [
            {
                tag: 't',
            },
        ]
    },

    renderHTML({ HTMLAttributes }) {
        return ['h1', mergeAttributes(this.options.HTMLAttributes, HTMLAttributes), 0]
    },
})
