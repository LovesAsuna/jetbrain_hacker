<script lang="ts">
    let {data} = $props();
    import Error from '$lib/error.svelte'
    import Loading from '$lib/loading.svelte'

    function onclick() {
        navigator.clipboard.writeText((document.getElementById("textContent") as HTMLTextAreaElement).value)
            .then(() => {
                alert('config is copied to clipboard');
            })
            .catch(() => {
                alert('copy failed！please copy manually');
            });
    }
</script>

{#await data.loader()}
    <Loading/>
{:then c }
    <textarea readonly id="textContent">{c.config}</textarea>
    <button class="copy-button" {onclick}>Copy Config</button>
{:catch _}
    <Error/>
{/await}

<style>
    textarea {
        width: 100vw;
        height: 100vh;
        padding: 15px;
        border: none;
        border-radius: 0;
        font-size: 16px;
        line-height: 1.5;
        color: #333;
        background: rgba(255, 255, 255, 0.5);
        resize: none;
        box-sizing: border-box;
        white-space: pre-wrap;
        overflow-wrap: break-word;
    }

    textarea[readonly] {
        cursor: default;
    }

    textarea::placeholder {
        color: #999;
        font-style: italic;
    }

    .copy-button {
        position: fixed;
        top: 10px;
        right: 10px;
        padding: 10px 20px;
        background-color: #007BFF;
        color: white;
        border: none;
        border-radius: 5px;
        font-size: 14px;
        cursor: pointer;
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
        transition: background-color 0.3s;
    }

    .copy-button:hover {
        background-color: #0056b3;
    }

    .copy-button:active {
        background-color: #004085;
    }
</style>