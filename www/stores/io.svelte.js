import createCRUD from './crud.svelte'

export default {
    ...createCRUD('maxwell'),

    async *chat(thread, abort) {
        let reader

        try {
            const response = await fetch('/api/maxwell/chat', {
                body: JSON.stringify(thread),
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
                return
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
    },
}