<script>
  import { api } from "../lib/api.js";

  let product = {
    productID: "",
    productCode: "",
    productName: "",
    productImage: "",
    createdUser: "",
    isFavourite: false,
    active: true,
    hsnCode: "",
  };

  let variants = [{ name: "", options: [""] }];

  function addVariant() {
    variants = [...variants, { name: "", options: [""] }];
  }

  function addOption(i) {
    variants[i].options = [...variants[i].options, ""];
  }

  async function submit() {
    let payload = { product, variants };
    const res = await api("/products", "POST", payload);
    alert("Product created!");
  }
</script>

<div class="container">
  <h2>Create Product</h2>

  <div class="row">
    <div style="flex:1">
      <label>Product ID</label>
      <input bind:value={product.productID} />
    </div>
    <div style="flex:1">
      <label>Product Code</label>
      <input bind:value={product.productCode} />
    </div>
  </div>

  <div class="row">
    <div style="flex:1">
      <label>Product Name</label>
      <input bind:value={product.productName} />
    </div>
    <div style="flex:1">
      <label>Image URL</label>
      <input bind:value={product.productImage} />
    </div>
  </div>

  <div class="row">
    <div style="flex:1">
      <label>Created User UUID</label>
      <input bind:value={product.createdUser} />
    </div>
    <div style="flex:1">
      <label>HSN Code</label>
      <input bind:value={product.hsnCode} />
    </div>
  </div>

  <h3>Variants</h3>

  {#each variants as variant, i}
    <div class="variant-box">
      <label>Variant Name</label>
      <input bind:value={variant.name} />

      <h4>Options</h4>

      {#each variant.options as opt, o}
        <input
          bind:value={variant.options[o]}
          placeholder="Option value"
          style="margin-bottom:10px"
        />
      {/each}

      <button on:click={() => addOption(i)}>+ Add Option</button>
    </div>
  {/each}

  <button on:click={addVariant} style="margin-bottom: 20px;"
    >+ Add Variant</button
  >

  <div class="submit-btn">
    <button on:click={submit}>Submit</button>
  </div>
</div>

<style>
  .container {
    max-width: 800px;
    margin: auto;
    margin-top: 40px;
    padding: 30px;
    background: #1e1e1e;
    border-radius: 12px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.4);
    color: white;
  }
  h2,
  h3,
  h4 {
    text-align: center;
  }
  .row {
    display: flex;
    gap: 20px;
    margin-bottom: 15px;
  }
  .row input {
    flex: 1;
    padding: 8px;
    border-radius: 6px;
    background: #2d2d2d;
    border: 1px solid #444;
    color: white;
  }
  label {
    display: block;
    margin-bottom: 6px;
  }
  .variant-box {
    padding: 15px;
    margin-bottom: 20px;
    background: #2b2b2b;
    border-radius: 10px;
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
  .submit-btn {
    display: flex;
    justify-content: center;
    margin-top: 20px;
  }
</style>
