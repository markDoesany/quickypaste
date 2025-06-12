import { v4 as uuidv4 } from 'uuid';

// Generate a hash from an ID
export const generateHash = (id) => {
    // Use UUID v4 as our hash generator
    return uuidv4();
};

// Store a mapping of hashes to IDs in localStorage
export const hashStore = {
    getHash: (id) => {
        const mapping = JSON.parse(localStorage.getItem('noteHashes') || '{}');
        return mapping[id];
    },
    
    setHash: (id, hash) => {
        const mapping = JSON.parse(localStorage.getItem('noteHashes') || '{}');
        mapping[id] = hash;
        localStorage.setItem('noteHashes', JSON.stringify(mapping));
    },
    
    getIDFromHash: (hash) => {
        const mapping = JSON.parse(localStorage.getItem('noteHashes') || '{}');
        return Object.entries(mapping).find(([_, storedHash]) => storedHash === hash)?.[0];
    }
};
