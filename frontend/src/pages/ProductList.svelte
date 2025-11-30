<script>
  import { api } from "../lib/api";

  let products = [];
  let page = 1;
  let limit = 10;

  async function loadProducts() {
    const res = await api(`/products?page=${page}&limit=${limit}`);
    console.log("RAW RES:", res);

    products = res.data || [];
  }

  function nextPage() {
    page++;
    loadProducts();
  }

  function prevPage() {
    if (page > 1) {
      page--;
      loadProducts();
    }
  }

  loadProducts();
</script>

<div class="container">
  <h2>Product List</h2>

  <table>
    <thead>
      <tr>
        <th>Product</th>
        <th>Variants</th>
        <th>SubVariants</th>
      </tr>
    </thead>

    <tbody>
      {#each products as p}
        <tr>
          <!-- PRODUCT INFO -->
          <td>
            <strong>{p.ProductName}</strong><br />
            Code: {p.ProductCode}<br />
            HSN: {p.HSNCode}<br />
          </td>

          <!-- VARIANTS -->
          <td>
            {#if p.Variants}
              {#each p.Variants as v}
                <div style="margin-bottom: 6px;">
                  <strong>{v.Name}:</strong>

                  {#each v.Options as o, i}
                    {o.Value}{i < v.Options.length - 1 ? ", " : ""}
                  {/each}
                </div>
              {/each}
            {/if}
          </td>

          <!-- SUBVARIANTS -->
          <td>
            {#if p.SubVariants}
              {#each p.SubVariants as sv}
                <div style="margin-bottom: 6px;">
                  {sv.SKU} — <b>Stock:</b>
                  {sv.Stock}
                </div>
              {/each}
            {/if}
          </td>
        </tr>
      {/each}
    </tbody>
  </table>

  <div class="pagination">
    <button on:click={prevPage} disabled={page === 1}>Previous</button>
    <span class="page-label">Page {page}</span>
    <button on:click={nextPage}>Next</button>
  </div>
</div>

<style>
  .container {
    max-width: 1000px;
    margin: auto;
    margin-top: 40px;
    padding: 20px;
    background: #1e1e1e;
    color: white;
    border-radius: 12px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.4);
  }

  table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
    background: #2b2b2b;
  }

  th,
  td {
    padding: 12px;
    border: 1px solid #444;
    text-align: left;
    vertical-align: top;
  }

  button {
    padding: 8px 14px;
    background: #444;
    border-radius: 8px;
    border: none;
    color: white;
    cursor: pointer;
  }

  button:hover {
    background: #666;
  }

  button:disabled {
    background: #333;
    cursor: not-allowed;
  }

  h2 {
    text-align: center;
  }

  .pagination {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 14px;
    margin-top: 20px;
  }

  .page-label {
    font-size: 14px;
    color: #ccc;
  }
</style>
