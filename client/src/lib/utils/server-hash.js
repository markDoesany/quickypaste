import { v4 as uuidv4 } from 'uuid';

export const generateHash = (id) => {
    return uuidv4();
};

export const hashStore = {
    getHash: (id) => {
        const mapping = JSON.parse(process.env.NOTE_HASHES || '{}');
        return mapping[id];
    },
    
    setHash: (id, hash) => {
        const mapping = JSON.parse(process.env.NOTE_HASHES || '{}');
        mapping[id] = hash;
        process.env.NOTE_HASHES = JSON.stringify(mapping);
    },
    
    getIDFromHash: (hash) => {
        const mapping = JSON.parse(process.env.NOTE_HASHES || '{}');
        return Object.entries(mapping).find(([_, storedHash]) => storedHash === hash)?.[0];
    }
};
