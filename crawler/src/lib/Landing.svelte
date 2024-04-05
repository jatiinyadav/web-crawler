<script>
    import { onMount } from "svelte";
  
    let message = "";
    let items = [{ header: "", desc: "", href: "" }];
    const endpoint = "http://localhost:8080/api/message";
  
    onMount(async () => {
      try {
        const response = await fetch(endpoint);
        const data = await response.json();
        items = data;
      } catch (error) {
        message = "Error fetching data...";
      }
    });
  </script>
  
  <main>
    <h1>{message}</h1>
  
    <h1>List of Items</h1>
    <ul>
      {#each items as item}
        <li>{item.header} - {item.desc}</li>
      {/each}
    </ul>
  </main>
  