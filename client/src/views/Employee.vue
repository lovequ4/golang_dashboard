<template>
    <q-table
     title="Employees"
     :rows="rows"
     :columns="columns"
     row-key="name"
     >
     <template v-slot:top-right="props">
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
    
      <!-- Edit Dialog -->
      <q-dialog v-model="showEditDialog">
       
         <q-card style="width: 500px; max-width: 200vw;">
           <q-card-section>
             <q-input v-model="editedRow.name" label="Name"></q-input>
             <q-input v-model="editedRow.email" label="Email"></q-input>
                 <q-select
                 v-model="editedRow.role"
                 label="Role"
                 :options="identityOptions"
                 ></q-select>
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
 
 
 export default {
     data() {
     return {
       columns: [
         { name: 'index', label: 'Index', align: 'left', field: 'index', sortable: true },
         { name: 'name', label: 'Name', align: 'left', field: 'name', sortable: true },
         { name: 'email', label: 'Email', align: 'left', field: 'email', sortable: true },
         { name: 'role', label: 'Role', align: 'left', field: 'role', sortable: true },
         { name: 'actions', label: 'Action'}
       ], 
       rows: [],
       showEditDialog: false,
       editedRow: {},
       identityOptions: ['admin', 'employee'] ,
       showDeleteDialog: false,
       rowToDelete: null
     };
   },
   mounted() {
      const token = localStorage.getItem('token');
      const headers = {
        'Authorization': `${token}`
      };
     axios.get("http://localhost:8080/users",{headers})
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
 
     //edit
     onEdit(row) {
       this.editedRow = { ...row };
       this.showEditDialog = true;
     },
 
     saveEdit() {
       const updatedData = {
         name: this.editedRow.name,
         email: this.editedRow.email,
         role: this.editedRow.role
       };

       const token = localStorage.getItem('token');
       const headers = {
        'Authorization': `${token}`
       };
 
       axios.put(`http://localhost:8080/users/${this.editedRow.id}`, updatedData,{headers})
       .then(response => {
         console.log('User data updated:', response.data);
         this.showEditDialog = false;
         this.refreshData();
       })
       .catch(error => {
         console.error('Error updating user data:', error);
       });
       this.showEditDialog = false;
     },
 
     cancelEdit() {
       this.showEditDialog = false;
     },
 
 
     //delete
     onDelete(row) {
       this.rowToDelete = row;
       this.showDeleteDialog = true;
     },
 
     confirmDelete() {
       if (this.rowToDelete) {
         const rowIndex = this.rows.indexOf(this.rowToDelete);
         if (rowIndex !== -1) {
           this.rows.splice(rowIndex, 1);
         }
       }
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
 
        axios.get('http://localhost:8080/users',{headers})
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