<script>
    import User from "../components/user.svelte"

    const dataUrl = "https://github.com/kim-ahri/mercenary-server/blob/master/public/userdata.json"

    async function getData() {
        const r = await fetch(dataUrl)
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