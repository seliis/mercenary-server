<script>
    import Router from "svelte-spa-router"

    // Pages
    import Home from "./routes/home.svelte"

    // Routes
    const routes = {
        "/": Home
    }
</script>

<main onmousedown="return false">
    <header>
        <h1>Ssouchan's Mercenary Server</h1>
    </header>
    <section>
        <Router {routes}/>
    </section>
</main>

<style>
    :global(*) {
        margin: 0;
        padding: 0;
    }

    :global(body) {
        font-family: "Roboto", sans-serif;
        background-color: #151E27;
        color: white;
        cursor: default;
    }

    main {
        grid-template-rows: 100px 1fr;
        grid-template-columns: 1fr;
        position: relative;
        overflow: hidden;
        display: grid;
        height: 100vh;
        width: 100vw;
    }

    header {
        background-color: #0A0E12;
        grid-area: 1 / 1 / 2 / 2;
        justify-content: center;
        align-items: center;
        display: flex;
    }

    section {
        grid-area: 2 / 1 / 3 / 2;
        overflow-y: scroll;
    }

    h1 {
        font-weight: 700;
        font-size: 2.5rem;
    }

    @media only screen and (max-width: 1600px) {
        h1 {
            font-size: 1rem;
        }
    }

    @import url('https://fonts.googleapis.com/css2?family=Roboto:wght@700&display=swap');
</style>