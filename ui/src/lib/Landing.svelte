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

  let isLoading = false;

  async function fetchData() {
    items = [];

    try {
      isLoading = true;

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

    isLoading = false;
    inputValue1 = "";
    inputValue2 = "";
  }

  $: showLoader = isLoading;
</script>

{#if mount}
  <main>
    <!-- <h1>{message}</h1> -->
    <!-- svelte-ignore a11y-autofocus -->
    <input type="text" bind:value={inputValue1} />

    <!-- svelte-ignore a11y-autofocus -->
    <input type="text" bind:value={inputValue2} /><br />
    <button transition:fade={{ delay: 100, duration: 400 }} on:click={fetchData}>
      Fetch Data
    </button>

    {#if items.length > 0}
      <button transition:fade={{ delay: 100, duration: 400 }} on:click={() => { items = []; }}>
        Reset
      </button>

      {#if !showLoader}
        <div style="display: flex;">
          <div id="loi">List of Items</div>
          <div><b>Total items:</b> {items.length}</div>
        </div>
        <hr />
        <ul>
          {#each items as item}
            <li>
              <div transition:slide={{ duration: 400, delay: 100, axis: "y" }}>
                <span>Header: </span> {item.header}<br/>
                <span>Description: </span> {item.desc}<br/>
                <span>URL: </span> {item.href}
                <hr />
              </div>
            </li>
          {/each}
        </ul>
      {:else}
        <div class="loader"></div>
      {/if}
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
</style>
