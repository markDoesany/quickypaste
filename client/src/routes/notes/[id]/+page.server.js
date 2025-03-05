import { GetNote } from '$lib/api/note.js'

export async function load({params}){
        const { id } = params
        const note = await GetNote(id)
        return {note}
    }