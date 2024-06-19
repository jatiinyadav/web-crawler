<script lang="ts">
  import { onMount } from "svelte";
  import type { IResponseType } from "../interfaces/IResponseType";
  import { fade, slide } from "svelte/transition";
  let mount: boolean = false;
  onMount(() => {
    mount = true;
  });
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
        `http://localhost:8080/api/weburl?url1=${inputValue1}&url2=${inputValue2}`,
      );
      const data = await response.json();
      console.log(data); // Process the response data as needed
      items = data;
    } catch (error) {
      console.error("Error fetching data:", error);
    }

    inputValue1 = "";
    inputValue2 = "";
  }
</script>

{#if mount}
  <main>
    <!-- <h1>{message}</h1> -->
    <!-- svelte-ignore a11y-autofocus -->
    <input type="text" bind:value={inputValue1} />

    <!-- svelte-ignore a11y-autofocus -->
    <input type="text" bind:value={inputValue2} /><br />
    <button transition:fade={{ delay: 100, duration: 400 }} on:click={fetchData}
      >Fetch Data</button
    >
   

    {#if items.length > 0}
    <button
    transition:fade={{ delay: 100, duration: 400 }}
    on:click={() => {
      items = [];
    }}>Reset</button
  >
      <flex>
        <div id="loi">List of Items</div>
        <div>(<b> Total items:</b> {items.length})</div></flex
      >
      <hr />

      <ul>
        <div transition:slide={{ duration: 400, delay: 100, axis: "y" }}>
          {#each items as item}
            <div transition:slide={{ duration: 400, delay: 100, axis: "y" }}>
              <li><span>Header: </span> {item.header}</li>
              <li><span>Description: </span> {item.desc}</li>
              <li><span>URL: </span> {item.href}</li>
              <hr />
            </div>
          {/each}
        </div>
      </ul>
    {/if}
  </main>
{/if}

<style>
  ul {
    text-align: left;
  }
  input {
    border-radius: 5%;
    outline: none;
  }
  span {
    font-weight: bold;
    font-size: 1.3rem;
  }
  #loi {
    font-size: xx-large;
  }
</style>
