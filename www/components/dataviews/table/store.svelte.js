import { onMount } from 'svelte'
import { dialog } from '$/stores/portals.svelte.js'

export default function createTableStore(state, addColumnMenu) {
    // Data states
    let current = $derived(state.current)
    // Column data is stored as a field in the metadata JSON of each node
    let columns = $derived(deserialiseColumns(state.current))
    // Rows are nodes of category record
    let data = $derived(
        state.map[state.current].children
            .map((id) => state.map[id])
            .filter((node) => node.category === 'record')
            .map((node) => ({ id: node.id, data: deserialiseRow(node.id) })),
    )
    // Frontend state
    let columnRefs = $state(columns?.map(() => null))
    let columnWidths = $state(columns?.map(() => 0))
    let resizingColumn = $state()
    let tableHeight = $state()

    onMount(() => {
        // Table rows are children in the hierarchy
        state.children(current)
    })

    function addColumn(name, type) {
        // Check if column's type is valid
        if (!['bool', 'number', 'text', 'note'].includes(type)) {
            throw new Error('Column is invalid type')
        }
        // Get a copy of metadata/columns for mutations
        let columns = deserialiseColumns(current)
        if (columns.find((column) => column.name === name)) {
            throw new Error('Column already exists')
        }
        columns = [...columns, { name, type }]
        serialiseColumns(current, columns)
    }

    function addRow() {
        state.create(current, 'record')
    }

    function removeColumn(index) {
        // Get a copy of metadata/columns for mutations
        let columns = deserialiseColumns(current)
        columns = columns.filter((_, i) => i !== index)
        serialiseColumns(current, columns)
    }

    function handleColumnResize(index) {
        resizingColumn = index

        function handleMouseMove(event) {
            columnWidths[resizingColumn] = event.clientX - columnRefs[resizingColumn].offsetLeft
        }

        function handleMouseUp() {
            window.removeEventListener('mousemove', handleMouseMove)
            window.removeEventListener('mouseup', handleMouseUp)
            resizingColumn = undefined
        }

        window.addEventListener('mousemove', handleMouseMove)
        window.addEventListener('mouseup', handleMouseUp)
    }

    // TODO: add server side validation

    function showAddColumnMenu(event) {
        dialog.show(event.target, addColumnMenu, 'bottom', '', event)
    }

    function updateColumn(id, column, value) {
        let row = deserialiseRow(id)
        switch (column.type) {
            case 'bool':
                if (typeof value !== 'boolean')
                    throw new Error('Value is not correct type for column')
                break
            case 'number':
                if (typeof value !== 'number')
                    throw new Error('Value is not correct type for column')
            case 'text':
                if (typeof value !== 'string')
                    throw new Error('Value is not correct type for column')
            case 'note':
                if (typeof value !== 'string')
                    throw new Error('Value is not correct type for column')
            default:
                throw new Error('Column is invalid type')

        }
        row[column.name] = value
        serialiseRow(id, row)
        state.update(id)
    }

    // Using JSON serialisation in to content for row
    // Not the best idea but fits nicely in to the data model
    function deserialiseRow(id) {
        let node = state.map[id]
        try {
            let data = JSON.parse(node.content)
            return data
        } catch {
            return {}
        }
    }

    function serialiseRow(id, data) {
        // We should ensure each column in data still exists here
        state.map[id].content = JSON.stringify(data)
    }

    function serialiseColumns(id, columns) {
        const metadata = state.map[id].metadata
        state.map[id].metadata = { ...metadata, columns }
        state.update(id)
    }

    function deserialiseColumnsAndMetadata(id) {
        const metadata = state.map[id].metadata
        let columns = metadata?.columns
        if (!Array.isArray(columns)) {
            columns = []
        }
        return { columns, metadata }
    }

    function deserialiseColumns(id) {
        return deserialiseColumnsAndMetadata(id).columns
    }

    return {
        get columns() { return columns },
        get data() { return data },
        get columnRefs() { return columnRefs },
        get columnWidths() { return columnWidths },
        get tableHeight() { return tableHeight },
        set tableHeight(newTableHeight) { tableHeight = newTableHeight },

        addColumn,
        addRow,
        removeColumn,
        handleColumnResize,
        showAddColumnMenu,
        updateColumn,
    }
}