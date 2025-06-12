import { GetNote } from '$lib/api/note.js';
import { hashStore } from '$lib/utils/server-hash.js';

export async function load({ params }) {
    const { id } = params;
    const noteId = hashStore.getIDFromHash(id);
    
    if (!noteId) {
        return { status: 404, error: 'Note not found' };
    }
    
    const note = await GetNote(noteId);
    if (!note) {
        return { status: 404, error: 'Note not found' };
    }
    
    return { note };
}