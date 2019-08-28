<script>
  import Navbar from "./components/Navbar"
  import Tweet from "./components/Tweet"

  let name = "Delirium Products!"
  let avi =
    "https://pbs.twimg.com/profile_images/1157472456505999363/T8w9d7hq.jpg"
  let loggedIn = false

  let tweets = [
    {
      username: "Dan Abramov",
      handle: "@dan_abramov",
      avi:
        "https://pbs.twimg.com/profile_images/1166344766210150401/amRnWzl-.jpg",
      rt: 99999,
      content: "react hooks",
    },
    {
      username: "Delirium Products!",
      handle: "@deliprods",
      avi:
        "https://pbs.twimg.com/profile_images/1157472456505999363/T8w9d7hq.jpg",
      rt: 50,
      content: "We are pleased to announce our latest project - LuncherBox!",
    },
    {
      username: "IMPZERO",
      handle: "@impzero",
      avi:
        "https://pbs.twimg.com/profile_images/1092758263744118785/-3MA9koq.jpg",
      rt: 20,
      content: "i am not do the cheating",
    },
    {
      username: "fr3fou!! 🎀",
      handle: "@fr3fou",
      avi:
        "https://pbs.twimg.com/profile_images/1064209784944189440/QtTtTJuT.jpg",
      rt: 2,
      content: "i am the retarded",
    },
  ]

  function onClickLogin() {
    loggedIn = true
  }
</script>

<style>
  .container {
    min-height: 100%;
    background: #e5edf5;
  }

  .content {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    padding: 1rem;
    text-align: center;
    padding-top: 100px;
    width: 40%;
    margin-left: auto;
    margin-right: auto;
  }

  .description {
    margin-top: 1rem;
    text-align: center;
  }

  .title {
    font-weight: 900;
    text-align: center;
  }

  @media (max-width: 768px) {
    .content {
      width: 80%;
    }
  }

  .divider {
    width: 100%;
    margin-top: 1rem;
    margin-bottom: 1rem;
    border-radius: 50px;
    background-color: #f1f6fa;
    height: 0.2rem;
    border: none;
  }

  .username {
    background: #f1f6fa;
    border: none;
    width: 100%;
    border-radius: 50px;
    padding: 0.4rem;
    padding-left: 1rem;
    padding-right: 1rem;
  }

  .username:focus {
    outline: none;
  }

  .tweet-wrapper {
    width: 100%;
    margin-top: 1rem;
  }
</style>

<div class="container">
  {#if !loggedIn}
    <Navbar {onClickLogin} />
  {:else}
    <Navbar username={name} {avi} />
  {/if}
  <div class="content">
    <h1 class="title">Find your top tweets!</h1>
    <span class="description">
      Just login with your Twitter account, enter a username and go!
    </span>
    <div class="divider" />
    <input placeholder="Username" class="username" type="text" />
    <div class="divider" />
    {#each tweets as tweet}
      <div class="tweet-wrapper">
        <Tweet {...tweet} />
      </div>
    {/each}
  </div>
</div>
