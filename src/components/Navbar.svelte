<script>
  import Avatar from "./Avatar"

  export let username
  export let avi
  export let onClickLogin
</script>

<style>
  .navbar {
    width: 100%;
    display: flex;
    background-color: #f1f6fa;
    flex-direction: row;
    padding: 1rem;
    position: fixed;
    height: 80px;
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

  .avatar-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .avatar-wrapper > h2 {
    color: #1da1f2;
    margin-right: 1rem;
  }

  @media (max-width: 768px) {
    .avatar-wrapper > h2 {
      display: none;
    }
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
    <div class="avatar-wrapper">
      <h2>Welcome, {username}!</h2>
      <Avatar url={avi} alt={username} />
    </div>
  {:else}
    <button class="login-btn" on:click={onClickLogin}>Login</button>
  {/if}
</nav>
