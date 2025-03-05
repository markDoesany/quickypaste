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
        console.log('Notes:', notes);
    })

    const handleAddClick = () => {
        console.log('Add button clicked');
        addNote = true;
    };

    const handleCloseAddNote = () => {
        console.log('Close add note');
        addNote = false;
    };

    const handleAddNewNote = async(note) => {
        console.log('Add new note',note.content);
        const res = await AddNewNote(note);
        if (!res) return
        
        notes = [...notes, res]
        handleCloseAddNote();
    };

    const handleDeleteNote = async(noteId) => {
        console.log('Delete note', noteId);
        const res = await DeleteNote(noteId);
        if (res !== true) return
        notes = notes.filter(note => note.ID !== noteId);
    };

    const handleUpdateNote = async(note) => {
        console.log('Update note', note);
        const res = await UpdateNote(note);
        if (!res) return

        notes = notes.map(n => n.ID === note.ID ? note : n);
    }
</script>

<ProtectedRoute redirect="/login"/>
    {#if notes.length > 0}
        <main class="p-5">
            <div class="grid gap-5 mt-5 grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 justify-items-center sm:justify-items-start">
                {#each notes as note, index}
                    <Note 
                        note={note}
                        onDeleteNote={handleDeleteNote}
                        onUpdateNote={handleUpdateNote}
                    />
                {/each}
            </div>

            {#if addNote}
                <AddNote onClose={handleCloseAddNote} onSaveNewNote={handleAddNewNote}/>
            {/if}

        </main>
        {:else}
            <p class="text-center">Create new note</p>
        {/if}
        <button 
            onclick={handleAddClick} 
            class="fixed bottom-4 right-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full focus:outline-none focus:shadow-outline cursor-pointer z-30"
        >
            <Plus class="w-6 h-6" />
        </button>

