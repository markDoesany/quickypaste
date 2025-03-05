<script>
    import './css/index.css'
    let form = $state({ username: '', password: '' });
    let { onSubmit, buttonText = "Submit" } = $props();

    const handleSubmitForm = (event, sendForm) =>{
        event.preventDefault();
        onSubmit(sendForm);
    }
</script>

<form class="form-container max-w-md mx-auto p-8 rounded-lg shadow-md" onsubmit={(event) => handleSubmitForm(event, form)}>
    <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="username">Username</label>
        <input 
            id="username" 
            type="text" 
            bind:value={form.username} 
            placeholder="Username" 
            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        >
    </div>
    <div class="mb-6">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="password">Password</label>
        <input 
            id="password" 
            type="password" 
            bind:value={form.password} 
            placeholder="Password" 
            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
        >
    </div>
    <div class="flex items-center justify-between">
        <button 
            type="submit" 
            class="focus:outline-none focus:shadow-outline py-2 px-4 rounded cursor-pointer"
        >
            {buttonText}
        </button>
    </div>
    <div class="mt-4 text-center">
        {#if buttonText === "Login"}
            <a href="/register" class="text-blue-500 hover:text-blue-700">Don't have an account? Register</a>
        {:else}
            <a href="/login" class="text-blue-500 hover:text-blue-700">Already have an account? Login</a>
        {/if}
    </div>
</form>