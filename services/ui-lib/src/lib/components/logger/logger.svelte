<script lang="ts">
  import { Icon } from "../icon";

  export let showTime: boolean = true;
  export let showEmojis: boolean = false;

  const data = [
    {
      time: new Date(),
      icon: "home",
      message: "Hello World",
    },
    {
      time: new Date(),
      icon: "home",
      message: "Hello World",
    },
    {
      time: new Date(),
      icon: "home",
      message: "Hello World",
    },
    {
      time: new Date(),
      icon: "home",
      message:
        "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consectetur distinctio aliquid accusamus deserunt aspernatur sed. Laboriosam earum, aspernatur asperiores quam ut impedit fugit neque dicta, laborum qui soluta iusto ipsum.",
    },
  ];
</script>

<div class="logger">
  <div class="header">⤵ Execute some stuff</div>
  <ul>
    {#each data as row, i}
      <li class="row">
        <span class="identifier">{i}</span>
        {#if showTime}
          <span class="date">{row?.time?.toLocaleTimeString()}</span>
        {/if}
        {#if showEmojis}
          <span>
            <Icon icon={row?.icon} --size="1rem"></Icon>
          </span>
        {/if}
        <span class="value">{row?.message}</span>
      </li>
    {/each}
  </ul>
</div>

<style lang="sass">
    .logger
        background-color: #191919
        color: #fff
        padding: 1rem
        height: 200px
        overflow-y: scroll
        resize: vertical
        .header
            background-color: #555
            padding: .5rem 1rem
            border-radius: 10px
            margin-bottom: 1rem
        ul
            list-style-type: none
            padding: 0
            margin: 0

            li.row
                display: flex
                gap: 1rem
                &:hover
                    background-color: lighten(#191919, 10%)
                .identifier, .date, .emoji
                    display: block
                    color: darken(#fff, 25%)
                    overflow: hidden
                    width: 8ch
                .identifier
                    width: 6ch
                    text-align: right
                .date
                    text-align: center
                .emoji
                    text-align: center

</style>
