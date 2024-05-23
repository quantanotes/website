<script>
    import darkMode from '$/stores/dark-mode.svelte.js'
    import Icon from '@iconify/svelte'
</script>

<button class="icon-btn" onclick={darkMode.toggle}>
    {#if darkMode.darkMode}
        <Icon icon="mdi:weather-sunny" />
    {:else}
        <Icon icon="mdi:weather-night" />
    {/if}
</button>
