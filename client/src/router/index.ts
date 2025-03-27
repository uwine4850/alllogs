import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Project from '@/views/Project.vue'
import ProjectGroup from '@/views/ProjectGroup.vue'
import Profile from '@/views/Profile.vue'
import Register from '@/views/Register.vue'
import Login from '@/views/Login.vue'
import NewProject from '@/views/NewProject.vue'
import NewGroup from '@/views/NewGroup.vue'
import MyOwnGroups from '@/views/MyOwnGroups.vue'
import Group from '@/views/Group.vue'
import ProfileUpdate from '@/views/ProfileUpdate.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/project',
      name: 'project',
      component: Project,
    },
    {
      path: '/project-group',
      name: 'project-group',
      component: ProjectGroup,
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile,
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
    },
    {
      path: '/new-group',
      name: 'new-group',
      component: NewGroup,
    },
    {
      path: '/my-own-groups',
      name: 'my-own-groups',
      component: MyOwnGroups,
    },
    {
      path: '/group',
      name: 'group',
      component: Group,
    },
    {
      path: '/profile/update',
      name: 'profile-update',
      component: ProfileUpdate,
    },
  ],
})

export default router
