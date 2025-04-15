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
import { type AxiosResponse } from 'axios'
import type { LoginResponseMessage } from '@/dto/auth'
import { AsyncRequest } from '@/classes/request'

async function isAuthenticated(){
  let isOK: boolean = false
  const req = new AsyncRequest("http://localhost:8000/isauth", {
    headers: {
      "Content-Type": "application/json",
      "Authorization": sessionStorage.getItem("authJWT")
    },
    withCredentials: true,
  });
  req.onResponse((response: AxiosResponse) => {
    const loginResponse = response.data as LoginResponseMessage;
    if (loginResponse.Error != ""){
      sessionStorage.removeItem("authJWT");
      sessionStorage.removeItem("profile");
      isOK = false;
    } else if (loginResponse.JWT != "") {
      sessionStorage.setItem("authJWT", loginResponse.JWT);
      isOK = true;
    } else {
      isOK = true;
    }
  });
  req.onError((error: unknown) => {
    sessionStorage.removeItem("authJWT");
    isOK = false;
  });
  await req.get();
  return isOK
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
      path: '/profile/:id',
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
