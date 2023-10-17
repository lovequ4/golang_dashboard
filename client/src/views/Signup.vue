<template>
    <div class="q-pa-md" >
        <div class="form-container">
            <q-form
            @submit="onSubmit"
            @reset="onReset"
            class="q-gutter-md"
            >
            <div class="label" >SignUp</div>
  
            <q-input
                filled
                v-model="name"
                label="Your name *"
                type="text"
                lazy-rules
                :rules="[ val => val && val.length > 0 || 'Please Enter your Name']"
            />
  
            <q-input
                filled
                v-model="email"
                label="Your Email *"
                type="email"
                lazy-rules
                :rules="[ val => val && val.length > 0 || 'Please Enter your Email']"
            />
            
            <q-input
                filled
                v-model="password"
                label="Your Password *"
                type="password"
                lazy-rules
                :rules="[ val => val && val.length > 0 || 'Please Enter your Password']"
            />
            
          
            
            <q-select color="orange" filled v-model="model" :options="options"  label="Identity">
              <template v-if="model" v-slot:append>
                <q-icon name="cancel" @click.stop.prevent="model = null" class="cursor-pointer" />
              </template>
            </q-select>
  
            <div>
                <q-btn label="Submit" type="submit" color="primary"/>
                <q-btn label="Reset" type="reset" color="primary" flat class="q-ml-sm" />
            </div>
  
            </q-form>
        </div>
    </div>
</template>
  
<script>
import { useQuasar } from 'quasar'
import { ref } from 'vue'
import axios from 'axios';
import router from '../router'

export default {
setup () {
const $q = useQuasar()
const name = ref(null)
const email = ref(null)
const password = ref(null)
const model = ref(null)


return {
    name,
    email,
    password,
    model,
    options: [
         
        'employee'
    ],

    onSubmit () {
    if (!model.value) {
        $q.notify({
        color: 'red-5',
        textColor: 'white',
        icon: 'warning',
        message: 'please chose one option'
        });
        return;
    }

    const formData = {
        name: name.value,
        email: email.value,
        password: password.value,
        role: model.value
    };
    


        axios.post('http://localhost:8080/users/signup', formData)
        .then(response => {
            
            console.log(response.data)
            $q.notify({
            color: 'green-4',
            textColor: 'white',
            icon: 'cloud_done',
            message: 'Submitted'
            })
            router.push({ name: 'signin' });
        })
        .catch(error => {
            console.error(error)
            $q.notify({
            color: 'red-5',
            textColor: 'white',
            icon: 'warning',
            message: 'submit error'
            })
        })
    },

    onReset () {
        name.value = null
        email.value = null
        password.value = null
        model.value = null
    }
    }
}
}
</script>
  
<style scoped>
.q-pa-md{
margin-top: 10%;
}

.label{
font-size:xx-large;
margin-bottom: 5%;
}

.link{
margin-top: 10%;
}

.form-container {
border: 1px solid #ccc; /* 設置外框樣式 */
padding: 20px; /* 設置內邊距 */
text-align: center; /* 文字置中 */
margin: 0 auto; /* 自動置中 */
max-width: 400px; /* 最大寬度 */
background-color: #f5f5f5;
border-radius: 5px;
box-shadow: 0px 5px 10px #dadadacc;
}
</style>