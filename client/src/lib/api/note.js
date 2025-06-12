import axios from "axios"
import {PUBLIC_BASE_API_URL} from '$env/static/public'
import { hashStore } from '../utils/hash';

export const AddNewNote = async (note) => {
    const token = localStorage.getItem('token');
    
    if (!token) {
        return "User not logged in."
    }

    try {
        const res  = await axios.post(`${PUBLIC_BASE_API_URL}/notes`, note, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        })
        return res.data
    } catch (error) {
        console.error(error)
        alert(error.response?.data.message || "Failed to create new note.")
    }
}

export const GetNotes = async () => {
    const token = localStorage.getItem('token');
    
    if (!token) {
        return "User not logged in."
    }

    try {
        const res  = await axios.get(`${PUBLIC_BASE_API_URL}/notes`, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        })
        return res.data
    } catch (error) {
        console.error(error)
        alert(error.response?.data.message || "Unable to fetch the notes")
    }
}

export const GetNote = async (hash) => {
    const id = hashStore.getIDFromHash(hash);
    if (!id) return null;
    try {
        const res  = await axios.get(`${PUBLIC_BASE_API_URL}/notes/${id}`)
        return res.data
    } catch (error) {
        console.error(error)
        alert(error.response?.data.message || "Unable to fetch the note.")
    }
}

export const DeleteNote = async (hash) => {
    const id = hashStore.getIDFromHash(hash);
    if (!id) return false;
    try {
        const token = localStorage.getItem('token');
        
        if (!token) {
            return "User not logged in."
        }
        try {
            const res  = await axios.delete(`${PUBLIC_BASE_API_URL}/notes/${id}`, {
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                }
            })
            return true
        } catch (error) {
            console.error(error)
            alert(error.response?.data.message || "Unable to delete the note.")
            return false
        }
    } catch (error) {
        console.error(error)
    }
}

export const UpdateNote = async (note) => {
    const id = hashStore.getIDFromHash(note.ID);
    if (!id) return false;
    try {
        const token = localStorage.getItem('token');
        
        if (!token) {
            return "User not logged in."
        }
        try {
            const res  = await axios.put(`${PUBLIC_BASE_API_URL}/notes/${id}`, note, {
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                }
            })
            return res.data
        } catch (error) {
            console.error(error)
            alert(error.response?.data.message || "Unable to save changes.")
        }
    } catch (error) {
        console.error(error)
        alert(error.response?.data.message || "Unable to save changes.")
    }
}