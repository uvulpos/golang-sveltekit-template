<script>
  import { _ } from "svelte-i18n";
</script>

<div class="content">
  <span>Go SvelteKit! ❤️</span>
  <!-- <img alt="dancing gopher" src="/assets/gifs/gopher-dance.gif"> -->
  <ul>
    <li><a href="/">{$_("header.navigation.home")}</a></li>
    <li><a href="/cat">{$_("header.navigation.cat-on-demand")}</a></li>
    <li>
      <a
        href="https://github.com/uvulpos/golang-sveltekit-binary"
        target="_blank">{$_("header.navigation.github")}</a
      >
    </li>
  </ul>
</div>

<style lang="scss">
  .content {
    background-color: rgb(0, 0, 0);
    color: white;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    align-items: center;
    max-width: 100vw;
    padding: 0rem 2rem;

    span {
      display: block;
    }

    ul {
      margin: 0 0 0 0;
      margin-left: auto;
      display: flex;
      flex-direction: row;
      // gap: .5rem;
      list-style-type: none;

      li {
        a {
          display: block;
          text-decoration: none;
          color: #ff3e00;
          padding: 1rem 1rem;
          font-weight: 500;
          transition: background-color 200ms ease-in-out;

          &:hover {
            background-color: #272727;
          }
        }
      }
    }
  }
</style>
