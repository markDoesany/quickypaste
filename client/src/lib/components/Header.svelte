<script>
    import { goto } from '$app/navigation';
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import './css/index.css';

    let isLoggedIn = $state(false);

    onMount(() => {
        const token = localStorage.getItem('token');
        isLoggedIn = !!token;
    });

    const handleLogout = () => {
        localStorage.removeItem('token');
        isLoggedIn =  false
        goto('/login');
    };
</script>

<header class="header">
    <button class="text-2xl font-bold cursor-pointer max-sm:text-lg" onclick={() => goto('/dashboard')}>QuickPaste</button>
    <nav class="cursor-pointer">
        {#if isLoggedIn}
            <button 
                onclick={handleLogout} 
                class="focus:outline-none focus:shadow-outline max-sm:text-sm max-sm:p-2"
            >
                Logout
            </button>
        {:else}
            <a 
                href="/login" 
                class=" focus:outline-none focus:shadow-outline max-sm:text-sm max-sm:p-2"
            >
                Login
            </a>
            <a 
                href="/register" 
                class="focus:outline-none focus:shadow-outline max-sm:text-sm max-sm:p-2"
            >
                Sign Up
            </a>
        {/if}
    </nav>
</header>