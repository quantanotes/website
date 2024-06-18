<script>
    import Icon from '@iconify/svelte'
    import DropdownMenu from '$/components/atoms/dropdown-menu.svelte'
    import AccountMenu from './account-menu.svelte'
    import DarkModeToggle from './dark-mode-toggle.svelte'

    let { start, middle, end, ...restProps } = $props()
</script>

<div
    class={restProps.class ||
        'backdrop-blur-lg border-b border-contrast flex h-11 justify-between px-2 w-full'}
>
    <div class="basis-0 flex flex-grow flex-row gap-4 items-center justify-start">
        {#if start}
            {@render start()}
        {/if}
    </div>
    <div class="flex flex-row gap-4 items-center justify-center">
        {#if middle}
            {@render middle()}
        {/if}
    </div>
    <div class="basis-0 flex flex-grow flex-row gap-4 items-center justify-end">
        {#if end}
            {@render end()}
        {/if}
        <DarkModeToggle />
        <DropdownMenu class="icon-btn" placement="bottom-end">
            <Icon icon="mdi:account" />
            {#snippet menu()}
                <AccountMenu />
            {/snippet}
        </DropdownMenu>
    </div>
</div>
