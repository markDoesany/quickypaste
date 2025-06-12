<script>
    import AuthenticatedRoute from "$lib/components/AuthenticatedRoute.svelte";
    import Form from "$lib/components/Form.svelte";
    import {register} from "$lib/api/auth";
    import { goto } from '$app/navigation';
    import MessageModal from "$lib/components/MessageModal.svelte";

    let showMessage = $state(false);
    let message = $state('');
    let modalOnClose = () => {
        showMessage = false;
    };

    let loading = $state(false);

    const handleRegister = async({username, password}) => {
        if (username === "" || password === ""){
            message = "Please fill in all fields";
            showMessage = true;
            modalOnClose();
            return;
        }

        loading = true;
        try {
            await register(username, password);
            message = 'Registration successful! Redirecting to login...';
            showMessage = true;
            modalOnClose();
            goto('/login')
        } finally {
            loading = false;
        }
    }
</script>

<style>
     @import url('https://fonts.googleapis.com/css2?family=Indie+Flower&display=swap');
     .register {
        font-family: 'Indie Flower', cursive; 
    }
</style>

<AuthenticatedRoute redirect="/dashboard" />
<main class="register flex items-center justify-center">
    <div class="w-full max-w-md p-5 text-center">
        <h1 class="text-4xl font-bold mb-4">QuickyPaste</h1>
        <p class="text-xl mb-8">Your personal notes manager</p>
        <div class="p-4 rounded-lg shadow-md">
            <h2 class="text-center text-2xl font-bold mb-6">Register Form</h2>
        <Form onSubmit={handleRegister} buttonText="Register" loading={loading}/>
        <MessageModal 
            message={message}
            show={showMessage}
            onClose={modalOnClose}
        />
    </div>
</main>