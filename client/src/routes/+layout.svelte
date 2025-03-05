<!-- filepath: /c:/Users/emman/OneDrive/Documents/Project/QuickyPaste/client/src/routes/__layout.svelte -->
<script>
    import { page } from '$app/stores';
    import Header from '$lib/components/Header.svelte';
    import '../app.css';

    function shouldShowHeader(pathname) {
        return !pathname.startsWith('/notes/');
    }

	let { children } = $props();
    let showHeader = $state(shouldShowHeader($page.url.pathname));
	$effect(()=>{
		showHeader = shouldShowHeader($page.url.pathname);
	})

</script>


<main >
    {#if showHeader}
        <Header />
    {/if}
    {@render children()}
</main>