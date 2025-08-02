<template>
  <header :class="['header-box', { 'header-fixer': showMenu }]">
    <router-link to="/" class="dev-head-main">
      <img class="devschool-cat" :src="devschoolPNG" />
      дев.школа
    </router-link>

    <nav :class="['nav-bar', { 'nav-visible': showMenu }]">
      <router-link to="/courses">Курсы</router-link>
      <router-link to="/tests">Тесты</router-link>
      <router-link to="/help">Поддержка</router-link>

      <button class="profile-box-mobile" @click="profileButton">
        <p class="profile-link">Профиль</p>
      </button>
    </nav>

    <button class="profile-box" @click="profileButton">
      <p class="profile-link">Профиль</p>
    </button>

    <button class="display-button" @click="toggleMenu">
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
  </header>
</template>

<script setup>
import devschoolPNG from "@/assets/devschool.png";
import { ref, onMounted, onBeforeUnmount } from "vue";
import { hasCookie } from "@/modules/auth/cookie";
import router from "@/router";
const token = hasCookie("tkn");
const showMenu = ref(false);

function toggleMenu() {
  showMenu.value = !showMenu.value;
}

function handleResize() {
  if (window.innerWidth > 930) {
    showMenu.value = false;
  }
}

function profileButton() {
  if (token) {
    router.push("/profile");
  } else {
    router.push("/login");
  }
}

onMounted(() => {
  window.addEventListener("resize", handleResize);
});

onBeforeUnmount(() => {
  window.removeEventListener("resize", handleResize);
});
</script>

<style scoped>
.header-box {
  position: fixed;
  top: 30px;
  display: flex;
  left: 0;
  right: 0;
  margin: 0 auto;
  flex-flow: row wrap;
  justify-content: space-between;
  align-items: center;
  z-index: 1;
  width: 99%;
  min-height: 64px;
  max-width: 1200px;
  padding: 4px 4px;
  background: white;
  border: 1px solid rgb(221, 221, 221);
  border-radius: 16px;
}

.devschool-cat {
  width: 50px;
  height: 50px;
}

.dev-head-main {
  display: flex;
  align-items: center;
  font-size: 24px;
  font-weight: bolder;
  padding: 0px 0px 0px 15px;
  gap: 5px;
  color: rgb(8, 8, 8);
  text-decoration: none;
}

.nav-bar {
  display: flex;
  flex-flow: wrap;
  gap: 20px;
}

.nav-bar a {
  text-decoration: none;
  color: rgb(95, 95, 95);
  font-size: 16px;
  border: none;
  background: rgb(255, 255, 255);
  border-radius: 8px;
  padding: 15px 24px;
  transition: background 0.3s ease;
}

.nav-bar a:hover {
  background: rgba(216, 216, 216, 0.253);
  color: rgb(0, 0, 0);
}

.profile-box {
  display: flex;
  flex-wrap: wrap;
  gap: 40px;
  border: none;
  background: white;
}

.profile-box-mobile {
  border: none;
  display: none;
  background: white;
}

.profile-link {
  background: rgb(59, 59, 59);
  color: white;
  border: none;
  padding: 15px;
  font-size: 16px;
  font-weight: bold;
  border-radius: 10px;
  cursor: pointer;
  text-decoration: none;
  transition: background 0.4s ease;

  &:hover {
    background: #555555;
  }
}

.display-button {
  border: none;
  background: none;
  display: none;
  padding: 0px 25px;
  cursor: pointer;
}

svg {
  width: 24px;
  height: 24px;
}

@media (max-width: 930px) {
  .profile-box {
    display: none;
  }

  .display-button {
    display: block;
  }

  .profile-box-mobile {
    display: block;
  }

  .header-fixer {
    border-radius: 16px 16px 0px 0px;
  }

  .nav-bar {
    display: none;
    flex-direction: column;
    width: calc(100% + 2px);
    margin-top: -10px;
    background-color: white;
    padding: 10px 10px;
    margin-left: -1px;
    gap: 10px;
    position: absolute;
    top: 100%;
    left: 0;
    border: 1px solid rgb(221, 221, 221);
    border-top: none;
    border-radius: 0px 0px 16px 16px;
  }

  .nav-visible {
    display: flex;
  }
}

@media (max-width: 450px) {
  .dev-head-main {
    font-size: 20px;
  }

  .header-box {
    min-height: 55px;
  }

  .three-lines-image {
    width: 18px;
    height: 18px;
  }

  .x-image {
    width: 15px;
    height: 15px;
  }

  .nav-bar {
    a {
      font-size: 13px;
    }
  }

  .profile-link {
    font-size: 12px;
  }
}
</style>
