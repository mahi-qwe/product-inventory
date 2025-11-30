<script>
  import { api } from "../lib/api";

  let productID = "";
  let subVariantID = "";
  let quantity = "";

  async function submit() {
    const payload = {
      product_id: productID,
      subvariant_id: subVariantID,
      quantity,
    };

    const res = await api("/stock/remove", "POST", payload);
    alert(res.message || JSON.stringify(res));
  }
</script>

<div class="container">
  <h2>Remove Stock</h2>

  <label>Product ID</label>
  <input bind:value={productID} />

  <label>SubVariant ID</label>
  <input bind:value={subVariantID} />

  <label>Quantity</label>
  <input bind:value={quantity} />

  <button on:click={submit}>Remove Stock</button>
</div>

<style>
  .container {
    max-width: 600px;
    margin: auto;
    margin-top: 40px;
    padding: 25px;
    background: #1e1e1e;
    border-radius: 12px;
    color: white;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.4);
  }
  label {
    display: block;
    margin-bottom: 6px;
    margin-top: 15px;
  }
  input {
    width: 100%;
    padding: 8px;
    background: #2b2b2b;
    border-radius: 6px;
    border: 1px solid #444;
    color: white;
  }
  button {
    margin-top: 20px;
    padding: 10px 20px;
    background: #444;
    color: white;
    border-radius: 8px;
    border: none;
    cursor: pointer;
  }
  button:hover {
    background: #666;
  }
  h2 {
    text-align: center;
  }
</style>
