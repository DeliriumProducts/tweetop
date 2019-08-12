<script>
  import Navbar from "./components/Navbar"
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
    width: 50%;
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
  </div>
</div>
