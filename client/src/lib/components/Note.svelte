<script>
    import { Copy, Edit, X, Share2, Trash } from 'lucide-svelte';
    import { formatReadableDate } from '../../utils/helper';
    import { generateHash, hashStore } from '../../utils/hash';
    import './css/index.css'
    import MessageModal from './MessageModal.svelte';

    // Generate hash when note is created
    if (!hashStore.getHash(note.ID)) {
        const hash = generateHash(note.ID);
        hashStore.setHash(note.ID, hash);
    }

    let showMessage = false;
    let message = '';
    let modalOnClose = () => {};

    export let note;
    export let onDeleteNote;
    export let onUpdateNote;
    export let disable = false;

    let isEditing = false;
    let expanded = false;
    
    let editableContent = note.content;

    const expandNote = () => {
        if (!isEditing) {
            expanded = true;
        }
    };

    const handleToggleEdit = (e) => {
        e.stopPropagation();
        isEditing = !isEditing;
        if (!isEditing) {
            note.content = editableContent;
        }
    };

    const handleSave = () => {
        note.content = editableContent;
        onUpdateNote(note);
        isEditing = false;
    };

    const handleCancel = () => {
        editableContent = note.content;
        isEditing = false;
    };

    const handleCopyNoteToClipboard = async(e) => {
        e.stopPropagation();
        try {
            await navigator.clipboard.writeText(note.content);
            message = 'Note copied to clipboard';
            showMessage = true;
            modalOnClose = () => {
                showMessage = false;
            };
        } catch (error) {
            console.error("Failed to copy: ", error);
        }
    }

    const handleShareNote = async(e) => {
        e.stopPropagation();
        try {
            await navigator.clipboard.writeText(note.shareable_link);
            message = 'Link copied to clipboard';
            showMessage = true;
            modalOnClose = () => {
                showMessage = false;
            };
        } catch (error) {
            console.error("Failed to copy: ", error);
        }
    }

    const handleDeleteNote = (e) => {
        e.stopPropagation();
        onDeleteNote(note.ID);
    }

    const handleKeyDown = (e) => {
        if (e.key === 'Enter' || e.key === ' ') {
            expandNote();
        }
    };

    const handleCloseNote = (e) => {
        e.stopPropagation();
        expanded = false;
        isEditing = false;
    };

</script>
<MessageModal 
        message={message}
        show={showMessage}
        onClose={modalOnClose}
    />
<div 
    class="note {expanded ? 'expanded' : 'collapsed transform transition-transform duration-200 hover:scale-110'} bg-yellow-200 p-4 rounded shadow-lg text-left focus:outline-none flex flex-col gap-2 select-text" 
    onclick={expandNote}
    onkeydown={handleKeyDown}
    tabindex="0"
    data-hash={hashStore.getHash(note.ID)}
    aria-expanded={expanded}
    role="button"
>
    <div class="flex justify-between items-center">
        <div class="text-gray-600 text-sm mt-3">{formatReadableDate(note.UpdatedAt)}</div>
        {#if expanded}
            <button onclick={(e)=> handleCloseNote(e)}>
                <X class="w-5 h-5 text-gray-600 cursor-pointer" />
            </button>
            <span class="absolute -top-3 text-3xl">📌</span>
            {#if disable === false}
                <div class="flex gap-4 absolute bottom-5 right-5">
                    <button onclick = {(e) => handleToggleEdit(e)}>
                        <Edit class="w-6 h-6 max-sm:w-4 max-sm:h-4 text-gray-600 cursor-pointer transform transition-transform duration-300 hover:scale-110 hover:rotate-6" />
                    </button>
                    <button onclick={(e) => handleCopyNoteToClipboard(e)}>
                        <Copy class="w-6 h-6 max-sm:w-4 max-sm:h-4 text-gray-600 cursor-pointer transform transition-transform duration-300 hover:scale-110 hover:rotate-6" />
                    </button>
                    <button onclick={(e) => handleShareNote(e)}>
                        <Share2 class="w-6 h-6 max-sm:w-4 max-sm:h-4 text-gray-600 cursor-pointer transform transition-transform duration-300 hover:scale-110 hover:rotate-6" />
                    </button>
                    <button onclick={(e)=>handleDeleteNote(e)}>
                        <Trash class="w-6 h-6 max-sm:w-4 max-sm:h-4 text-gray-600 cursor-pointer transform transition-transform duration-300 hover:scale-110 hover:rotate-6" />
                    </button>
                </div>
                {:else}
                    <div class="flex gap-4 absolute bottom-5 right-5">
                        <button onclick={(e) => handleCopyNoteToClipboard(e)}>
                            <Copy class="w-6 h-6 text-gray-600 cursor-pointer transform transition-transform duration-300 hover:scale-110 hover:rotate-6" />
                        </button>
                    </div>
            {/if}
        {/if}
    </div>
    {#if isEditing}
        <textarea bind:value={editableContent} class="w-full h-full px-2 py-4 border rounded focus:outline-none"></textarea>
        <div class="flex gap-2 mt-2">
            <button onclick={handleSave} class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded 
                                                    focus:outline-none focus:shadow-outline cursor-pointer max-sm:font-semibold max-sm:text-sm max-sm:px-2">Save</button>
            <button onclick={handleCancel} class="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded 
                                                    focus:outline-none focus:shadow-outline cursor-pointer max-sm:font-semibold max-sm:text-sm max-sm:px-2">Cancel</button>
        </div>
    {:else}
        <div class="text-gray-800 content mt-4 overflow-y-hidden p-2">{editableContent}</div>
    {/if}
    
    {#if !expanded && disable === false}
        <div class="overlay">
            <div class="flex gap-4">
                <button onclick={(e) => handleCopyNoteToClipboard(e)}>
                    <Copy class="w-6 h-6 text-white cursor-pointer" />
                </button>
                <button onclick={(e) => handleShareNote(e)}>
                    <Share2 class="w-6 h-6 text-white cursor-pointer" />
                </button>
                <button onclick={(e) => handleDeleteNote(e)}>
                    <Trash class="w-6 h-6 text-white cursor-pointer" />
                </button>
            </div>
        </div>
    {/if}
</div>

<style>
    .note.expanded {
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    position: fixed;
    z-index: 50;
    width: 90%;
    max-width: 600px;
    max-height: 90vh;
    overflow-y: auto;
}

</style>