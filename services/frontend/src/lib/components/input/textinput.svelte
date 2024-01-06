<script lang="ts">
  import { toast } from "$lib/stores/toast";
  import { debounce } from "throttle-debounce";
  import type { HTMLInputTypeAttribute } from "svelte/elements";
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  export let labelName: string | undefined = undefined;
  export let showDefaultMargin: boolean = true;

  export let name: string = "";
  export let autocomplete: string | null | undefined = undefined;
  export let type: HTMLInputTypeAttribute | null | undefined = "text";
  export let placeholder: string = "";
  export let value: string = "";
  export let disabled: boolean = false;
  export let validateFunction:
    | undefined
    | ((input: string) => Promise<boolean | null>) = undefined;

  type validStatus = "empty" | "invalid" | "valid" | "checking" | "unknown";

  let isValid: validStatus = "empty";

  const debounceFunc = debounce(2000, (textValue: string) => {
    runValidation(textValue);
  });

  async function onInputHandler() {
    dispatch("inputchange", { value: value });

    if (validateFunction === undefined) return;
    isValid = "checking";
    debounceFunc(value);
  }

  async function runValidation(textValue: string) {
    if (validateFunction === undefined) return;
    const result = await validateFunction(textValue);
    if (result === null) {
      toast.push(
        "error",
        "Internal Server Error",
        `Could not validate ${name}`
      );
      isValid = "unknown";
      toast.push("error", "Internal Server Error", "could not validate input");
      return;
    }

    if (result === true) {
      isValid = "valid";
      return;
    }

    isValid = "invalid";
  }
</script>

<label class:showDefultMargin={showDefaultMargin}>
  {#if labelName !== undefined}
    <span>{labelName}</span>
  {/if}

  <div class="input">
    {#if type === "password"}
      <input
        {name}
        {autocomplete}
        type="password"
        {placeholder}
        bind:value
        {disabled}
        on:input={onInputHandler}
      />
    {:else}
      <input
        {name}
        {autocomplete}
        type="text"
        {placeholder}
        bind:value
        {disabled}
        on:input={onInputHandler}
      />
    {/if}
    {#if isValid !== "empty"}
      <img
        src={`/assets/vector/status-${isValid}.svg`}
        class="validation-icon-{isValid}"
        alt="validation icon"
      />
    {/if}
  </div>
</label>

<style lang="sass">
    @import "../../variables/sass/main"

    label
      display: flex
      flex-direction: column
      gap: .5rem
      &.showDefultMargin
        margin: 1rem
      
      .input
        display: flex
        align-items: center
        justify-content: space-between
        padding: .5rem 1rem
        border-radius: 5px
        font-size: 1rem
        background-color: $ui-inputs-textinputs-background
        input
            unset: all
            flex-grow: 1
            border: none
            outline: none
            background: none
            color: $ui-font-color
            &:disabled
              cursor: not-allowed
              opacity: .3

        img.validation-icon-checking
          animation: rotateSpinner
          animation-duration: 1.5s
          animation-iteration-count: infinite


    @keyframes rotateSpinner 
      0%   
        rotate: 0deg
      100% 
        rotate: 360deg


</style>
