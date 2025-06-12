<script>
    import Note from "$lib/components/Note.svelte";
	let { data } = $props();
</script>

<main class="p-10 flex justify-center items-center min-h-screen bg-gray-100">
    <Note note={data.note} disable={true}/>
</main>