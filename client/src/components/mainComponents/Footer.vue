<template>
    <footer class='footer-box'>
        <router-link to="/" class="dev-head-main">
            <img class='devschool-cat' :src="devschoolPNG">
            дев.школа
        </router-link>
        <h1 class="devschool-footer-text">{{ curDate }} pre-alpha</h1>
    </footer>
</template>

<script setup>
import devschoolPNG from '@/views/bgImages/devschool.png'
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue'
const isFixed = ref(false)
const curDate = new Date().getFullYear();

const updateFooterPosition = () => {
    nextTick(() => {
        isFixed.value = document.body.scrollHeight <= window.innerHeight
    })
}

onMounted(() => {
    updateFooterPosition()
    window.addEventListener('resize', updateFooterPosition)
})

onBeforeUnmount(() => {
    window.removeEventListener('resize', updateFooterPosition)
})
</script>

<style scoped>
.footer-box {
    display: flex;
    flex-flow: row wrap;
    justify-content: space-between;
    align-items: center;
    z-index: 1;
    width: 90%;
    max-width: 1200px;
    min-height: 64px;
    margin-top: 40px;
    height: 85px;
    border-top: 1px solid rgba(151, 151, 151, 0.815);
}

.dev-head-main {
    text-decoration: none;
}

.devschool-footer-text {
    font-weight: normal;
    font-size: 20px;
    color: rgb(116, 116, 116);
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
  color: rgb(33, 33, 33);
  text-decoration: none;
}

@media (max-width: 750px) {

    .devschool-footer-text,
    .dev-head-main {
        font-size: 17px;
    }
}

@media (max-width: 500px) {

    .devschool-footer-text,
    .dev-head-main {
        font-size: 15px;
    }
}
</style>