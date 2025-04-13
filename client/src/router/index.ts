import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Project from '@/views/Project.vue'
import ProjectGroup from '@/views/ProjectGroup.vue'
import Profile from '@/views/profile/Profile.vue'
import Register from '@/views/auth/Register.vue'
import Login from '@/views/auth/Login.vue'
import NewProject from '@/views/NewProject.vue'
import NewGroup from '@/views/NewGroup.vue'
import MyOwnGroups from '@/views/MyOwnGroups.vue'
import Group from '@/views/Group.vue'
import ProfileUpdate from '@/views/profile/ProfileUpdate.vue'
import axios from 'axios'
import type { LoginResponse } from '@/dto/auth'

async function isAuthenticated(){
  try{
    const response = await axios.get("http://localhost:8000/isauth", {
      headers: {
        "Content-Type": "application/json",
        "Authorization": sessionStorage.getItem("authJWT")
      },
      withCredentials: true,
    });
    const loginResponse = response.data as LoginResponse;
    if (loginResponse.Error != ""){
      sessionStorage.removeItem("authJWT");
      return false
    } else if (loginResponse.JWT != "") {
      sessionStorage.setItem("authJWT", loginResponse.JWT)
      return true
    } else {
      return true
    }
  } catch (error){
    console.log(error)
    sessionStorage.removeItem("authJWT");
    return false
  }
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      meta: {requiresAuth: true},
    },
    {
      path: '/project',
      name: 'project',
      component: Project,
      meta: {requiresAuth: true},
    },
    {
      path: '/project-group',
      name: 'project-group',
      component: ProjectGroup,
      meta: {requiresAuth: true},
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile,
      meta: {requiresAuth: true},
    },
    {
      path: '/register',
      name: 'register',
      component: Register,
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/new-project',
      name: 'new-project',
      component: NewProject,
      meta: {requiresAuth: true},
    },
    {
      path: '/new-group',
      name: 'new-group',
      component: NewGroup,
      meta: {requiresAuth: true},
    },
    {
      path: '/my-own-groups',
      name: 'my-own-groups',
      component: MyOwnGroups,
      meta: {requiresAuth: true},
    },
    {
      path: '/group',
      name: 'group',
      component: Group,
      meta: {requiresAuth: true},
    },
    {
      path: '/profile/update',
      name: 'profile-update',
      component: ProfileUpdate,
      meta: {requiresAuth: true},
    },
  ],
})

router.beforeEach(async (to, from, next) => {
  if (to.meta.requiresAuth && !await isAuthenticated()) {
    return next('/login')
  }
  next()
})

export default router
