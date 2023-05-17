<script lang="ts">
  import { _ } from "svelte-i18n";
import {pokemonData} from "./data.js";

  let pokemonApi: string = $_("api.endpoints.pokemon-all");

</script>

<div class="content">
    <div class="pokemons">
        {#each pokemonData as pokemon}
            <div class="pokemon">
                <img alt="Pokemon {pokemon?.name}" loading="lazy" src={pokemon?.ThumbnailImage}>
                <span>{pokemon?.name}</span>
            </div>
        {/each}
    </div>
</div>

<style lang="scss">
    .content {
        max-width: 1200px;
        margin: auto;

        .pokemons {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(calc(160px + 2rem), 1fr));
            gap: .5rem;
            margin-top: 6rem;

            .pokemon {
                background-color: #f0f0f0;
                display: flex;
                align-items: center;
                flex-direction: column;
                padding: 1rem;
                border-radius: 5px;
                img {
                    width: 160px;
                    height: 160px;
                    object-fit: contain;
                }
            }
        }
    }
</style>