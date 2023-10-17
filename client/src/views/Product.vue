<template>
    <q-table
     title="Products"
     :rows="rows"
     :columns="columns"
     row-key="name"
     >
    
     
     <template v-slot:top-right="props">
         <q-btn icon="add" @click="onAddProduct" color="primary" ></q-btn>
         <q-input
         outlined
         dense
         debounce="300"
         v-model="filter"
         placeholder="Search"
         >
         <template v-slot:append>
             <q-icon name="search" />
         </template>
         </q-input>
     </template>
     
     <template v-slot:body-cell-actions="props">
         <q-td :props="props">
         <q-btn icon="mode_edit" @click="onEdit(props.row)"></q-btn>
         <q-btn icon="delete" @click="onDelete(props.row)"></q-btn>
         </q-td>
     </template>
     
     </q-table>

     <!-- Create Dialog -->
    <q-dialog v-model="showCreateDialog">
        <q-card style="width: 500px; max-width: 200vw;">
            <q-card-section>
                <q-input v-model="newProduct.product_name" label="Product Name"></q-input>
                <q-input v-model="newProduct.quantity" @input="handleQuantityInput"  label="Quantity" type="number" :min="1"></q-input>
                <q-input v-model="newProduct.price" @input="handlePriceInput" label="Price" type="number" :min="1"></q-input>
            </q-card-section>
            <q-card-actions align="right">
                <q-btn label="Create" color="primary" @click="onCreateProduct"></q-btn>
                <q-btn label="Cancel" color="negative" @click="cancelCreate"></q-btn>
            </q-card-actions>
        </q-card>
    </q-dialog>
    
      <!-- Edit Dialog -->
      <q-dialog v-model="showEditDialog">
       
         <q-card style="width: 500px; max-width: 200vw;">
           <q-card-section>
             <q-input v-model="editedRow.product_name" label="Product Name"></q-input>
             
                <q-input
                    v-model="editedRow.quantity"
                    label="Quantity"
                    type="number"
                    @input="handleQuantityInput"
                    ></q-input>

                <q-input
                    v-model="editedRow.price"
                    label="Price"
                    type="number"
                    @input="handlePriceInput"
                ></q-input>

           </q-card-section>
           <q-card-actions align="right">
             <q-btn label="Save" color="deep-orange" @click="saveEdit"></q-btn>
             <q-btn label="Cancel" color="primary" @click="cancelEdit"></q-btn>
           </q-card-actions>
         </q-card>
      
     </q-dialog>
 
     <!-- Delete Dialog -->
     <q-dialog v-model="showDeleteDialog">
        
         <q-card>
             <q-card-section>
             <q-card-title>Delete</q-card-title>
             <p>Are you sure you want to delete this?</p>
             </q-card-section>
             <q-card-actions align="right">
             <q-btn label="Delete" color="negative" @click="confirmDelete"></q-btn>
             <q-btn label="Cancel" color="primary" @click="cancelDelete"></q-btn>
             </q-card-actions>
         </q-card>
         
     </q-dialog>
 
 </template>
 
 <script>
 import axios from 'axios'
 import { format,  parseISO } from 'date-fns';
 import jwtDecode from 'jwt-decode';

 export default {
     data() {
     return {
       columns: [
         { name: 'index', label: 'Index', align: 'left', field: 'index', sortable: true },
         { name: 'product_name', label: 'Product Name', align: 'left', field: 'product_name', sortable: true },
         { name: 'quantity', label: 'Quantity', align: 'left', field: 'quantity', sortable: true },
         { name: 'price', label: 'Price', align: 'left', field: 'price', sortable: true },
         { name: 'created_date', label: 'Created Date', align: 'left', field: (row) => format(parseISO(row.create_date), 'yyyy/MM/dd HH:mm'), sortable: true },
         { name: 'actions', label: 'Action'}
       ], 
       rows: [],
       showEditDialog: false,
       editedRow: {},
       showDeleteDialog: false,
       rowToDelete: null,
       showCreateDialog: false,
       newProduct: {},
       userIsAdmin: false
     };
   },
   mounted() {
    const token = localStorage.getItem('token');
    const headers = {
        'Authorization': `${token}`
    };

     axios.get("http://localhost:8080/products",{headers})
       .then(response => {
         console.log(response.data);
         this.rows = response.data;
         this.rows = response.data.map((row, index) => ({ ...row, index: index + 1 }));
       })
       .catch(error => {
         console.error(error);
       });
   }, 
   methods: {
 
    //create
    onAddProduct() {
      const token = localStorage.getItem('token');
      if (token) {
        const decodedToken = jwtDecode(token);
        if (decodedToken.role === 'admin') {
          this.showCreateDialog = true; 
        } else {
         
          this.$q.notify({
            message: 'Your Not Administrator',
            color: 'negative',
            position: 'center',
            timeout: 1000,
          });
        }
      }else {
        this.$router.push({ name: 'signin' });
      }
    },

    handleQuantityInput(value) {
        this.editedRow.quantity = value; 
    },

    handlePriceInput(value){
        this.editedRow.price = value;
    },

    onCreateProduct() {
        const createdData = {
            product_name: this.newProduct.product_name,
            quantity: parseInt(this.newProduct.quantity),
            price: parseInt(this.newProduct.price),
        };
        console.log(createdData)

        const token = localStorage.getItem('token');
        const headers = {
            'Authorization': `${token}`
        };
   
    

        axios.post("http://localhost:8080/products", createdData, { headers })
       .then(response => {
         console.log('product data created:', response.data);
         this.showEditDialog = false;
         this.refreshData();
       })
       .catch(error => {
         console.error('Error updating product  data:', error);
       });
       this.showEditDialog = false;
        
       this.cancelCreate();
    },
    cancelCreate() {
        
        this.newProduct = {};
        this.showCreateDialog = false;
    },


     //edit
     onEdit(row) {
      
      const token = localStorage.getItem('token');
      if (token) {
        const decodedToken = jwtDecode(token);
        if (decodedToken.role === 'admin') {
          this.editedRow = { ...row };
          this.showEditDialog = true;
        } else {
         
          this.$q.notify({
            message: 'Your Not Administrator',
            color: 'negative',
            position: 'center',
            timeout: 1000,
          });
        }
      }else {
        this.$router.push({ name: 'signin' });
      }
     },
 
    handleQuantityInput(value) {
        this.editedRow.quantity = value; 
    },

    handlePriceInput(value){
        this.editedRow.price = value;
    },

     saveEdit() {
       const updatedData = {
         product_name: this.editedRow.product_name,
         quantity: parseInt(this.editedRow.quantity),
         price: parseInt(this.editedRow.price),
       };
      

       const token = localStorage.getItem('token');
       const headers = {
          'Authorization': `${token}`
        };

       axios.put(`http://localhost:8080/products/${this.editedRow.id}`, updatedData, { headers })
       .then(response => {
         console.log('product data updated:', response.data);
         this.showEditDialog = false;
         this.refreshData();
       })
       .catch(error => {
         console.error('Error updating product  data:', error);
       });
       this.showEditDialog = false;
     },
 
     cancelEdit() {
       this.showEditDialog = false;
     },
 
 
     //delete
     onDelete(row) {
      const token = localStorage.getItem('token');
      if (token) {
        const decodedToken = jwtDecode(token);
        if (decodedToken.role === 'admin') {
          this.rowToDelete = row;
          this.showDeleteDialog = true;
        } else {
         
          this.$q.notify({
            message: 'Your Not Administrator',
            color: 'negative',
            position: 'center',
            timeout: 1000,
          });
        }
      }else {
        this.$router.push({ name: 'signin' });
      }
       
       
     },
 
     confirmDelete() {
        const token = localStorage.getItem('token');
        const headers = {
          'Authorization': `${token}`
        };
        console.log(this.rowToDelete.id)
       axios.delete(`http://localhost:8080/products/${this.rowToDelete.id}`, { headers })       
       .then(response => {
         console.log('product data deleted:', response.data);
         this.showDeleteDialog = false;
         this.refreshData();
       })
       .catch(error => {
         console.error('Error updating product  data:', error);
       });
       
  
       this.showDeleteDialog = false;
     },
 
     cancelDelete() {
       this.rowToDelete = null;
       this.showDeleteDialog = false;
     },
 
 
     refreshData() {
        const token = localStorage.getItem('token');
        const headers = {
          'Authorization': `${token}`
        };
        axios.get('http://localhost:8080/products',{headers})
         .then(response => {
             console.log('Data refreshed:', response.data);
             this.rows = response.data;
             this.rows = response.data.map((row, index) => ({ ...row, index: index + 1 }));
         })
         .catch(error => {
             console.error('Error refreshing data:', error);
         });
     }
   }
 };
 </script>
 
 
 <style>
 .my-table-details {
   font-size: 0.85em;
   font-style: italic;
   max-width: 200px;
   white-space: normal;
   color: #555;
   margin-top: 4px;
 }
 </style>