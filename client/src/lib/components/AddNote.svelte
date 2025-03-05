<script> 
    import './css/index.css'
    let {onSaveNewNote, 
         onClose, 
         content="", 
         date=new Date().toISOString().split('T')[0]} = $props()
</script>

<button class="fixed inset-0 bg-opacity-50 backdrop-blur-xs z-40" onclick={onClose}><div>backdrop</div></button>

<div 
    class="addNote bg-yellow-200 p-4 rounded shadow-lg text-left focus:outline-none flex flex-col gap-2 select-text" 
    tabindex="0"
    aria-expanded={true}
    role="button"
>
    <div class="flex justify-between items-center">
        <div class="text-gray-600 text-sm mt-3">{date}</div>
        </div>
        <textarea bind:value={content} class="w-full h-full px-2 py-4 border rounded focus:outline-none"></textarea>
        <div class="flex gap-2 mt-2">
            <button onclick={() => onSaveNewNote({content:content})} class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline cursor-pointer">Save</button>
            <button onclick={onClose} class="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline cursor-pointer">Cancel</button>
        </div>
</div>