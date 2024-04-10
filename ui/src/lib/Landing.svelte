<script>
  // import { onMount } from "svelte";

  // let message = "";
  let items = [{ header: "", desc: "", href: "" }];
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

  let inputValue = "";

  async function fetchData() {
    try {
      const response = await fetch(
        `http://localhost:8080/api/weburl?input=${inputValue}`
      );
      const data = await response.json();
      console.log(data); // Process the response data as needed
      items = data;
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }
</script>

<main>
  <!-- <h1>{message}</h1> -->

  <h1>List of Items</h1>
  <ul>
    {#each items as item}
      <li>{item.header} - {item.desc}</li>
    {/each}
  </ul>

  <input type="text" bind:value={inputValue} />
  <button on:click={fetchData}>Fetch Data</button>
</main>
