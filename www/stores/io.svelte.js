import createNodeStore from './nodes.svelte'

export const io = createNodeStore('thread')

io.chat = async function* (thread, abort) {
    let reader

    try {
        if (!io.current) {
            throw new Error("You're a cheeky rapscallion if you're seeing this.")
        }
        const threadID = io.current

        const response = await fetch('/api/io/chat', {
            body: JSON.stringify({ threadID, thread }),
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'POST',
            signal: abort.signal,
        })

        if (!response.ok) {
            const error = new Error(response.statusText)
            error.status = response.status
            throw error
        }

        if (!response.body) {
            throw new Error('No response body')
        }

        let buffer = ''
        reader = response.body.pipeThrough(new TextDecoderStream()).getReader()

        while (true) {
            const { value, done } = await reader.read()
            if (done) break

            buffer += value

            const lines = buffer.split('\n')
            if (!value.endsWith('\n')) {
                buffer = lines.pop()
            } else {
                buffer = ''
            }

            for (const line of lines) {
                if (!line.trim()) continue
                try {
                    const obj = JSON.parse(line)
                    yield obj
                } catch (_) {
                    throw new Error('An unexpected error has occurred, try again later.')
                }
            }
        }
    } catch (error) {
        throw error
    } finally {
        if (reader) reader.releaseLock()
    }
}

export default io
