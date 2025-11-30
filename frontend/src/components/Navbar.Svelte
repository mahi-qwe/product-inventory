<script>
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  function go(route) {
    dispatch("navigate", route);
  }
</script>

<nav>
  <button on:click={() => go("create")}>Create Product</button>
  <button on:click={() => go("list")}>Product List</button>
  <button on:click={() => go("add")}>Add Stock</button>
  <button on:click={() => go("remove")}>Remove Stock</button>
  <button on:click={() => go("report")}>Stock Report</button>
</nav>

<style>
  nav {
    display: flex;
    justify-content: center;
    gap: 20px;
    padding: 15px;
    background: #121212;
    border-bottom: 1px solid #333;
    position: sticky;
    top: 0;
    z-index: 10;
  }

  button {
    padding: 10px 18px;
    background: #2b2b2b;
    border: 1px solid #444;
    border-radius: 8px;
    color: #f0f0f0;
    font-size: 14px;
    cursor: pointer;
    transition: 0.2s ease;
  }

  button:hover {
    background: #444;
    border-color: #666;
  }

  button:active {
    transform: scale(0.97);
  }
</style>
