<script lang="ts">
  import {
    Navbar,
    NavBrand,
    NavLi,
    NavUl,
    NavHamburger,
  } from "flowbite-svelte";

  import { Darkmode } from "$lib/components/darkmode";
</script>

<Navbar let:hidden let:toggle>
  <NavBrand href="/">
    <img
      src="/assets/gifs/gopher-dance.gif"
      class="mr-3 h-6 sm:h-9"
      alt="Dancing Gopher as a Logo"
    />
    <span
      class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
    >
      Go Svelte!
    </span>
  </NavBrand>
  <NavHamburger on:click={toggle} />
  <NavUl {hidden}>
    <NavLi href="/">Home</NavLi>
    <NavLi href="/github">GitHub</NavLi>
    <NavLi><Darkmode /></NavLi>
  </NavUl>
</Navbar>
