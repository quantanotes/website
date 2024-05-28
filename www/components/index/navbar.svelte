<script>
    import { fade } from 'svelte/transition'
    import { inertia } from '@inertiajs/svelte'
    import Navbar from '$/components/common/navbar/navbar.svelte'

    const links = [
        { name: 'home', href: '/'},
        { name: 'products', children: [
            { name: 'notes', href: '/notes' },
            { name: 'io', href: '/io' },
            { name: 'space', href: '/space' },
            { name: 'heisenberg', href: '/heisenberg' },
        ]},
    ]

    let active = $state(null)

    function show(link) {
        active = link
    }

    function hide() {
        active = null 
    }
</script>

{#if active}
    <div
        class="absolute backdrop-blur-md h-screen mt-11 overflow-hidden w-full z-10"
        transition:fade={{ duration: 200 }}
    >
    </div>
{/if}

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="relative" onmouseleave={hide}>
    <Navbar>
        {#snippet middle()}
            <div class="grid grid-cols-2 w-full">
                {#each links as link}
                    <a
                        class="hover:font-bold transition text-center uppercase"
                        href={link.href || '#'}
                        onmouseenter={() => link.children && show(link)}
                        use:inertia
                    >
                        {link.name}
                    </a>
                {/each}
            </div>
        {/snippet}
    </Navbar>
    {#if active}
        <div
            class="absolute bg-primary gap-4 p-4 w-full z-20"
            transition:fade={{ duration: 200 }}
        >
            {#each active.children as link}
                <a
                    class="block hover:font-bold max-w-4xl mx-auto transition uppercase"
                    href={link.href || '#'}
                    use:inertia
                >
                    {link.name}
                </a>
            {/each}
        </div>
    {/if}
</div>
