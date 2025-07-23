import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/mainPages/Home.view.vue";
import RegisterView from "@/views/auth/Register.view.vue";
import ProfileView from "@/views/userPages/Profile.view.vue";
import LoginView from "@/views/auth/Login.view.vue";
import CoursesView from "@/views/mainPages/Courses.view.vue";
import HelpView from "@/views/mainPages/Help.view.vue";
import TestsView from "@/views/mainPages/Tests.view.vue";
import UserPanelView from "@/views/userPages/UserPanel.view.vue";

const routes = [
  {
    path: "/",
    name: "home",
    meta: { title: "Дев.Школа" },
    component: HomeView,
  },
  {
    path: "/login",
    name: "login",
    meta: { title: "Авторизация" },
    component: LoginView,
  },
  {
    path: "/registration",
    name: "registration",
    meta: { title: "Регистрация" },
    component: RegisterView,
  },
  {
    path: "/profile",
    name: "profile",
    meta: { title: "Профиль" },
    component: ProfileView,
  },
  {
    path: "/courses",
    name: "courses",
    meta: { title: "Курсы" },
    component: CoursesView,
  },
  {
    path: "/help",
    name: "help",
    meta: { title: "Поддержка" },
    component: HelpView,
  },
  {
    path: "/tests",
    name: "tests",
    meta: { title: "Тесты" },
    component: TestsView,
  },
  {
    path: "/user-panel",
    name: "user-panel",
    meta: { title: "Панель пользователя" },
    component: UserPanelView,
  },
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: () => import("@/views/NotFound.vue"),
    meta: { title: "Страница не найдена" },
  },
];

const coursePages = import.meta.glob("@/views/courses/*/*Course.vue");
for (const path in coursePages) {
  const match = path.match(/\/courses\/([^/]+)\/[^/]+Course\.vue$/);
  if (match) {
    const courseFolder = match[1];
    routes.push({
      path: `/courses/${courseFolder}`,
      name: `${courseFolder}-course`,
      meta: { title: `${courseFolder} | Список курса` },
      component: coursePages[path],
    });
  }
}

const lessonPages = import.meta.glob("@/views/courses/*/chapters/*/*.vue");
for (const path in lessonPages) {
  const match = path.match(
    /\/courses\/([^/]+)\/chapters\/([0-9]+)\/([0-9.]+)\.vue$/,
  );
  if (match) {
    const [_, courseFolder, chapter, lesson] = match;
    routes.push({
      path: `/courses/${courseFolder}/${chapter}/${lesson}`,
      name: `${courseFolder}-${chapter}-${lesson}`,
      meta: { title: `${courseFolder} | Урок ${lesson}` },
      component: lessonPages[path],
    });
  }
}

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  document.title = to.meta.title || "дев.школа";
  next();
});

export default router;
