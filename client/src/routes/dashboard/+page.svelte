<script>
    import { onMount } from 'svelte';
    import { Plus } from 'lucide-svelte';
    import { AddNewNote, GetNotes, DeleteNote, UpdateNote } from '$lib/api/note';
    import ProtectedRoute from '$lib/components/ProtectedRoute.svelte';
    import Note from '$lib/components/Note.svelte';
    import AddNote from '$lib/components/AddNote.svelte';

    let addNote = $state(false);
    let notes = $state([]);
    let filteredNotes = $state([]);
    let sortOption = $state('date-desc');
    let filterDateRange = $state({
        startDate: null,
        endDate: null
    });
    let searchQuery = $state('');
    let showEditedOnly = $state(false);

    onMount(async() => {
        const fetchedNotes = await GetNotes();
        if (fetchedNotes.length === 0) return
        
        notes = fetchedNotes;
        updateFilteredNotes();
    })

    // Watch for changes in notes, sortOption, filterDateRange, searchQuery, and showEditedOnly
    $: updateFilteredNotes();

    function updateFilteredNotes() {
        let filtered = [...notes];

        // Apply date range filter
        if (filterDateRange.startDate || filterDateRange.endDate) {
            filtered = filtered.filter(note => {
                const noteDate = new Date(note.created_at);
                const startDate = filterDateRange.startDate ? new Date(filterDateRange.startDate) : null;
                const endDate = filterDateRange.endDate ? new Date(filterDateRange.endDate) : null;
                
                if (startDate && noteDate < startDate) return false;
                if (endDate && noteDate > endDate) return false;
                return true;
            });
        }

        // Apply search query filter
        if (searchQuery) {
            const query = searchQuery.toLowerCase();
            filtered = filtered.filter(note => 
                note.title.toLowerCase().includes(query) || 
                note.content.toLowerCase().includes(query)
            );
        }

        // Apply edited filter
        if (showEditedOnly) {
            filtered = filtered.filter(note => note.edited_at !== null);
        }

        // Sort the notes
        filtered.sort((a, b) => {
            switch (sortOption) {
                case 'date-desc':
                    return new Date(b.created_at) - new Date(a.created_at);
                case 'date-asc':
                    return new Date(a.created_at) - new Date(b.created_at);
                case 'title-asc':
                    return a.title.localeCompare(b.title);
                case 'title-desc':
                    return b.title.localeCompare(a.title);
                default:
                    return 0;
            }
        });

        filteredNotes = filtered;
    }

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

    const handleDeleteNote = async(noteId) => {
        if (!window.confirm("Delete note?")) return

        const res = await DeleteNote(noteId);
        if (res !== true) return
        notes = notes.filter(note => note.ID !== noteId);
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
        <div class="flex flex-col gap-4">
            <div class="flex justify-between items-center">
                <h1 class="text-2xl font-bold">Dashboard</h1>
                <div class="flex gap-4">
                    <select
                        class="px-3 py-1 bg-yellow-200 rounded text-gray-800 border-2 border-gray-800"
                        bind:value={sortOption}
                    >
                        <option value="date-desc">Newest First</option>
                        <option value="date-asc">Oldest First</option>
                        <option value="title-asc">Title A-Z</option>
                        <option value="title-desc">Title Z-A</option>
                    </select>
                    <button
                        class="px-4 py-2 bg-yellow-400 hover:bg-yellow-500 rounded font-bold text-gray-800"
                        on:click={handleAddClick}
                    >
                        Add Note
                    </button>
                </div>
            </div>

            <div class="flex flex-wrap gap-4">
                <div class="flex items-center gap-2">
                    <label class="text-gray-800">Date Range:</label>
                    <input
                        type="date"
                        class="px-2 py-1 bg-yellow-200 rounded text-gray-800 border-2 border-gray-800"
                        bind:value={filterDateRange.startDate}
                    />
                    <input
                        type="date"
                        class="px-2 py-1 bg-yellow-200 rounded text-gray-800 border-2 border-gray-800"
                        bind:value={filterDateRange.endDate}
                    />
                </div>
                <div class="flex items-center gap-2">
                    <label class="text-gray-800">Search:</label>
                    <input
                        type="text"
                        class="px-2 py-1 bg-yellow-200 rounded text-gray-800 border-2 border-gray-800"
                        bind:value={searchQuery}
                        placeholder="Search notes..."
                    />
                </div>
                <div class="flex items-center gap-2">
                    <label class="text-gray-800">Show Edited Only:</label>
                    <input
                        type="checkbox"
                        class="bg-yellow-200 rounded text-gray-800 border-2 border-gray-800"
                        bind:checked={showEditedOnly}
                    />
                </div>
            </div>
            <div class="grid gap-5 mt-5 grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 justify-items-center sm:justify-items-start">
                {#each filteredNotes as note}
                    <Note 
                        note={note}
                        onDeleteNote={handleDeleteNote}
                        onUpdateNote={handleUpdateNote}
                    />
                {/each}
            </div>
        </div>
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

      

