import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ChatView from '../views/ChatView.vue'
import UserInfoView from '../views/UserInfoView.vue'
import GroupInfoView from '../views/GroupInfoView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/session'},
		{path: '/session', component: LoginView},
		{path: '/home', component: HomeView},
		{path: '/link2', component: HomeView},
		{path: '/some/:id/link', component: HomeView},
		{ path: '/conversation/:conversation', component: ChatView },
		{path : '/profile',component : UserInfoView},
		{path: '/group',component:GroupInfoView},
	]
})

export default router
