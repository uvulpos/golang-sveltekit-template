<script lang="ts">
  import { goto } from "$app/navigation";
  import type { SelfInformation } from "$lib/api/user/models/SelfInformation";
  import { logoutSession } from "$lib/functions/logout/logout";
  import { Exit } from "radix-icons-svelte";
  import { _ } from "svelte-i18n";

  export let user: SelfInformation;
</script>

<div class="content-element sidebar-account-content-element">
  <div class="account">
    <img src={user.profile_picture} alt="" class="profilepicture" />
    <div class="account-name" title="{user.display_name} ({user.username})">
      <span class="displayname">{user.display_name} </span>
      <span class="username">
        ({user.username})
      </span>
    </div>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div class="margin-left">
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <!-- svelte-ignore a11y-missing-attribute -->
      <a
        on:click={() => {
          logoutSession();
        }}
      >
        <div class="logout-button">
          <Exit />
        </div>
      </a>
    </div>
  </div>
</div>

<style lang="sass">
  .content-element
    .account
      // border: 1px solid #000
      border-radius: 5px
      display: flex
      align-items: center
      gap: .5rem
      padding: .25rem .5rem

      .profilepicture
        $size: 30px
        border-radius: 100%
        display: block
        height: $size
        width: $size
        background-color: #f60
        object-fit: cover

      .account-name
        display: flex
        flex-direction: column
        gap: .1rem

        .displayname
          font-size: 1rem
        .username
          font-size: .8rem
          color: var(--sidebar-font-color-secondary)

      .margin-left
        margin-left: auto

      .logout-button
        :global(svg)
          height: 1.2rem
          width: 1.2rem
          object-fit: contain
          color: var(--sidebar-font-color-red)
          transition: 200ms color ease-in-out
          &:hover
            color: var(--sidebar-font-color-red-hover)
        :global(svg path)
          transition: 200ms stroke ease-in-out
          stroke: .2px var(--sidebar-font-color-red)
        :global(svg:hover path)
          stroke-color: .2px var(--sidebar-font-color-red-hover)

      :global(.menu button:hover)
        background-color: var(--sidebar-background-color-hover)
      :global(.menu button svg)
        color: var(--sidebar-svg-color)

</style>
