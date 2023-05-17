<script lang="ts">
  import Button from "../../lib/components/button/Button.svelte";

  let buttonDisabled: boolean = false;
  let catImageURL: string = "https://cataas.com/cat/gif";

  function loadCatImage() {
    buttonDisabled = true;
    catImageURL = `https://cataas.com/cat/gif?myTime=${+new Date()}`;
  }

  function enableButton() {
    buttonDisabled = false;
  }
</script>

<div class="content">
  <div class="image">
    <!-- svelte-ignore a11y-img-redundant-alt -->
    <img alt="cat gif" src={catImageURL} on:load={enableButton} />
  </div>
  <div class="control">
    <Button bind:disabled={buttonDisabled} on:click={loadCatImage}
      >New Cat</Button
    >
  </div>
</div>

<style lang="scss">
  .content {
    max-width: 1200px;
    margin: 10vh auto 0;

    .image {
      height: 50vh;
      text-align: center;
      background-color: #ebebeb;
      padding: 1rem;
      img {
        height: 100%;
        object-fit: contain;
      }
    }

    .control {
      display: flex;
      justify-content: center;
      gap: 0.5rem;
      margin-top: 2rem;
    }
  }
</style>
