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
    padding: 1rem;
  }

  .filler {
    flex: 2;
  }

  .login-btn {
    border: none;
    background-color: #1da1f2;
    color: white;
    font-size: 1rem;
    padding: 0.5rem;
    padding-left: 2rem;
    padding-right: 2rem;
    border-radius: 1rem;
    text-align: center;
  }

  .login-btn:focus {
    border: none;
    outline: 0;
  }

  .logo {
    width: 2.5rem;
    height: 2.5rem;
    margin-right: 1rem;
  }

  .logo-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  h1 {
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
    font-weight: 900;
  }
</style>

<nav class="navbar">
  <div class="logo-wrapper">
    <img class="logo" src="assets/Twitter_Logo_Blue.png" alt="twitter-logo" />
    <h1>tweetop</h1>
  </div>
  <div class="filler" />
  {#if username && avi}
    <img src={avi} alt={username} />
  {:else}
    <button class="login-btn">Login</button>
  {/if}
</nav>
