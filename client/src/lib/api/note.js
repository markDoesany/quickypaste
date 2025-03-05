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
        console.log("New Note:",res.data)
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
        console.log("Notes:",res.data)
        return res.data
    } catch (error) {
        console.error(error)
        alert(error.response?.data.message || "Unable to fetch the notes")
    }
}

export const GetNote = async (id) => {
    try {
        const res  = await axios.get(`${PUBLIC_BASE_API_URL}/notes/${id}`)
        console.log("Note:",res.data)
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
        const res  = await axios.delete(`${PUBLIC_BASE_API_URL}/notes/${id}`, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        })
        console.log("Deleted Note:",res.data)
        alert(res.data.message)
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
        const res  = await axios.put(`${PUBLIC_BASE_API_URL}/notes/${updatedNote.ID}`, updatedNote, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        })
        console.log("Updated Note:",res.data)
        return res.data
    } catch (error) {
        console.error(error)
        alert(error.response?.data.message || "Unable to save changes.")
    }
}