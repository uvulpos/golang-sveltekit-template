<script>
    import { _ } from "svelte-i18n";
    import { GithubLogo } from "radix-icons-svelte";

    $: header = $_("page.legal.licenses.header-software-licenses")
    $: header_participating_developers = $_("page.legal.licenses.header-participating-developers")
    $: text_participating_developers = $_("page.legal.licenses.text-participating-developers")

</script>

# Software Lizenzen

## Partizipierende Developer

Hier dürfen sich alle Entwickler verewigen, welche an diesem Projekt gearbeitet haben:

- Tim Riedl [<GithubLogo />](github.com/uvulpos)
