<script lang="ts">
    import type {Product} from "./+page";
    import JbIcon from './icon.svelte'
    import '$lib/tailwind.css'

    let {data} = $props();
    let productSources: Array<Product> = $state([]);
    let filterInput: string = $state("");
    let productList: Array<Product> = $derived.by(() => {
        return productSources.filter((product) => {
            return product.name.toLowerCase().includes(filterInput.toLowerCase());
        });
    })

    function filterSources(family: Array<string>) {
        productSources = data.products.filter((product) => {
            return family.includes(product.productFamilyName);
        });
    }

    async function copyLicense(event: Event, code: string) {
        let p = event.target as HTMLParagraphElement;
        try {
            let response = await fetch(`api/rpc/license?code=${code}`);
            let licenseCode = await response.json();
            await navigator.clipboard.writeText(licenseCode.licenseCode)
            p.setAttribute('data-content', 'Copied!');
        } catch (error) {
            p.setAttribute('data-content', 'Copy failed!');
        }
        setTimeout(
            () => {
                p.setAttribute('data-content', 'Copy to clipboard');
            }, 1500
        );
    }
</script>

<!-- svelte-ignore a11y_invalid_attribute -->
<JbIcon filter={filterSources}/>

<div id="container">
    <header class="top-[2.3%] bg-(--card-bg) text-(--text-main) z-99 w-[80%] mx-auto rounded-[16px] shadow-[0_8px_40px_-12px_rgba(0,0,0,0.3)] transition-[transform,box-shadow] duration-250 ease-in-out hover:translate-y-[2px] hover:shadow-[0_4px_20px_0_rgba(0,0,0,0.12)] sticky flex items-center pl-6 pr-6">
        <p class="block my-[1em] mx-0 break-words">
            Download <a
                href="https://gitee.com/ja-netfilter/ja-netfilter/releases/download/2022.2.0/ja-netfilter-2022.2.0.zip"
                title="Download jetbra first"
                class="text-(--accent) no-underline">jetbra.zip</a>, and configure as described in
            <strong>readme.txt</strong>! For testing purposes only, not for
            commercial use! <br>
            <strong>Please note that this is just a personal page, not an official website!</strong>
        </p>
    </header>

    <main class="px-6 py-10 grid gap-(--gutter,1rem) grid-cols-[repeat(auto-fill,minmax(min(var(--space,10rem),100%),1fr))]"
          style="--space: 20rem; --gutter: 3.5rem">
        {#each productList as product}
            <article
                    class="group shadow-lg rounded-2xl transition-all duration-400 ease-in-out w-[90%] relative overflow-visible bg-(--card-bg) mx-auto hover:-translate-y-0.5"
                    data-sequence={product.Code}>
                <header>
                    <div class="flex items-center justify-between px-6 pt-(--spacing) pb-0 bg-(--card-bg) rounded-(--radius)">
                        <div class="relative w-(--size) h-(--size) text-[1.25rem] select-none translate-y-1/2 flex items-center justify-center overflow-hidden shrink-0">
                            <svg class="w-full h-full m-0 bg-card-bg text-transparent object-cover text-center text-indent-10000"
                                 role="img">
                                <use href={`#${product.productFamilyName}`}></use>
                            </svg>
                        </div>
                        <button data-version={product.version} class="cursor-pointer outline-none select-none inline-block items-center justify-between border border-transparent rounded-(--radius) box-border px-[21px] py-[12px] font-normal tracking-[1.2px] bg-transparent transition-[border,color] duration-250 ease-out max-w-[60%]
                     dark:text-gray-500 text-sm text-right relative
                     before:content-[attr(data-version)] before:whitespace-nowrap before:truncate before:w-full before:block hover:border-(--accent) hover:text-(--accent)">
                            <ul class="absolute top-full left-0 bg-(--main-bg) backdrop-blur-[18px] w-fit rounded-(--text-sm) shadow-[0_4px_12px_rgba(0,0,0,0.1)] text-left opacity-0 invisible transition ease-in-out duration-300 z-99
                            before:content-[''] before:absolute before:-top-[6px] before:left-5 before:w-0 before:h-0 before:shadow-[2px_-2px_6px_rgba(0,0,0,0.05)] before:border-t-[6px] before:border-t-(--main-bg) before:border-r-[6px] before:border-r-(--main-bg) before:border-b-[6px] before:border-b-transparent before:border-l-[6px] before:border-l-transparent before:-rotate-45 before:mix-blend-multiply">
                                <li class="active z-99 relative bg-transparent px-5 text-(--text-main) transition-colors duration-250 ease-out hover:bg-(--hover-color) first:rounded-t-(--text-sm) last:rounded-b-(--text-sm)">
                                    <a href="#"
                                       class="block border-b border-(--border-color) py-[16px] text-inherit no-underline whitespace-nowrap active:text-(--accent) last:border-b-0">{product.version}</a>
                                </li>
                            </ul>
                        </button>
                    </div>
                    <hr class="m-0 p-0 bg-(--border-color) h-[1px] border-none"/>
                </header>
                <div class="p-6 overflow-hidden bg-(--card-bg) pt-10 rounded-(--radius)">
                    <h1 class="line-clamp-1 text-(--text-main) mt-0 text-ellipsis font-bold text-[2em] my-[0.67em]"
                        title={product.name}>{product.name}</h1>
                    <p title="Click to copy full license text" class="
                   my-[1em] relative cursor-pointer transition-all duration-300 ease-in-out line-clamp-3 text-sm hover:text-transparent
                   dark:text-gray-500 after:content-[attr(data-content)] after:absolute after:text-transparent after:top-0 after:left-0 after:w-full after:h-full after:flex after:items-center after:justify-center after:rounded-[var(--radius)] after:transition-all after:duration-300 after:ease-in-out
                   hover:after:text-[var(--text-main)] hover:after:bg-[var(--hover-color)]"
                       onclick={(event) => {copyLicense(event, product.code)}}
                       data-content="Copy to clipboard">
                        *********************************************************************************************************************************************************
                    </p>
                </div>
                <div class="transition duration-200 absolute -z-10 w-[88%] h-full bottom-0 rounded-2xl bg-[var(--grey-600)] left-1/2 -translate-x-1/2 group-hover/card:bottom-[-1.5rem]"></div>
                <div class="transition duration-200 absolute -z-10 w-[88%] h-full bottom-0 rounded-2xl bg-[var(--grey-600)] left-1/2 -translate-x-1/2 group-hover/card:bottom-[-2.5rem]"></div>
            </article>
        {/each}
    </main>
    <footer class="pt-10 w-[96%] mt-10 mx-auto pb-10 border-t border-(--border-color) flex items-center justify-between">
        <div class="lt-panel">
            <span class="text-base dark:text-gray-500">All the above keys are collected from the Internet and are for testing purposes only, not for commercial use!</span>
        </div>
        <div class="text-sm dark:text-gray-500">Theme by QieTuZai</div>
    </footer>
</div>

<!-- svelte-ignore css_unused_selector -->
<style lang="postcss">
    :root {
        --text-grey: #9e9e9e;
        --text-main: rgba(0, 0, 0, 0.87);
        --spacing: 4px;
        --size: 64px;
        --radius: 1.5rem;
        --accent: #5380f7;
        --text-sm: 0.875rem;
        --main-bg: #fff;
        --card-bg: #fff;
        --hover-color: #eee;
        --border-color: rgba(0, 0, 0, 0.05);
        --grey-400: rgba(0, 0, 0, 0.04);
        --grey-600: rgba(0, 0, 0, 0.06);
    }

    @media (prefers-color-scheme: dark) {
        :root {
            --main-bg: rgb(0, 0, 0);
            --card-bg: rgb(31, 34, 38);
            --text-main: #d9d9d9;
            --text-grey: #6e767d;
            --accent: #1d9bf0;
            --hover-color: rgba(255, 255, 255, 0.07);
            --border-color: #4b4648;
        }
    }

    #container {
        font-size: 1rem;
        line-height: 1.5;
        word-wrap: break-word;
        font-kerning: normal;
        font-family: 'Gotham SSm A', 'Gotham SSm B', 'Arial Unicode MS', Helvetica, sans-serif;
        margin: 0;
        padding: 0;
        -webkit-font-smoothing: antialiased;
        background-color: var(--main-bg);
    }
</style>