<script lang="ts">
  import { _ } from "svelte-i18n";
</script>

<div class="footer">
  <div class="socialmedia">
    <a
      href="https://github.com/uvulpos/golang-sveltekit-binary"
      target="_blank"
      title="GitHub"
    >
      <img alt="GitHub logo" src="/assets/logos/github.svg" />
    </a>
    <a
      href="https://www.linkedin.com/in/riedl-tim/"
      target="_blank"
      title="Linkedin"
    >
      <img alt="Linkedin logo" src="/assets/logos/linkedin.svg" />
    </a>
  </div>
  <p class="note">© 2023 uVulpos {$_("page.footer.allrightsreserved")}</p>
</div>

<style lang="sass">
  .footer
    padding: 1.5rem
    background-color: rgba(0,0,0,.3)
    display: flex
    flex-direction: column
    gap: .5rem

    .socialmedia
      display: grid
      // grid-auto-flow: column
      gap: 2rem
      justify-content: center
      align-content: center

      grid-auto-flow: column
      // grid-template-columns: repeat(auto-fill, minmax(3rem, 1fr))
      a
        width: 1.5rem
        transition: opacity 200ms ease-in-out
        img
          width: 1.5rem
          height: 1.5rem
          filter: invert(1)
          opacity: .8
        &:hover
          opacity: .5
    
    .note
      text-align: center
      margin: 0


</style>
