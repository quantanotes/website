<script>
	import Icon from '@iconify/svelte'
	import notes from '$/stores/notes.svelte.js'
    import Navbar from '../common/navbar/navbar.svelte'
    import Input from './input.svelte'

	let { toggleSidebar, inputValue, inputSend } = $props()
	let current = $derived(notes.map[notes.current])
</script>

<div class="fixed w-full z-50">
	<Navbar class="backdrop-blur-lg border-b border-contrast grid grid-flow-col h-11 justify-stretch px-2 w-full">
		{#snippet start()}
			<div class="flex flex-row gap-4 items-center justify-start">
				<button class="icon-btn" onclick={toggleSidebar}>
					<Icon icon="mdi:menu" />
				</button>
				<div class="h-6 text-contrast">
					{current && current.title}
				</div>
			</div>
		{/snippet}

		{#snippet middle()}
			<Input bind:value={inputValue} bind:send={inputSend} />
		{/snippet}
	</Navbar>
</div>