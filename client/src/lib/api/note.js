import axios from "axios"
import {PUBLIC_BASE_API_URL} from '$env/static/public'

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
        console.error(error);
        if (error.response?.status === 401) {
            alert(error.response?.data?.message || "Session expired. Log in again.");
            localStorage.removeItem('token');
            localStorage.removeItem('username');
            window.location.href = '/login';
            return null;
        }
        alert(error.response?.data?.message || "Unable to fetch the notes");
        return null;
    }
}

export const GetNote = async (id) => {
    try {
        const res  = await axios.get(`${PUBLIC_BASE_API_URL}/notes/${id}`)
        return res.data
    } catch (error) {
        console.error(error)
        alert(error.response?.data.message || "Unable to fetch the note.")
    }
}

export const DeleteNote = async (id) => {
    const token = localStorage.getItem('token');
    
    if (!token) {
        return "User not logged in."
    }
    try {
        await axios.delete(`${PUBLIC_BASE_API_URL}/notes/${id}`, {
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
}

export const UpdateNote = async (updatedNote) => {
    const token = localStorage.getItem('token');
    
    if (!token) {
        return "User not logged in."
    }
    try {
        const res  = await axios.put(`${PUBLIC_BASE_API_URL}/notes/${updatedNote.id}`, updatedNote, {
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
}