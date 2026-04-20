import { createRouter, createWebHistory } from 'vue-router'

import SelectFenrir from '@/views/SelectFenrir.vue'
import TGFenrir from '@/views/tg/fenrir.vue'
import SettingsFenrir from '@/views/Settings.vue'


const routes = [
  { path: '/', component: SelectFenrir },
  { path: '/tg', component: TGFenrir },
  { path: '/settings', component: SettingsFenrir }
]

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
})

export default router
