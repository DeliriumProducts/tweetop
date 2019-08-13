<script>
  import Avatar from "./Avatar"
  import { RepeatIcon } from "svelte-feather-icons"

  export let username
  export let handle
  export let avi
  export let rt
  export let content
</script>

<style>
  .container {
    background: #f1f6fa;
    border-radius: 1.6rem;
    display: flex;
    flex-direction: column;
    padding: 1.5rem;
  }

  .header {
    display: flex;
    flex-direction: row;
  }

  .filler {
    flex: 2;
  }

  .user-info {
    margin-left: 1rem;
  }

  .user-handle {
    color: #1da1f2;
    text-align: left;
  }

  .user-name {
    text-align: left;
  }

  .rt {
    color: #17bf63;
    font-weight: 400;
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 1.6rem;
  }

  .rt-icon-wrapper {
    margin-left: 0.5rem;
    width: 1.6rem;
    height: 1.6rem;
  }

  .content {
    margin-top: 1rem;
    text-align: left;
    font-weight: 400;
  }
</style>

<div class="container">
  <div class="header">
    <Avatar url={avi} alt={username} />
    <div class="user-info">
      <h2 class="user-name">{username}</h2>
      <h3 class="user-handle">{handle}</h3>
    </div>
    <div class="filler " />
    <h2 class="rt">
      {rt}
      <div class="rt-icon-wrapper">
        <RepeatIcon />
      </div>
    </h2>
  </div>
  <div class="content">{content}</div>
</div>
