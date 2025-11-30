<script>
  import { api } from "../lib/api";

  let from = "";
  let to = "";
  let report = [];

  async function loadReport() {
    if (!from || !to) {
      alert("Please select both dates");
      return;
    }

    const res = await api(`/stock/report?from=${from}&to=${to}`);
    console.log(res);
    report = res;
  }
</script>

<div class="container">
  <h2>Stock Report</h2>

  <label>From</label>
  <input type="date" bind:value={from} />

  <label>To</label>
  <input type="date" bind:value={to} />

  <button on:click={loadReport}>Load Report</button>

  <table>
    <thead>
      <tr>
        <th>Type</th>
        <th>Qty</th>
        <th>Date</th>
        <th>Product ID</th>
        <th>SubVariant ID</th>
      </tr>
    </thead>
    <tbody>
      {#each report as r}
        <tr>
          <td>{r.TransactionType}</td>
          <td>{r.Quantity}</td>
          <td>{r.TransactionDate}</td>
          <td>{r.ProductID}</td>
          <td>{r.SubVariantID}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>

<style>
  .container {
    max-width: 900px;
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
    margin-top: 10px;
  }

  input {
    width: 100%;
    padding: 8px;
    background: #2b2b2b;
    border-radius: 6px;
    border: 1px solid #444;
    color: white;
    margin-bottom: 10px;
  }

  button {
    padding: 10px 20px;
    background: #444;
    color: white;
    border-radius: 8px;
    border: none;
    cursor: pointer;
    margin-top: 10px;
  }

  button:hover {
    background: #666;
  }

  table {
    width: 100%;
    margin-top: 20px;
    border-collapse: collapse;
    background: #2b2b2b;
  }

  th,
  td {
    padding: 10px;
    border: 1px solid #444;
    text-align: left;
  }

  h2 {
    text-align: center;
  }
</style>
