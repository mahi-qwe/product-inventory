<script>
  import CreateProduct from "./pages/CreateProduct.svelte";
  import ProductList from "./pages/ProductList.svelte";
  import StockAdd from "./pages/StockAdd.svelte";
  import StockRemove from "./pages/StockRemove.svelte";
  import StockReport from "./pages/StockReport.svelte";
  import Navbar from "./components/Navbar.svelte";

  let route = "create";

  function navigate(to) {
    route = to;
  }
</script>

<Navbar on:navigate={(e) => navigate(e.detail)} />

{#if route === "create"}
  <CreateProduct />
{:else if route === "list"}
  <ProductList />
{:else if route === "add"}
  <StockAdd />
{:else if route === "remove"}
  <StockRemove />
{:else if route === "report"}
  <StockReport />
{/if}
