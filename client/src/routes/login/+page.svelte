<script>
    import AuthenticatedRoute from "$lib/components/AuthenticatedRoute.svelte";
    import Form from "$lib/components/Form.svelte";
    import { login } from "$lib/api/auth";
    import MessageModal from "$lib/components/MessageModal.svelte";

    let showMessage = false;
    let message = '';
    let modalOnClose = () => {};

    let loading = false;

    const handleLogin = async({username, password}) =>{
        if (username === "" || password === ""){
            message = "Please fill in all fields";
            showMessage = true;
            modalOnClose = () => {
                showMessage = false;
            };
            return;
        }

        loading = true;
        try {
            await login(username, password);
        } finally {
            loading = false;
        }
    }
</script>

<style>
     .login {
        font-family: 'Indie Flower', cursive; 
    }
</style>

<AuthenticatedRoute redirect="/dashboard" />
<main class="login flex justify-center">
    <div class="w-full max-w-md p-5 text-center">
        <h1 class="text-4xl font-bold mb-4">QuickyPaste</h1>
        <p class="text-xl mb-8">Your personal notes manager</p>
        <div class="p-4 rounded-lg shadow-md">
            <h2 class="text-center text-2xl font-bold mb-6">Login Form</h2>
        <Form onSubmit={handleLogin} buttonText="Login"/>
    </div>
    <MessageModal 
        message={message}
        show={showMessage}
        onClose={modalOnClose}
    />
</main>