import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import QuestionList from '../components/QuestionList.vue';

const routes = [
  { path: '/', component: Home },
  { path: '/survey', component: QuestionList }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
