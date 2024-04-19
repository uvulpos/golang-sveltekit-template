<script lang="ts">
  import { Icon, IconType } from "../icon";

  export let icon: IconType | string;
  export let animateButton: boolean = true;

  /**
   * Button contents
   */
  export let label: string = "";
</script>

<button class="btn information-button" on:click>
  <div class="icon">
    <!-- <Icon --uui-icon-size="1.5rem" type={icon} /> -->
    <Icon {icon} />
  </div>
  <div class="content" class:animate={animateButton}>
    <div class="label">
      {label}
    </div>
  </div>
</button>

<style lang="sass">
    .btn
        $extra-space: 25px
        border: none 
        border-radius: 10px
        overflow: hidden
        cursor: pointer
  
        color: #fff
        background-color: #58A1E3
        box-shadow: 2px 2px 3px rgba(#000, .1)
        font-size: 1em
        height: 2.4em
        padding: 0
  
        display: flex
        flex-direction: row
        align-items: center
  
        &:hover,&:active,&:focus
          border: none
          outline: none
          .content.animate::before
              width: calc(100% + $extra-space)
  
        .icon 
          background-color: darken(#58A1E3, 25%)
          border-radius: 10px 0 $extra-space 10px
          z-index: 10
          padding: .2rem
          height: 2.4rem
          width: 2.4rem
          display: flex
          align-items: center
          justify-content: center
        
        .content
            padding: 0 10px
            height: 100%
            position: relative
            display: flex
            align-items: center
          
  
            .label
              z-index: 20
            &.animate::before
              content: ""
              height: 100%
              width: 0
              background-color: darken(#58A1E3, 25%)
              position: absolute
              top: 0
              left: (0-$extra-space)
              animation: fade-out 0.5s ease-in forwards
              transition: 0.5s all

</style>
