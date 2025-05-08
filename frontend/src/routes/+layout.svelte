<script lang="ts">
    let {children} = $props();
</script>

{@render children()}

<style>
    :global(body) {
        margin: 0;
        padding: 0;
    }
</style>