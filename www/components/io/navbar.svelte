
<script>
	import Icon from '@iconify/svelte'
	import io from '$/stores/io.svelte.js'
    import Navbar from '$/components/common/navbar/navbar.svelte'

	let { toggleSidebar } = $props()
	let current = $derived(io.map[io.current])
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
