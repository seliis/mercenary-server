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
            <User name={v.name}
                  balance={v.balance}
                  aircraft={v.record.aircraft}
                  air={v.record.air}
                  ground={v.record.ground}
                  sortie={v.record.sortie}
                  killed={v.record.killed}
            />
        {/each}
    {/await}
</div>