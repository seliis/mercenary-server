<script>
    import User from "../components/user.svelte"

    async function getData() {
        const r = await fetch("/userdata.json")
        return await r.json()
    }
</script>

<div id="home">
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