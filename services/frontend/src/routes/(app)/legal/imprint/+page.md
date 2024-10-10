 <script>
    import { _ } from "svelte-i18n";

    $: header = $_("page.legal.imprint.header")
    $: linefirst = $_("page.legal.imprint.this-website-template")
    $: linesecond = $_("page.legal.imprint.using-template-liable")
    $: github = $_("page.legal.imprint.project-on-github")

</script>

# {header}

{linefirst}

{linesecond}

[{github}](https://github.com/uvulpos/golang-sveltekit-template)
