<div class="loading_container">
    <div class="sh1"></div>
    <div class="sh2"></div>
    <h4 class="lt">Loading data...</h4>
</div>

<style>
    :global(body) {
        min-width: 100vw;
        min-height: 100vh;
        overflow-x: hidden;
        overflow-y: hidden;
        margin: 0;
    }

    .loading_container {
        background-color: #333;
        width: 100vw;
        height: 100vh;
        padding-top: calc(45vh - 25px);
    }

    .sh1 {
        width: 0;
        height: 0;
        border-style: solid;
        border-width: 50px 50px 0 0;
        border-color: #354952 transparent transparent transparent;
        margin: 0 auto;
        animation: shk1 1s ease-in-out infinite normal;
    }

    .sh2 {
        width: 0;
        height: 0;
        border-style: solid;
        border-width: 0 0 50px 50px;
        border-color: transparent transparent #df1a54 transparent;
        margin: -50px auto 0;
        animation: shk2 1s ease-in-out infinite alternate;
    }

    /** animation starts here **/
    @keyframes shk1 {
        0% {
            transform: rotate(-360deg);
        }

        100% {
        }
    }

    @keyframes shk2 {
        0% {
            transform: rotate(360deg);
        }
        100% {
        }
    }

    .lt {
        color: #fff;
        font-family: 'Roboto', sans-serif;
        margin: 30px auto;
        text-align: center;
        font-weight: 100;
        letter-spacing: 10px;
    }

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