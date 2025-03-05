export const formatReadableDate = (isoString) => {
    const date = new Date(isoString);

    const options = {
        weekday: 'long',
        year: 'numeric', 
        month: 'long', 
        day: 'numeric',
        hour: 'numeric', 
        minute: 'numeric', 
        second: 'numeric',
    };

    return date.toLocaleDateString('en-US', options);
}