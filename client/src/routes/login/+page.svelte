<script>
    import AuthenticatedRoute from "$lib/components/AuthenticatedRoute.svelte";
    import Form from "$lib/components/Form.svelte";
    import { login } from "$lib/api/auth";

    const handleLogin = async({username, password}) =>{
        if (username === "" || password === ""){
            alert("Please fill in all fields")
            return
        }

       await login(username, password);
}
</script>

<style>
     .login {
        font-family: 'Indie Flower', cursive; 
    }
</style>

<AuthenticatedRoute redirect="/dashboard" />
<main class="login flex justify-center">
    <div class="w-full max-w-md p-5 ">
        <h2 class="text-center text-2xl font-bold mb-6">Login Form</h2>
        <Form onSubmit={handleLogin} buttonText="Login"/>
    </div>
</main>