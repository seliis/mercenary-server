<script>
    import User from "../components/user.svelte"

    async function getData() {
        const r = await fetch("/userdata.json")
        return await r.json()
    }
</script>

<div id="home">
    <div id="field">
        <ul>
            <li>Player</li>
            <li>Balance</li>
            <li>Aircraft</li>
            <li>Sortie</li>
            <li>Air</li>
            <li>Ground</li>
            <li>Killed</li>
        </ul>
    </div>
    {#await getData() then a}
        {#each a as v}
            <User name={v.Name}
                  balance={v.Balance}
                  aircraft={v.Aircraft}
                  air={v.Air}
                  ground={v.Ground}
                  sortie={v.Sortie}
                  killed={v.Killed}
            />
        {/each}
    {/await}
</div>

<style>
    #field {
        font-family: "Ubuntu Mono", monospace;
        background-color: #0A0E12;
        text-transform: uppercase;
        letter-spacing: 1px;
        font-size: 1.25rem;
        position: relative;
        height: 50px;
        margin: 5px;
        width: 100%;
    }

    ul {
        display: grid;
        grid-template-columns: repeat(7, 1fr);
        grid-template-rows: 100%;
        width: 100%;
        height: 100%;
    }

    li {
        list-style: none;
        margin: auto;
    }

    @media only screen and (max-width: 1600px) {
        #field {
            letter-spacing: 3px;
            font-size: 1rem;
            width: 1500px;
            height: 50px;
        }
    }
</style>