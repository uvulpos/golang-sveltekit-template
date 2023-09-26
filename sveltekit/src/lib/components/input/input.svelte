<script lang="ts">
  import { Label, Input } from "flowbite-svelte";

  export let label: string;
  export let placeholder: string = "";
  export let type: "text" | "email" = "text";
  export let value: string = "";
</script>

<Label class="space-y-2">
  <span>{label}</span>
  <Input bind:value {type} {placeholder} size="md" />
</Label>
