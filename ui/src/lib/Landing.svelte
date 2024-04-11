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

  async function fetchData() {
    items = [];

    try {
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

    inputValue1 = "";
    inputValue2 = "";
  }
</script>

<main>
  <!-- <h1>{message}</h1> -->

  <!-- svelte-ignore a11y-autofocus -->
  <input type="text" bind:value={inputValue1} />

  <!-- svelte-ignore a11y-autofocus -->
  <input type="text" bind:value={inputValue2} />
  <button on:click={fetchData}>Fetch Data</button>

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
</main>

<style>
  ul {
    text-align: left;
  }

  span {
    font-weight: bold;
    font-size: 1.3rem;
  }
</style>
