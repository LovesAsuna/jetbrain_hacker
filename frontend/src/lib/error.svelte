<!-- svelte-ignore a11y_invalid_attribute -->
<div class="error_container">
    <nav class="shelf">
        <a class="book home-page">Home page</a>
        <a class="book about-us">About us</a>
        <a class="book contact">Contact</a>
        <a class="book faq">F.A.Q.</a>

        <span class="book not-found"></span>

        <span class="door left"></span>
        <span class="door right"></span>
    </nav>
    <h1>Error 404</h1>
    <p>The page you're looking for can't be found</p>
</div>


<style>
    @import url(https://fonts.googleapis.com/css?family=Dancing+Script:400,700);

    .error_container {
        height: calc(100vh - 2rem);
        margin-top: unset;
        padding-top: 2rem;
        background-color: #446072;
        font-family: monospace;
        color: white;
        overflow: hidden;
    }

    h1 {
        margin-top: 2rem;
        text-align: center;
    }
    h1 + p {
        text-align: center;
    }
    .shelf {
        position: relative;
        width: 30rem;
        height: 14rem;
        margin: 0 auto;
        border: 0.5rem solid #374d5b;
        border-radius: 0.5rem;
        background-color: rgba(255, 255, 255, 0.1);
        perspective: 130rem;
        box-shadow: inset 0 0 2rem rgba(0, 0, 0, 0.2);
    }
    .door {
        position: absolute;
        width: 14.8rem;
        height: 14rem;
        display: flex;
        box-sizing: border-box;
        padding: 1rem;
        background-color: #374d5b;
        align-items: center;
        box-shadow: 0 1px 1px rgba(0, 0, 0, 0.3);
        outline: 1px solid transparent;
    }
    .door::before {
        width: 1.5rem;
        height: 1.5rem;
        border-radius: 50%;
        background-color: rgba(0, 0, 0, 0.1);
        content: "";
    }
    .door.left {
        border-radius: 0 0.75rem 0.75rem 0;
        justify-content: flex-end;
        -webkit-animation:
                leftDoorOpen 3.5s ease-out forwards 1s,
                leftDoorFlap 15s linear infinite forwards 9s;
        animation:
                leftDoorOpen 3.5s ease-out forwards 1s,
                leftDoorFlap 15s linear infinite forwards 9s;
        transform-origin: 0 0 0;
    }
    .door.right {
        right: 0;
        border-radius: 0.75rem 0 0 0.75rem;
        -webkit-animation:
                rightDoorOpen 3s ease-out forwards 1.5s,
                rightDoorFlap 10s linear infinite forwards 8s;
        animation:
                rightDoorOpen 3s ease-out forwards 1.5s,
                rightDoorFlap 10s linear infinite forwards 8s;
        transform-origin: 100% 0 0;
    }
    .book {
        position: absolute;
        box-sizing: border-box;
        padding: 0.8rem 4rem 0.8rem 2rem;
        border-radius: 0.25rem;
        background-color: rgba(255, 255, 255, 0.1);
        color: white;
        font-size: 1rem;
        text-transform: uppercase;
        letter-spacing: 0.1rem;
        cursor: pointer;
        box-shadow: inset 0 0 0.75rem rgba(255, 255, 255, 0.1);
    }
    .book.home-page {
        transform: rotate(-90deg) translate(-12.4rem, 3rem);
        transform-origin: 0;
    }
    .book.about-us {
        transform: rotate(-100deg) translate(-13.4rem, 6.1rem);
        transform-origin: 0;
        outline: 1px solid transparent;
    }
    .book.contact {
        right: 2rem;
        bottom: 0.2rem;
        border-radius: 0.3rem 0 0 0.3rem;
    }
    .book.faq {
        right: 0.8rem;
        bottom: 3.3rem;
        border-radius: 0.3rem 0 0 0.3rem;
    }
    .book.not-found {
        width: 12rem;
        height: 3.5rem;
        border: 1px dashed rgba(255, 255, 255, 0.3);
        background-color: transparent;
        transform: rotate(-90deg) translate(-12rem, 13rem) scale(1);
        transform-origin: 0;
        cursor: default;
        -webkit-animation: bookFadeOut 1s 3s infinite forwards;
        animation: bookFadeOut 1s 3s infinite forwards;
    }
    .book.not-found::after {
        display: block;
        width: 10rem;
        padding-left: 5rem;
        background-image: url("data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' width='80' height='30'><path fill='rgb(255, 255, 255)' d='M7.688,3.737C6.1,10.409,4.624,16.982,2.475,23.517c-0.424,1.29,0.724,2.338,1.957,1.916 c5.879-2.021,11.743-4.107,17.409-6.696c1.246-0.572,0.443-2.366-0.815-1.932c-3.856,1.333-7.695,2.69-11.565,3.959 c2.879-2.526,5.485-5.215,9.013-7.17c4.441-2.459,9.299-4.109,14.281-4.915c10.903-1.772,22.052,0.562,31.979,5.04 c9.241,4.162,12.895,8.725,13.164,4.942c0.121-1.69-5.57-4.953-14.125-8.714C53.19,5.292,41.094,3.539,29.636,6.058 c-7.122,1.567-16.708,5.374-22.064,11.544c1.359-4.455,2.461-9.028,3.032-13.563C10.852,2.091,8.12,1.927,7.688,3.737z'/></svg>");
        background-repeat: no-repeat;
        background-size: 4rem;
        background-position: left center;
        font-family: "Dancing Script";
        text-transform: lowercase;
        font-size: 25px;
        content: "this page can't be found";
        transform: rotate(90deg) translate(6rem, -1rem);
        opacity: 0.3;
    }
    .book:hover:not(.not-found) {
        background-color: rgba(255, 255, 255, 0.2);
    }
    @-webkit-keyframes leftDoorOpen {
        60% {transform: rotateY(-115deg)}
        100% {transform: rotateY(-110deg)}
    }
    @keyframes leftDoorOpen {
        60% {transform: rotateY(-115deg)}
        100% {transform: rotateY(-110deg)}
    }
    @-webkit-keyframes rightDoorOpen {
        60% {transform: rotateY(125deg)}
        100% {transform: rotateY(120deg)}
    }
    @keyframes rightDoorOpen {
        60% {transform: rotateY(125deg)}
        100% {transform: rotateY(120deg)}
    }
    @-webkit-keyframes rightDoorFlap {
        0% { transform: rotateY(120deg)}
        5% {transform: rotateY(125deg)}
        15% {transform: rotateY(117deg)}
        25% {transform: rotateY(123deg)}
        30% {transform: rotateY(120deg)}
        100% {transform: rotateY(120deg)}
    }
    @keyframes rightDoorFlap {
        0% { transform: rotateY(120deg)}
        5% {transform: rotateY(125deg)}
        15% {transform: rotateY(117deg)}
        25% {transform: rotateY(123deg)}
        30% {transform: rotateY(120deg)}
        100% {transform: rotateY(120deg)}
    }
    @-webkit-keyframes leftDoorFlap {
        0% { transform: rotateY(-110deg)}
        5% {transform: rotateY(-115deg)}
        15% {transform: rotateY(-107deg)}
        25% {transform: rotateY(-113deg)}
        30% {transform: rotateY(-110deg)}
        100% {transform: rotateY(-110deg)}
    }
    @keyframes leftDoorFlap {
        0% { transform: rotateY(-110deg)}
        5% {transform: rotateY(-115deg)}
        15% {transform: rotateY(-107deg)}
        25% {transform: rotateY(-113deg)}
        30% {transform: rotateY(-110deg)}
        100% {transform: rotateY(-110deg)}
    }
    @-webkit-keyframes bookFadeOut {
        50% {border: 1px dashed rgba(255, 255, 255, 0.1);}
    }
    @keyframes bookFadeOut {
        50% {border: 1px dashed rgba(255, 255, 255, 0.1);}
    }
</style>