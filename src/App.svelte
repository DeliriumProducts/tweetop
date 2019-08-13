<script>
  import Navbar from "./components/Navbar"
  import Tweet from "./components/Tweet"
  let name
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
</style>

<div class="container">
  <Navbar />
  <div class="content">
    <h1 class="title">Find your top tweets!</h1>
    <span class="description">
      Just login with your Twitter account, enter a username and go!
    </span>
    <div class="divider" />
    <input placeholder="Username" class="username" type="text" />
    <div class="divider" />
    <Tweet
      username="Dan Abramov"
      handle="@dan_abramov"
      avi="https://pbs.twimg.com/profile_images/1096807971374448640/rVCDhxkG.png"
      rt="256"
      content=" Lorem, ipsum dolor sit amet consectetur adipisicing elit.
      Minima, iste voluptatem omnis repellendus voluptatum excepturi vel
      distinctio repellat cumque quos dolorem quisquam praesentium reprehenderit
      atque quod fugit voluptatibus aperiam? Ipsum. " />
  </div>
</div>
