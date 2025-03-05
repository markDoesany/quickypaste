<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';

    let { children } = $$props;

    onMount(() => {
        const token = localStorage.getItem('token');
        if (token) {
            goto('/dashboard');
        }
    });
</script>

<slot />