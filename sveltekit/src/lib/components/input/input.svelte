<script lang="ts">
  import { InputTypes } from "./input-types";

  export let label: string;
  export let name: string = "";
  export let inputType: InputTypes = InputTypes.Text;
  export let placeholder: string = "";

  export let textinput: string = "";

  $: placeholderWfallback = placeholder !== "" ? placeholder : label;
  $: currentInputType = inputType;

  function togglePasswordVisibility() {
    if (currentInputType === InputTypes.Password) {
      currentInputType = InputTypes.Text;
    } else {
      currentInputType = InputTypes.Password;
    }
  }
</script>

<label>
  <span>{label}</span>
  <div class="input">
    {#if inputType === InputTypes.Password}
      <div class="show-password" title="show / hide current password">
        {#if currentInputType === InputTypes.Password}
          <img
            src="/assets/vector/eye.svg"
            alt="Show Password Eye"
            on:click={togglePasswordVisibility}
            on:keypress
          />
        {:else}
          <img
            src="/assets/vector/no-eye.svg"
            alt="Show Password Eye"
            on:click={togglePasswordVisibility}
            on:keypress
          />
        {/if}
      </div>
    {/if}
    <input
      placeholder={placeholderWfallback}
      type={currentInputType}
      {name}
      value={textinput}
      on:change
      on:input
    />
  </div>
</label>

<style lang="scss">
  label {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    .input {
      position: relative;
      input {
        outline: 0;
        border: none;
        background-color: rgb(52, 52, 56);
        color: white;
        padding: 0.5rem 1rem;
        border-radius: 5px;
        width: 40ch;
        &[type="password"] {
          -webkit-text-security: square;
        }
      }
      .show-password {
        position: absolute;
        z-index: 110;
        transform: translate(-50%, -50%);
        top: 55%;
        right: 0.3rem;

        img {
          height: 1rem;
          width: 1rem;
          cursor: pointer;
          filter: invert(1);
          opacity: 0.4;
        }
      }
    }
  }
</style>
