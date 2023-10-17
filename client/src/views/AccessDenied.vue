<template>
  
      <q-dialog v-model="dialog" @hide="redirectToIndex">
      <q-card style="width: 700px; max-width: 100vw;">
        <q-card-section>
          <div class="text-h6"> Access Denied</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          Sorry, you don't have permission to access this page.
        </q-card-section>

        <q-card-actions align="right" class="bg-white text-teal">
          <q-btn color="primary" @click="redirectToIndex">OK</q-btn>
        </q-card-actions>
      </q-card>
    </q-dialog>

  
  </template>
  
  <script>
  export default {
    data() {
      return {
        dialog: true, // 控制對話框的顯示
      };
    },
    methods: {
      redirectToIndex() {
        // 跳轉到 /index 路由
        this.$router.push({ name: 'index' });
      },
    },
  };
  </script>
  