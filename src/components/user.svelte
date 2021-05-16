<script>
    import { slide } from "svelte/transition"
    import { quintOut } from "svelte/easing"

    export let name, balance, aircraft, air, ground, sortie, killed
</script>

<div id="user" transition:slide="{{delay: 250, duration: 1000, easing: quintOut}}">
    <div id="container">
        <div id="base">
            <div id="name">{name}</div>
            <div id="balance">${balance.toLocaleString("ko-KR")}</div>
        </div>
        <div id="desc">
            <div id="primary">
                <div id="aircraft">
                    {aircraft}
                </div>
                <div id="sortie">
                    {sortie} TIMES ACTION
                </div>
            </div>
            <div id="secondary">
                <div id="air" class="sec-data">
                    AIR: {air}
                </div>
                <div id="ground" class="sec-data">
                    GND: {ground}
                </div>
                <div id="killed" class="sec-data">
                    KIA: {killed}
                </div>
            </div>
        </div>
    </div>
</div>

<style>
    #user {
        background-color: #0A0E12;
        justify-content: center;
        border-radius: 10px;
        align-items: center;
        position: relative;
        display: flex;
        height: 200px;
        margin: 20px;
    }

    #container {
        width: 100%;
        height: 100%;
        position: relative;
        display: grid;
        grid-template-columns: 40% 60%;
        grid-template-rows: 100%;
    }

    #base {
        display: grid;
        grid-template-columns: 100%;
        grid-template-rows: 50% 50%;
        background: linear-gradient(135deg, #C33864 0%, #822A99 50%, #1E2670 100%);
        border-top-left-radius: 10px;
        border-bottom-left-radius: 10px;
        text-shadow: 1px 3px 3px black;
    }

    #desc {
        display: grid;
        grid-template-columns: 75% 25%;
        grid-template-rows: 100%;
        font-family: "Ubuntu Mono", monospace;
    }

    #name {
        font-size: 2.5rem;
        font-weight: 700;
        margin: auto 0;
        padding: 0 50px;
    }

    #balance {
        font-size: 4rem;
        font-weight: 700;
        color: #FFD700;
        margin: auto 0;
        text-align: right;
        padding: 0 50px;
    }

    #aircraft {
        font-size: 3.5rem;
        font-weight: 700;
        text-align: center;
        padding: 1rem;
    }

    #sortie {
        font-size: 1.75rem;
        font-weight: 700;
        letter-spacing: 1px;
        color: #ADFF2F;
    }

    #killed {
        color: #FF0000;
    }

    #primary {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

    #secondary {
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        justify-content: space-around;
        font-size: 2rem;
        text-transform: uppercase;
        letter-spacing: 2px;
        font-weight: 700;
    }

    @media only screen and (max-width: 1600px) {
        #container {
            grid-template-columns: 100%;
            grid-template-rows: 25% 75%;
        }

        #base {
            display: flex;
            align-items: center;
            justify-content: space-between;
            padding: 0 30px;
            background: linear-gradient(135deg, #C33864 0%, #822A99 50%, #1E2670 100%);
            border-top-right-radius: 10px;
            border-bottom-left-radius: 0;
        }

        #desc {
            display: flex;
            align-items: center;
            justify-content: space-between;
            padding: 0 30px;
        }

        #name {
            font-size: 1rem;
            font-weight: 700;
            margin: 0;
            padding: 0;
        }

        #balance {
            font-size: 1.25rem;
            font-weight: 700;
            color: #FFD700;
            margin: 0;
            text-align: unset;
            padding: 0;
        }

        #aircraft {
            font-size: 1.5rem;
            padding: 5px 0;
        }

        #sortie {
            font-size: 1rem;
            letter-spacing: 0;
        }

        #primary {
            align-items: flex-start;
        }

        #secondary {
            font-size: 1rem;
            font-weight: 700;
        }

        .sec-data {
            padding: 5px 0;
        }
    }
</style>