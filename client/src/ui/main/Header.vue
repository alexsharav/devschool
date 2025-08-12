<template>
  <header :class="['header-box', { 'header-fixer': showMenu }]" role="banner">
    <router-link to="/" class="dev-head-main">
      <img class="devschool-cat" :src="devschoolPNG" alt="dev.school logo" />
      дев.школа
    </router-link>

    <!-- Десктопная навигация -->
    <nav class="nav-bar desktop-only" aria-label="Главная навигация">
      <router-link to="/courses">Курсы</router-link>
      <router-link to="/tests">Тесты</router-link>
      <router-link to="/media">Медиа</router-link>
    </nav>

    <button
      class="profile-box desktop-only"
      @click="profileButton"
      type="button"
    >
      <p class="profile-link">Профиль</p>
    </button>

    <!-- Кнопка бургер (мобайл) -->
    <button
      class="display-button mobile-only"
      :aria-expanded="showMenu ? 'true' : 'false'"
      aria-controls="mobile-nav"
      @click="toggleMenu"
      type="button"
    >
      <svg
        v-if="showMenu"
        class="x-image"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        viewBox="0 0 24 24"
      >
        <line x1="18" y1="6" x2="6" y2="18" />
        <line x1="6" y1="6" x2="18" y2="18" />
      </svg>
      <svg
        v-else
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <line x1="3" y1="6" x2="21" y2="6" />
        <line x1="3" y1="12" x2="21" y2="12" />
        <line x1="3" y1="18" x2="21" y2="18" />
      </svg>
    </button>

    <!-- Мобильная навигация с плавной высотой (0.5s) -->
    <Transition @enter="onEnter" @after-enter="onAfterEnter" @leave="onLeave">
      <nav
        v-show="showMenu"
        id="mobile-nav"
        class="nav-bar mobile-panel"
        aria-label="Главная навигация (мобильная)"
      >
        <router-link to="/courses" @click="closeMenu">Курсы</router-link>
        <router-link to="/tests" @click="closeMenu">Тесты</router-link>
        <router-link to="/media" @click="closeMenu">Медиа</router-link>

        <button
          class="profile-box-mobile"
          @click="onMobileProfile"
          type="button"
        >
          <p class="profile-link">Профиль</p>
        </button>
      </nav>
    </Transition>
  </header>
</template>

<script setup>
import devschoolPNG from "@/assets/devschool.png";
import { ref, onMounted, onBeforeUnmount } from "vue";
import token from "@/modules/tokens/token";
import router from "@/router/router";

const showMenu = ref(false);

function toggleMenu() {
  showMenu.value = !showMenu.value;
}
function closeMenu() {
  showMenu.value = false;
}

/* Плавное раскрытие/сворачивание: ровно 0.5s (height + opacity) */
function onEnter(el) {
  el.style.height = "0px";
  el.style.opacity = "0";
  void el.offsetHeight; // force reflow
  el.style.transition = "height 500ms ease, opacity 100ms ease";
  el.style.height = el.scrollHeight + "px";
  el.style.opacity = "1";
}
function onAfterEnter(el) {
  el.style.height = "auto";
  el.style.transition = "";
}
function onLeave(el) {
  el.style.height = el.scrollHeight + "px";
  el.style.opacity = "1";
  void el.offsetHeight;
  el.style.transition = "height 500ms ease, opacity 500ms ease";
  el.style.height = "0px";
  el.style.opacity = "0";
}

function profileButton() {
  if (token.getAccessToken()) router.push("/profile");
  else router.push("/login");
}

function onMobileProfile() {
  closeMenu();
  profileButton();
}

/* Закрывать меню при ресайзе на десктоп */
function handleResize() {
  if (window.innerWidth > 930) showMenu.value = false;
}
onMounted(() => window.addEventListener("resize", handleResize));
onBeforeUnmount(() => window.removeEventListener("resize", handleResize));
</script>

<style scoped>
.header-box {
  position: fixed;
  top: 30px;
  left: 0;
  right: 0;
  margin: 0 auto;
  z-index: 10;

  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-wrap: wrap;

  width: 99%;
  max-width: 1200px;
  min-height: 64px;
  padding: 4px;

  background: #fff;
  border: 1px solid rgb(221, 221, 221);
  border-radius: 16px;
}
.header-fixer {
  border-radius: 16px 16px 0 0;
}

.devschool-cat {
  width: 50px;
  height: 50px;
}
.dev-head-main {
  display: flex;
  align-items: center;
  gap: 8px;
  padding-left: 15px;
  font-size: 24px;
  font-weight: 800;
  color: #080808;
  text-decoration: none;
}

.nav-bar {
  display: flex;
  gap: 20px;
}

.nav-bar a {
  text-decoration: none;
  color: rgb(95, 95, 95);
  font-size: 16px;
  border: none;
  background: #fff;
  border-radius: 8px;
  padding: 12px 20px;
  transition: background 0.25s ease, color 0.25s ease, box-shadow 0.25s ease;
}
.nav-bar a:hover {
  background: rgba(216, 216, 216, 0.25);
  color: #000;
}

.profile-box,
.profile-box-mobile {
  border: none;
  background: transparent;
}
.profile-link {
  background: rgb(59, 59, 59);
  color: #fff;
  border: none;
  padding: 12px 16px;
  font-size: 16px;
  font-weight: 700;
  border-radius: 10px;
  cursor: pointer;
  transition: background 0.25s ease;
}
.profile-link:hover {
  background: #555555;
}

.display-button {
  border: none;
  background: none;
  padding: 0 16px;
  cursor: pointer;
}
svg {
  width: 24px;
  height: 24px;
}

/* Desktop / Mobile switches */
.desktop-only {
  display: flex;
}
.mobile-only {
  display: none;
}

/* Мобильная панель (анимация через JS-хуки, поэтому только overflow:hidden) */
.mobile-panel {
  flex-direction: column;
  width: calc(100% + 2px);
  margin-top: -10px;
  margin-left: -1px;

  position: absolute;
  top: 100%;
  left: 0;

  background: #fff;
  padding: 10px;
  gap: 10px;

  border: 1px solid rgb(221, 221, 221);
  border-top: none;
  border-radius: 0 0 16px 16px;

  overflow: hidden;
}
.profile-box-mobile {
  display: block;
}

@media (max-width: 930px) {
  .desktop-only {
    display: none;
  }
  .mobile-only {
    display: block;
  }
}
@media (max-width: 450px) {
  .dev-head-main {
    font-size: 20px;
  }
  .header-box {
    min-height: 55px;
  }
  .display-button svg {
    width: 20px;
    height: 20px;
  }
  .nav-bar a {
    font-size: 14px;
    padding: 10px 14px;
  }
  .profile-link {
    font-size: 14px;
  }
}
</style>
