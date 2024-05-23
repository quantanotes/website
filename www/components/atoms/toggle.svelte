<script>
    let { active = false, onToggle }  = $props()

    function toggle() {
        active = !active
        onToggle(active)
    }
</script>

<button
    class={`
        bg-primary border border-secondary cursor-pointer h-4 relative transition-all w-8
        ${active && 'shadow-secondary shadow-md'}
    `}
    onclick={toggle}
>
    <div
        class="absolute bg-secondary inset-0 w-1/2 transition-transform transform"
        style="transform: translateX({active ? '100%' : '0%'})"
    ></div>
</button>
