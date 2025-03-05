<script>
    import AuthenticatedRoute from "$lib/components/AuthenticatedRoute.svelte";
    import Form from "$lib/components/Form.svelte";
    import {register} from "$lib/api/auth";

    const handleRegister = async({username, password}) => {
        await register(username, password);
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
    <div class="w-full max-w-md p-5">
        <h2 class="text-center text-2xl font-bold mb-6">Register Form</h2>
        <Form onSubmit={handleRegister} buttonText="Register"/>
    </div>
</main>