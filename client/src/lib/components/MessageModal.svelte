<script>
    export let message = '';
    export let show = false;
    export let onClose = () => {};

    const handleClose = () => {
        onClose();
    };

    const handleKeydown = (e) => {
        if (e.key === 'Escape') {
            handleClose();
        }
    };
</script>

{#if show}
    <div 
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
        tabindex="0"
        on:keydown={handleKeydown}
    >
        <div class="bg-white rounded-lg shadow-lg p-6 max-w-sm w-full relative border-2 border-gray-800">
            <button
                class="absolute top-2 right-2 text-gray-800 hover:text-gray-600"
                aria-label="Close"
                on:click={handleClose}
            >
                &times;
            </button>
            <div class="text-center text-lg font-semibold mb-4 font-['Indie_Flower'] text-gray-800">
                {message}
            </div>
            <button
                class="mt-2 px-4 py-2 bg-yellow-400 hover:bg-yellow-500 rounded font-bold text-gray-800 border-2 border-gray-800 transition-all duration-300 transform hover:scale-105"
                on:click={handleClose}
            >
                OK
            </button>
        </div>
    </div>
{/if}