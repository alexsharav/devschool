import { createRouter, createWebHistory } from "vue-router";
import { requireAuth, requireGuest } from "@/modules/tokens/accesschecker";
import HomeView from "@/pages/mainPages/Home.view.vue";
import RegisterView from "@/pages/auth/Register.view.vue";
import LoginView from "@/pages/auth/Login.view.vue";
import CoursesView from "@/pages/mainPages/Courses.view.vue";
import TestsView from "@/pages/mainPages/Tests.view.vue";
import MediaView from "@/pages/mainPages/Media.view.vue";
import ProfileView from "@/pages/userPages/Profile.view.vue";

const routes = [
  {
    path: "/",
    name: "home",
    meta: { title: "Дев.Школа" },
    component: HomeView,
  },
  {
    path: "/profile",
    name: "profile",
    meta: { title: "Профиль" },
    component: ProfileView,
    beforeEnter: requireAuth,
  },
  {
    path: "/media",
    name: "media",
    meta: { title: "Медиа" },
    component: MediaView,
  },
  {
    path: "/login",
    name: "login",
    meta: { title: "Аутентификация" },
    component: LoginView,
    beforeEnter: requireGuest,
  },
  {
    path: "/register",
    name: "register",
    meta: { title: "Регистрация" },
    component: RegisterView,
    beforeEnter: requireGuest,
  },
  {
    path: "/courses",
    name: "courses",
    meta: { title: "Курсы" },
    component: CoursesView,
  },
  {
    path: "/tests",
    name: "tests",
    meta: { title: "Тесты" },
    component: TestsView,
  },
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: () => import("@/pages/otherPages/NotFound.vue"),
    meta: { title: "Страница не найдена" },
  },
  {
    path: "/accessdenied",
    name: "AccessDenied",
    component: () => import("@/pages/otherPages/AccessDenied.vue"),
    meta: { title: "Запрос отказан" },
  },
];

const coursePages = import.meta.glob("@/pages/courses/*/*Course.vue");
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

const lessonPages = import.meta.glob("@/pages/courses/*/chapters/*/*.vue");
for (const path in lessonPages) {
  const match = path.match(
    /\/courses\/([^/]+)\/chapters\/([0-9]+)\/([0-9.]+)\.vue$/
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
