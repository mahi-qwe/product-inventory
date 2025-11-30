📦 Product Inventory System (Go + Svelte)

A complete inventory system built as part of a machine test challenge.
Includes product management, variants, auto-generated subvariants, stock operations, and stock reports.

📂 Folder Structure:
product-inventory/  
│  
├── backend/     # Go + Gin + GORM + PostgreSQL  
├── frontend/    # Svelte (Vite)  
└── README.md    # Main documentation  

⛏️ Technologies Used  
Go (1.24.5)  
Gin Framework  
GORM ORM  
PostgreSQL  
Decimal package  
UUID for IDs  

📌 Backend Setup  
cd backend  
go mod tidy  
createdb inventory  
Update database connection in internal/database/database.go(as your system)  
go run main.go  
Backend runs on: http://localhost:8080  
 
📚 API Endpoints  
👉🏻 POST /products  
Create product + variants + auto-subvariants.  
👉🏻 GET /products?page=1&limit=10  
List products with variants and subvariants.  
👉🏻 POST /stock/add  
Add stock (purchase).  
👉🏻 POST /stock/remove  
Remove stock (sale).  
👉🏻 GET /stock/report?from=YYYY-MM-DD&to=YYYY-MM-DD  
Stock report filtered by date.  

🎨 Frontend — Svelte (Vite)  
cd frontend  
npm install  
npm run dev  
Frontend runs on: http://localhost:5173  
