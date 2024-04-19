<script lang="ts">
  import type { AuthorInformations } from ".";
  import { Icon } from "../icon";

  export let roundCorner: boolean = false;
  export let authorData: AuthorInformations = [];
</script>

<div class="author" class:round-corner={roundCorner}>
  <ul>
    {#each authorData as information}
      {@const showPicture =
        information?.imagePath !== null && information?.imagePath !== undefined}
      {@const showIcon =
        information?.iconName !== null &&
        information?.iconName !== undefined &&
        !showPicture}
      <li>
        <div class="information">
          <span>{information?.header}:</span>
          <span>{information?.value}</span>
        </div>
        {#if showPicture}
          <!-- svelte-ignore a11y-img-redundant-alt -->
          <!-- svelte-ignore a11y-missing-attribute -->
          <img src={information?.imagePath} />
        {/if}
        {#if showIcon}
          <Icon icon={information?.iconName ?? ""} />
        {/if}
      </li>
    {/each}
  </ul>
</div>

<style lang="sass">
    .author
        padding: 1rem 5%
        background-color: #191919
        color: #ffffff
        &.round-corner
            border-radius: 10px
            overflow: hidden
        ul
            padding: 0
            margin: 0
            list-style-type: none
            display: flex
            gap: 5rem
            li
                display: flex
                flex-direction: row-reverse
                gap: 1rem
                align-items: center
                img
                    height: calc(1rem + .8rem + .25rem)
                    aspect-ratio: 1/1
                    object-fit: cover
                    border-radius: 100%
                .information
                    display: flex
                    flex-direction: column
                    gap: .25rem
                    &>:first-child
                        font-weight: 200
                        font-size: .8rem
                        color: darken(#ffffff, 25%)
</style>
