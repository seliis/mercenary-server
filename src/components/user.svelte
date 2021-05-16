<script>
    import { slide } from "svelte/transition"
    import { quintOut } from "svelte/easing"

    export let name, balance, aircraft, air, ground, sortie, killed
</script>

<div id="user" transition:slide="{{delay: 250, duration: 1000, easing: quintOut}}">
    <ul>
        <li id="name">
            {name}
        </li>
        <li style="color: #FFD700;">
            ${balance.toLocaleString("ko-KR")}
        </li>
        <li>
            {aircraft}
        </li>
        <li>
            {sortie}
        </li>
        <li>
            {air}
        </li>
        <li>
            {ground}
        </li>
        <li style="color: #FF0000;">
            {killed}
        </li>
    </ul>
</div>

<style>
    #user {
        font-family: "Ubuntu Mono", monospace;
        background-color: #0A0E12;
        position: relative;
        font-size: 1.25rem;
        height: 50px;
        margin: 5px;
    }

    ul {
        display: grid;
        grid-template-columns: repeat(7, 1fr);
        grid-template-rows: 100%;
        position: relative;
        width: 100%;
        height: 100%;
    }

    li {
        list-style: none;
        margin: auto;
    }

    #name {
        background: linear-gradient(135deg, #C33864 0%, #822A99 50%, #1E2670 100%);
        height: 100%;
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        font-family: "Noto Sans KR", sans-serif;
        font-weight: 700;
        font-size: 1.25rem;
    }
</style>