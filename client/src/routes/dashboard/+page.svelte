<script>
    import { onMount } from 'svelte';
    import { Plus } from 'lucide-svelte';
    import { AddNewNote, GetNotes, DeleteNote, UpdateNote } from '$lib/api/note';
    import ProtectedRoute from '$lib/components/ProtectedRoute.svelte';
    import Note from '$lib/components/Note.svelte';
    import AddNote from '$lib/components/AddNote.svelte';

    let addNote = $state(false);
    let notes = $state([]);

    onMount(async() => {
        const fetchedNotes = await GetNotes();
        if (fetchedNotes.length === 0) return
        
        notes = fetchedNotes;
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

    let showDeleteModal = false;
    let deleteNoteId = null;
    let deleteMessage = '';

    const handleDeleteNote = async(noteId) => {
        deleteNoteId = noteId;
        deleteMessage = 'Are you sure you want to delete this note?';
        showDeleteModal = true;
    };

    const confirmDelete = async() => {
        const res = await DeleteNote(deleteNoteId);
        if (res !== true) return;
        
        notes = notes.filter(note => note.ID !== deleteNoteId);
        showDeleteModal = false;
        deleteNoteId = null;
    };

    const cancelDelete = () => {
        showDeleteModal = false;
        deleteNoteId = null;
    };

    const handleUpdateNote = async(note) => {
        const res = await UpdateNote(note);
        if (!res) return

        notes = notes.map(n => n.ID === note.ID ? note : n);
    }
</script>

<ProtectedRoute redirect="/login"/>
<main class="p-5">
    {#if notes && notes.length > 0}
        <div class="grid gap-5 mt-5 grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 justify-items-center sm:justify-items-start">
            {#each notes as note}
                <Note 
                    note={note}
                    onDeleteNote={() => handleDeleteNote(note.ID)}
                    onUpdateNote={handleUpdateNote}
                />
            {/each}
        </div>
    {:else}
        <div class="text-center text-gray-600 py-10">
            No notes found
        </div>
    {/if}

    {#if showDeleteModal}
        <MessageModal 
            message={deleteMessage}
            show={showDeleteModal}
            onClose={cancelDelete}
            onConfirm={confirmDelete}
            isConfirm={true}
        />
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

      

