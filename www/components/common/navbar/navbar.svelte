<script>
    import Icon from '@iconify/svelte'
    import Dropdown from '$/components/atoms/dropdown.svelte'
    import AccountMenu from './account-menu.svelte'
    import DarkModeToggle from './dark-mode-toggle.svelte'
    
    let { start, middle, end } = $props()
</script>

<div class="backdrop-blur-lg border-b border-contrast h-11 grid grid-cols-3 px-2 py-2 w-full">
    <div class="flex flex-row gap-4 items-center justify-start">{#if start}{@render start()}{/if}</div>
    <div class="flex flex-row gap-4 items-center justify-start">{#if middle}{@render middle()}{/if}</div>
    <div class="flex flex-row gap-4 items-center justify-end">
        {#if end}{@render end()}{/if}
        <DarkModeToggle />
        <Dropdown class="icon-btn" placement="bottom-end">
            <Icon icon="mdi:account" />
            {#snippet menu()}
                <AccountMenu />
            {/snippet}
        </Dropdown>
    </div>
</div>
