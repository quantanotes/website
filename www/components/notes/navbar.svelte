<script>
	import Icon from '@iconify/svelte'
	import notes from '$/stores/notes.svelte.js'
    import Navbar from '../common/navbar/navbar.svelte'

	let { toggleSidebar } = $props()
	let current = $derived(notes.map[notes.current])
</script>

<div class="fixed w-full z-50">
	<Navbar>
		{#snippet start()}
			<button class="icon-btn" onclick={toggleSidebar}>
				<Icon icon="mdi:menu" />
			</button>
			<div class="h-6 text-contrast">
				{current && current.title}
			</div>
		{/snippet}
	</Navbar>
</div>