<template>
    <q-layout view="hHh lpR fFf">
  
      <q-header elevated class="bg-primary text-white">
        <q-toolbar>
          <q-btn dense flat round icon="menu" @click="toggleLeftDrawer" />
  
          <q-toolbar-title>
            <q-avatar>
              <img src="https://cdn.quasar.dev/logo-v2/svg/logo-mono-white.svg">
            </q-avatar>
            MyNode
          </q-toolbar-title>
        </q-toolbar>
      </q-header>
  
      <q-drawer show-if-above v-model="leftDrawerOpen" side="left" bordered>
        
            <q-list bordered padding class="rounded-borders text-primary">
            <q-item
                clickable
                v-ripple
                :active="link === 'inbox'"
                @click="link = 'inbox'"
                active-class="my-menu-link"
                to="/index"
            >
                <q-item-section avatar>
                    <span class="material-icons">dashboard</span>
                </q-item-section>

                <q-item-section>DashBoard</q-item-section>
            </q-item>

            <q-item
                clickable
                v-ripple
                :active="link === 'product'"
                @click="link = 'product'"
                active-class="my-menu-link"
                to="/index/product"
            >
                <q-item-section avatar>
                  <span class="material-icons">inventory</span>
                </q-item-section>

                <q-item-section>Product Management</q-item-section>
            </q-item>


            <q-item
                clickable
                v-ripple
                :active="link === 'employee'"
                @click="link = 'employee'"
                active-class="my-menu-link"
                to="/index/employee"
            >
                <q-item-section avatar>
                    <span class="material-icons">badge</span>
                </q-item-section>

                <q-item-section>Employee Management</q-item-section>
            </q-item>
          
            

            <!-- <q-separator spaced /> -->

            <q-item
                clickable
                v-ripple
                :active="link === 'settings'"
                @click="link = 'settings'"
                active-class="my-menu-link"
                class="setting-link"
            >
                <q-item-section avatar>
                <q-icon name="settings" />
                </q-item-section>

                <q-item-section>Settings</q-item-section>
            </q-item>

            <q-item
                clickable
                v-ripple
                :active="link === 'LogOut'"
                @click="LogOut"
                active-class="my-menu-link"
                class="LogOut-link"
                to="/signin"
            >
                <q-item-section avatar>
                    <span class="material-icons">
                    account_circle
                    </span>
                </q-item-section>

                <q-item-section>LogOut</q-item-section>
            </q-item>
            </q-list>
  
      </q-drawer>
  
      <q-page-container>
        <router-view />
      </q-page-container>
  
    </q-layout>
  </template>
  
<script>
  import { ref } from 'vue'
  import { RouterView } from 'vue-router';

  export default {
    setup () {
      const leftDrawerOpen = ref(false)

      return {
        leftDrawerOpen,
        toggleLeftDrawer () {
          leftDrawerOpen.value = !leftDrawerOpen.value
        },
        link: ref('inbox')
      }
    },
    methods: {
      LogOut() {
        localStorage.removeItem('token');
      }
    }
  }
</script>


<style lang="sass">
.my-menu-link
  color: white
  background: #F2C037


</style>

<style scoped>
.material-icons{
    font-size: x-large;
}
</style>
