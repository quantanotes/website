// Creates a generic CRUD controller for Notes and Io
// Category is just the type of node we want to mount as our root node in the hierarchy
export default function createNodeStore(category) {
    /// Storage states

    // Hierarchy root
    let root = $state('')
    // Key value store for each node containing metadata, children and parent
    let map = $state({})

    /// Frontend states

    // Current focus of frontend
    let current = $state('')
    // Drag focus for moving nodes
    let moving = $state('')
    // Drag context for moving nodes
    let dragover = $state('')

    function call(route, data) {
        return fetch(`/api/node/${route}`, {
            method: 'POST',
            body: data && JSON.stringify(data)
        })
    }

    return {
        get root() {
            return root
        },

        get map() {
            return map
        },

        get current() {
            return current
        },

        get moving() {
            return moving
        },

        set moving(newMoving) {
            moving = newMoving
        },

        get dragover() {
            return dragover
        },

        set dragover(newDragover) {
            dragover = newDragover
        },

        // Initialises the state
        async mount() {
            if (root !== '') return

            const response = await call(`root/${category}`)
            if (!response.ok) return

            const newRoot = { ... await response.json(), children: [] }
            current = root = newRoot.id
            map[root] = newRoot
        },

        // Sets the current focus
        open(id) {
            if (current)
                this.update(current, map[current].title, map[current].content)
            current = id
        },

        async create(parent) {
            const response = await call('create', { parent })
            if (!response.ok) return
            // Refresh the children
            // TODO: manually append the new id        
            await this.children(parent)
        },

        async update(id, title, content) {
            const response = await call('update', { id, title, content })
            // Incase we need to add new code
            if (!response.ok) return
        },

        async remove(id, parent) {
            if (!parent) return
            if (this.ascendants(current).includes(id))
                this.current = ''

            const response = await call('delete', { id, parent })
            if (!response.ok) return

            delete map[id]
            map[parent].children = map[parent].children.filter((child) => child !== id)
        },

        async move(id, to) {
            const response = await call('move', { id, to })
            if (!response.ok) return

            const from = map[id].parent
            // Set new parent
            map[id].parent = to
            // Remove child from old parent
            map[from].children = map[from].children.filter((child) => child !== id)
            // Add the new child to the new parent
            map[to].children = [...map[to].children, id]
        },

        async children(id) {
            const response = await call('children', { id })
            if (!response.ok) return

            const nodes = await response.json() || []
            // Ensure each node has a children field
            nodes.forEach((node) => (map[node.id] = { ...node, children: [] }))
            // Populate the parents children
            map[id].children = nodes.map((node) => node.id)
        },

        async publish(id) {
            const response = await call('publish', { id })
            if (!response.ok) return
            map[id].public = 1
        },

        async unpublish(id) {
            const response = await call('unpublish', { id })
            if (!response.ok) return
            map[id].public = 0
        },

        ascendants(id, result = []) {
            const parent = map[id]?.parent
            if (map[parent]) {
                result.push(parent)
                return this.ascendants(parent, result)
            }
            return result
        }
    }
}