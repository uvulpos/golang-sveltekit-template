<script>
  import { _ } from "svelte-i18n";
  import Avatar from "../avatar/avatar.svelte";
</script>

<header>
  <div class="logo">
    <img src="/assets/gifs/gopher-dance.gif" alt="Dancing Gopher as a Logo" />
    <span>Go Svelte! ❤️</span>
  </div>
  <nav>
    <ul>
      <li><a href="/">{$_("header.navigation.home")}</a></li>
      <li>
        <a href="/github">{$_("header.navigation.github")}</a>
      </li>
    </ul>
  </nav>
  <div class="account">
    <div class="not-login">
      <a href="/login">
        <button>Login</button>
      </a>
    </div>
    <div class="authorized">
      <Avatar
        imgAlt="user profil picture"
        imgUrl="/assets/gifs/gophesr-dance.gif"
        name="Tim Riedl"
      />
    </div>
  </div>
</header>

<style lang="scss">
  header {
    $header-padding: 10rem;
    width: calc(100% - ($header-padding * 2));
    padding: 0 10rem;
    background-color: rgb(37, 37, 41);
    z-index: 9999;
    display: flex;
    justify-content: space-between;
    align-items: center;

    div.logo {
      display: flex;
      align-items: center;
      gap: 1rem;
      color: #96d6ff;
      font-weight: bolder;
      img {
        height: 3rem;
      }
    }

    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 2rem;
    nav ul {
      padding: 0;
      margin: 0;
      list-style-type: none;
      display: flex;
      flex-direction: row;
      align-items: center;

      li a {
        display: block;
        text-decoration: none;
        color: rgba(255, 255, 255, 0.551);
        transition: color 200ms ease-in-out;
        font-size: 1rem;
        padding: 1rem;
        &:hover {
          color: white;
        }
      }
    }

    div.account {
      display: flex;
      flex-direction: row;
      align-items: center;
      div.not-login {
        a button {
          cursor: pointer;
          background-color: rgba(255, 102, 0, 0.05);
          border: 1px solid #f60;
          color: #f60;
          border-radius: 5px;
          padding: 0.5rem 2rem;
          font-size: 1rem;
          margin: 0;
          transition: background-color 200ms ease-in-out,
            color 200ms ease-in-out;

          &:hover {
            background-color: #f60;
            color: white;
          }
        }
      }
    }
  }
</style>
