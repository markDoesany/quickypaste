<script>
    import { onMount } from 'svelte';
    import { Plus } from 'lucide-svelte';
    import { AddNewNote, GetNotes, DeleteNote, UpdateNote } from '$lib/api/note';
    import ProtectedRoute from '$lib/components/ProtectedRoute.svelte';
    import Note from '$lib/components/Note.svelte';
    import AddNote from '$lib/components/AddNote.svelte';

    let addNote = $state(false);
    let notes = $state([]);
    let loading = $state(false)

    onMount(async() => {
        loading = true;
        const fetchedNotes = await GetNotes();
        if (fetchedNotes.length === 0) {
            loading = false;
            return
        }
        
        notes = fetchedNotes;
        loading = false;
    })

    const handleAddClick = () => {
        addNote = true;
    };

    const handleCloseAddNote = () => {
        addNote = false;
    };

    const handleAddNewNote = async(note) => {
        const res = await AddNewNote(note);
        if (!res) return
        
        notes = [...notes, res]
        handleCloseAddNote();
    };

    let showDeleteModal = $state(false);
    let currentNoteId = $state(null);

    const handleDeleteNote = async(noteId) => {
        currentNoteId = noteId;
        showDeleteModal = true;
    };

    const confirmDelete = async() => {
       try {
            loading = true;
            const res = await DeleteNote(currentNoteId);
            if (res !== true){
                loading = false;
                return
            } 
            notes = notes.filter(note => note.id !== currentNoteId);
            showDeleteModal = false;
       } catch (error) {
            console.error(error)
       } finally{
            loading = false
       }
    };

    const cancelDelete = () => {
        showDeleteModal = false;
        currentNoteId = null;
    };

    const handleUpdateNote = async(note) => {
        const res = await UpdateNote(note);
        if (!res) return

        notes = notes.map(n => n.id === note.id ? note : n);
    }
</script>

<ProtectedRoute redirect="/login"/>
<main class="p-5">
    {#if notes && notes.length > 0}
        <div class="grid gap-5 mt-5 grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 justify-items-center sm:justify-items-start">
            {#each notes as note}
                <Note 
                    note={note}
                    onDeleteNote={handleDeleteNote}
                    onUpdateNote={handleUpdateNote}
                />
            {/each}
        </div>
    {:else if loading}
        <p class="text-center">Loading notes...</p>
    {:else}
        <p class="text-center">Create new note</p>
    {/if}

    {#if addNote}
        <AddNote onClose={handleCloseAddNote} onSaveNewNote={handleAddNewNote}/>
    {/if}

    <button 
        onclick={handleAddClick} 
        class="fixed bottom-4 right-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full focus:outline-none focus:shadow-outline cursor-pointer z-30"
    >
        <Plus class="w-6 h-6" />
    </button>
</main>

<!-- Delete Confirmation Modal -->
{#if showDeleteModal}
    <div class="fixed inset-0 backdrop-blur-sm flex items-center justify-center z-99">
        <div class="bg-white rounded-lg shadow-lg p-6 max-w-sm w-full relative border-2 border-gray-800">
            <button
                class="absolute top-2 right-2 text-gray-800 hover:text-gray-600"
                aria-label="Close"
                onclick={cancelDelete}
            >
                &times;
            </button>
            <div class="text-center text-lg font-semibold mb-4 text-gray-800">
                Delete Note?
            </div>
            <div class="flex gap-4 justify-center">
                <button
                    class="px-4 py-2 bg-red-500 hover:bg-red-600 rounded font-bold text-white border-2 border-gray-800 transition-all duration-300 transform hover:scale-105"
                    onclick={confirmDelete}
                >
                    {loading ? 'Deleting...' : 'Delete'}
                </button>
                <button
                    class="px-4 py-2 bg-gray-500 hover:bg-gray-600 rounded font-bold text-white border-2 border-gray-800 transition-all duration-300 transform hover:scale-105"
                    onclick={cancelDelete}
                >
                    Cancel
                </button>
            </div>
        </div>
    </div>
{/if}
