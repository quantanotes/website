<script context="module">
    export { default as layout } from '$/layouts/page.svelte'
</script>

<script>
    import Signin from '$/components/auth/signin.svelte'
</script>

<div class="flex h-[calc(100vh-44px)] items-center justify-center w-full">
    <div>
        <Signin />
    </div>
</div>
