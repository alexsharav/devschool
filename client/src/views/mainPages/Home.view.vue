<template>
  <Header />

  <StarterTab
    mainText="Дев.Школа"
    secondaryText="образовательные курсы и тесты для всех"
  >
    <template #buttons>
      <div class="buttons">
        <button class="dev-tab-start-courses" @click="router.push('/courses')">
          <p class="dev-tab-courses">начни с наших курсов</p>
        </button>

        <button class="dev-tab-start-tests" @click="router.push('/tests')">
          <p class="dev-tab-tests">начни с наших тестов</p>
        </button>
      </div>
    </template>
  </StarterTab>

  <section class="devschool-base-info">
    <h1>
      <span class="devschool-base-info-text">Дев.Школа</span> - Платформа для
      тех, кто хочет войти в <b>IT</b>. Курсы по программированию, практические
      задания и поддержка от экспертов.
      <span style="text-decoration: underline"
        >Учиться легко — даже с нуля.</span
      >
    </h1>

    <div>
      <p>
        Наша образовательная платформа построена для развития технических
        навыков: от базового программирования до продвинутых технологий. Мы
        верим, что каждый может стать разработчиком.
      </p>

      <p>
        Наши курсы — это путь от первых строк кода до уверенного старта в
        карьере. Обучение, практика, результат.
      </p>
    </div>
  </section>

  <Footer />
</template>

<script setup>
import Header from "@/components/mainComponents/Header.vue";
import Footer from "@/components/mainComponents/Footer.vue";
import StarterTab from "@/components/mainComponents/StarterTab.vue";
import { useRouter } from "vue-router";

const router = useRouter();
</script>

<style scoped>
.course-features {
  margin-top: 20vh;
}

.tab {
  height: 600px;
}

.buttons {
  display: flex;
  flex-flow: row wrap;
  gap: 20px;
  margin-top: 15px;

  button {
    border-radius: 12px;
    display: flex;
    flex-flow: row wrap;
    padding: 16px 0px;
    align-items: center;
    cursor: pointer;
    border: none;
    transition: 0.3s ease-in-out;
    background: rgb(59, 59, 59);

    &:hover {
      background: rgb(65, 65, 65);
      transform: scale(1.03);
      box-shadow: 0 4px 8px rgb(150, 150, 150);
      background: white;

      p {
        color: rgb(45, 45, 45);
      }
    }

    p {
      font-size: 16px;
      text-decoration: none;
      font-weight: bold;
      padding: 0px 15px;
      color: white;
      transition: 0.3s ease-in-out;
    }
  }
}

.devschool-base-info {
  max-width: 1600px;
  width: 79%;
  display: flex;
  flex-flow: row;
  margin-top: 70px;
  align-items: center;
  justify-content: center;

  h1 {
    font-weight: normal;
    font-size: 35px;
    padding-right: 15px;
    border-right: 1px solid black;
    line-height: 1.4;
    letter-spacing: 0.02em;
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease;

    b {
      background: linear-gradient(90deg, #00c9ff 0%, #92fe9d 100%);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      font-weight: 700;
    }
  }

  div {
    display: flex;
    flex-flow: column wrap;
    gap: 10px;
  }

  p {
    padding-left: 10px;
    font-size: 18px;
    text-align: justify;
    transition: transform 0.3s ease;

    &:hover {
      transform: translateX(5px);
    }
  }
}

.devschool-base-info-text {
  background: linear-gradient(
    90deg,
    rgb(191, 255, 249) 0%,
    rgba(161, 255, 211, 1) 33%,
    rgb(209, 255, 208) 100%,
    rgb(195, 238, 255) 57%
  );
}

@media (max-width: 1350px) {
  .devschool-base-info {
    width: 75%;
    margin-top: 50px;

    h1 {
      font-size: 26px;
    }

    p {
      font-size: 18px;
    }
  }
}

@media (max-width: 830px) {
  .devschool-base-info {
    width: 80%;
    margin-top: 30px;

    h1 {
      font-size: 23px;
    }

    p {
      font-size: 15px;
    }
  }
}

@media (max-width: 770px) {
  .tab {
    height: 500px;
  }

  .buttons {
    flex-flow: column wrap;
    button p {
      font-size: 12px;
    }
  }

  .devschool-base-info {
    width: 95%;
    margin-top: 20px;
    flex-flow: column;

    h1 {
      border-right: none;
      border-bottom: 1px solid black;
      padding-bottom: 15px;
      padding-right: 10px;
      font-size: 18px;
    }

    p {
      margin-top: 10px;
      padding-left: 0px;
    }
  }
}

@media (max-width: 500px) {
  .tab {
    height: 440px;
  }
}
</style>
