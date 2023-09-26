<script>
  import { _ } from "svelte-i18n";
  import { Heading, P, Button } from "flowbite-svelte";
</script>

<div class="text-center mx-4">
  <img
    alt="Svelte Society Gopher"
    src="/assets/images/svelte-society-gopher.png"
    class="w-4/6 lg:w-1/4 m-auto my-4"
  />
  <Heading
    tag="h1"
    class="mb-4"
    style="color: #f60"
    customSize="text-3xl font-extrabold  md:text-4xl lg:text-5xl"
    >{$_("page.home.header-1")}</Heading
  >
  <P class="mb-6 text-sm lg:text-md sm:px-12 xl:px-48 dark:text-gray-400"
    >{$_("page.home.text")}</P
  >
  <Button href="/github"
    >Visit on GitHub
    <svg
      class="ml-2 -mr-1 w-5 h-5"
      fill="currentColor"
      viewBox="0 0 20 20"
      xmlns="http://www.w3.org/2000/svg"
      ><path
        fill-rule="evenodd"
        d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z"
        clip-rule="evenodd"
      /></svg
    >
  </Button>
</div>
