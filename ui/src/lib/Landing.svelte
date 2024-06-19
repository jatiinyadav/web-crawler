<script lang="ts">
  // import { onMount } from "svelte";
  import type { IResponseType } from "../interfaces/IResponseType";

  // let message = "";
  let items: IResponseType[] = [];
  // const endpoint = "http://localhost:8080/api/message";

  // onMount(async () => {
  //   try {
  //     const response = await fetch(endpoint);
  //     const data = await response.json();
  //     items = data;
  //   } catch (error) {
  //     message = "Error fetching data...";
  //   }
  // });

  let inputValue1 = "";
  let inputValue2 = "";

  let isLoading = false;

  async function fetchData() {
    items = [];

    try {
      isLoading = true;

      console.log(inputValue1);

      const response = await fetch(
        `http://localhost:8080/api/weburl?url1=${inputValue1}&url2=${inputValue2}`
      );
      const data = await response.json();

      console.log(data); // Process the response data as needed
      items = data
    } catch (error) {  
      console.error("Error fetching data:", error);
    }

    isLoading = false;
    inputValue1 = "";
    inputValue2 = "";
  }

  $: showLoader = isLoading;
</script>

<main>
  <!-- <h1>{message}</h1> -->

  <!-- svelte-ignore a11y-autofocus -->
  <input type="text" bind:value={inputValue1} />

  <!-- svelte-ignore a11y-autofocus -->
  <input type="text" bind:value={inputValue2} />
  <button on:click={fetchData}>Fetch Data</button>

  <br/>
  {#if !showLoader}
    <h1>List of Items</h1>
    <ul>
      {items.length}
      {#each items as item}
        <li><span>Header: </span> {item.header}</li>
        <li><span>Description: </span> {item.desc}</li>
        <li><span>URL: </span> {item.href}</li>
        <hr />
      {/each}
    </ul>
  {:else}
    <div class="loader"></div>
  {/if}
</main>

<style>
  ul {
    text-align: left;
  }

  span {
    font-weight: bold;
    font-size: 1.3rem;
  }

  .loader {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  display: inline-block;
  position: relative;
  border: 3px solid;
  border-color: #FFF #FFF transparent;
  box-sizing: border-box;
  animation: rotation 1s linear infinite;
}
.loader::after {
  content: '';  
  box-sizing: border-box;
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  margin: auto;
  border: 3px solid;
  border-color: transparent #FF3D00 #FF3D00;
  width: 24px;
  height: 24px;
  border-radius: 50%;
  animation: rotationBack 0.5s linear infinite;
  transform-origin: center center;
}

@keyframes rotation {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
} 
    
@keyframes rotationBack {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(-360deg);
  }
}
    
</style>
