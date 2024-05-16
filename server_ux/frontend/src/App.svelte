<script lang="ts">
  import { GetImages } from '../wailsjs/go/main/App.js'
  import { GetTags } from '../wailsjs/go/main/App.js'

  let resultText: string = "Please enter your name below 👇"
  let name: string

  let images = [];
  let tags = [];
  

  function init() {
    getImageList()
  }

  init();

  function getImageList() {
    GetImages().then((result) => (images = result))
  }

  function getTagList() {
    GetTags(name).then((result) => (tags = result))
  }

</script>

<main>
  <div>
    <div class="input-box">
      <select bind:value={name}>
        {#each images as image}
          <option value={image}>
            {image}
          </option>
          {/each}
      </select>
      <button class="btn" on:click={getTagList}>Get Information</button>
      <div class="result">{tags}</div>  
    </div>
  </div>
</main>

<style>

  main {
    text-align: center;
    padding: 1em;
    margin: 0 auto;
    height: 100vh;
  }


  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

</style>
