<script>
    import { marked } from 'marked'

    let { node } = $props()
</script>

<div
    class="border border-contrast max-w-none my-4 p-4 prose dark:prose-invert prose-neutral w-full"
>
    <h1 class="text-5xl">{node.title}</h1>
    <hr />
    {@html marked.parse(node.content)}
</div>
