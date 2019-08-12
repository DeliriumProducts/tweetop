<script>
  export let name
  import Navbar from "./components/Navbar"
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
  }

  .description {
    margin-top: 1rem;
    text-align: center;
  }

  .title {
    font-weight: 900;
    text-align: center;
  }

  .divider {
    width: 100%;
    background-color: #f1f6fa;
  }
</style>

<div class="container">
  <Navbar />
  <div class="content">
    <h1 class="title">Find your top tweets!</h1>
    <span class="description">
      Just login with your Twitter account, enter a username and go!
    </span>
    <hr class="divider" />
  </div>
</div>
