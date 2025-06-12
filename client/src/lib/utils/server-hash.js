import { v4 as uuidv4 } from 'uuid';

// Simple in-memory store for hashes
const hashMapping = new Map();

export const generateHash = (id) => {
    return uuidv4();
};

export const hashStore = {
    getHash: (id) => {
        return hashMapping.get(id);
    },
    
    setHash: (id, hash) => {
        hashMapping.set(id, hash);
    },
    
    getIDFromHash: (hash) => {
        for (const [id, storedHash] of hashMapping.entries()) {
            if (storedHash === hash) {
                return id;
            }
        }
        return null;
    }
};
