<script>
  export let username
  export let avi
</script>

<style>
  .navbar {
    width: 100%;
    display: flex;
    background-color: #f1f6fa;
    flex-direction: row;
  }

  .filler {
    flex: 3;
  }

  h1 {
    font-weight: bold;
  }
</style>

<nav class="navbar">
  <h1>tweetop</h1>
  <div class="filler" />
  {#if username && avi}
    <img src={avi} alt={username} />
  {:else}
    <button>Login</button>
  {/if}
</nav>
