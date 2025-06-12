<script>
    export let message = '';
    export let show = false;
    export let onClose = () => {};
    export let onConfirm = () => {};
    export let isConfirm = false;  // Set to true for confirm/cancel modal

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
        class="fixed inset-0 backdrop-blur-sm flex items-center justify-center z-99"
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
            {#if isConfirm}
                <div class="flex gap-4 justify-center">
                    <button
                        class="mt-2 px-4 py-2 bg-red-500 hover:bg-red-600 rounded font-bold text-white border-2 border-gray-800 transition-all duration-300 transform hover:scale-105"
                        on:click={() => {
                            onConfirm();
                            handleClose();
                        }}
                    >
                        Delete
                    </button>
                    <button
                        class="mt-2 px-4 py-2 bg-gray-500 hover:bg-gray-600 rounded font-bold text-white border-2 border-gray-800 transition-all duration-300 transform hover:scale-105"
                        on:click={handleClose}
                    >
                        Cancel
                    </button>
                </div>
            {:else}
                <button
                    class="mt-2 px-4 py-2 bg-yellow-400 hover:bg-yellow-500 rounded font-bold text-gray-800 border-2 border-gray-800 transition-all duration-300 transform hover:scale-105"
                    on:click={handleClose}
                >
                    OK
                </button>
            {/if}
        </div>
    </div>
{/if}