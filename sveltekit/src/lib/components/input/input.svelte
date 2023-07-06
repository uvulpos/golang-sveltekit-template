<script lang="ts">
  import { Label, Input } from "flowbite-svelte";

  export let label: string;
  export let placeholder: string = "";
  export let type: "text" | "email" = "text";
</script>

<Label class="space-y-2">
  <span>{label}</span>
  <Input {type} {placeholder} size="md" />
</Label>
