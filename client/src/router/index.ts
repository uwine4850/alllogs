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

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      meta: { authorized: true },
    },
    {
      path: '/project/:id',
      name: 'project',
      component: Project,
      meta: { authorized: true },
    },
    {
      path: '/project-group',
      name: 'project-group',
      component: ProjectGroup,
      meta: { authorized: true },
    },
    {
      path: '/profile/:id',
      name: 'profile',
      component: Profile,
      meta: { authorized: true },
    },
    {
      path: '/register',
      name: 'register',
      component: Register,
      meta: { isAuth: false },
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
      meta: { isAuth: false },
    },
    {
      path: '/new-project',
      name: 'new-project',
      component: NewProject,
      meta: { authorized: true },
    },
    {
      path: '/new-group',
      name: 'new-group',
      component: NewGroup,
      meta: { authorized: true },
    },
    {
      path: '/my-own-groups',
      name: 'my-own-groups',
      component: MyOwnGroups,
      meta: { authorized: true },
    },
    {
      path: '/group',
      name: 'group',
      component: Group,
      meta: { authorized: true },
    },
    {
      path: '/profile/update/:id',
      name: 'profile-update',
      component: ProfileUpdate,
      meta: { authorized: true },
    },
  ],
})

router.beforeEach((to, from, next) => {
  let isUserAuth: boolean
  if (sessionStorage.getItem('profile') === null && sessionStorage.getItem('authJWT') === null) {
    isUserAuth = false
  } else {
    isUserAuth = true
  }
  if ('isAuth' in to.meta) {
    if (to.meta.isAuth === isUserAuth) {
      next()
    } else {
      next('/')
    }
  }
  if ('authorized' in to.meta) {
    if (to.meta.authorized === isUserAuth) {
      next()
    } else {
      next('/login')
    }
  }
})

export default router
