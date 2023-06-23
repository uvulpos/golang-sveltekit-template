<script lang="ts">
  export let imgUrl: string;
  export let imgAlt: string;
  export let name: string;

  let imageObject: HTMLImageElement;
  let alternativeTextObject: HTMLDivElement;

  $: displayInitials = calculateInitialsFromName(name);
  function calculateInitialsFromName(name: string): string {
    return name
      .split(" ")
      .map((n) => n[0])
      .join("")
      .substring(0, 2);
  }

  function handleImageNotFound(e: Event) {
    imageObject.style.display = "none";
    alternativeTextObject.style.display = "block";
  }

  function handleImageFound(e: Event) {
    imageObject.style.display = "block";
    alternativeTextObject.style.display = "none";
  }
</script>

<div class="content">
  <img
    src={imgUrl}
    alt={imgAlt}
    bind:this={imageObject}
    on:load={handleImageFound}
    on:error={handleImageNotFound}
  />
  <div class="alternative-initials" bind:this={alternativeTextObject}>
    <p>{displayInitials}</p>
  </div>
</div>
