import { createRouter, createWebHistory } from 'vue-router'
import SignUp from '../views/Signup.vue'
import SignIn from '../views/Signin.vue'
import Index from '../views/Index.vue'
import Employee from '../views/Employee.vue'
import Product from '../views/Product.vue'
import AccessDenied from '../views/AccessDenied.vue'
import Navbar from '../components/Navbar.vue'
import jwtDecode from 'jwt-decode';

const routes = [
    {
      path: '/',
      redirect: '/signin',
    },
    {
      path: '/access-denied',
      name: 'access-denied',
      component: AccessDenied, 
    },  
    {
      path: '/signup',
      name: 'signup',
      component: SignUp
    },
    {
      path: '/signin',
      name: 'signin',
      component: SignIn
    },
    {
      path: '/index',
      
      component: Navbar, 
      children: [
        {
          path: '',
          name: 'index',
          component: Index
        },
        {
          path: 'employee',
          name: 'employee',
          component: Employee,
          beforeEnter: (to, from, next) => {
            const token = localStorage.getItem('token');
            if (token) {
              const decodedToken = jwtDecode(token);
              if (decodedToken.role === 'admin') {
                next();
              } else {
                next('/access-denied');
              }
            } else {
              next('/signin');
            }
          },
        },
        {
          path: 'product',
          name: 'product',
          component: Product
        },
      ],
    },
  ]
  
  const router = createRouter({
    history: createWebHistory(),
    routes
  })
  
  
  export default router
  