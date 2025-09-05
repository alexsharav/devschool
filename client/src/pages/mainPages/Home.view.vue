<template>
  <Header />

  <StarterTab
    mainText="Дев.Школа"
    secondaryText="образовательные курсы и тесты для всех"
  >
    <template #buttons>
      <div class="buttons">
        <button @click="router.push('/courses')">
          <p>начни с наших курсов</p>
        </button>

        <button @click="router.push('/tests')">
          <p>начни с наших тестов</p>
        </button>
      </div>
    </template>
  </StarterTab>

  <section class="devschool-base-info">
    <h1c>Платформа для тех</h1c>
    <h1c>кто хочет войти в IT</h1c>

    <article class="devschool-pluses-wrapper">
      <div>
        <div class="div-tags">
          <Chip>Экономия времени</Chip>
          <Chip>Гибкость учебы</Chip>
          <Chip>Удобство и комфорт</Chip>
        </div>

        <div class="div-pluses-main">
          <svg
            class="w-6 h-6 text-gray-800 dark:text-white"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              fill-rule="evenodd"
              d="M6 5V4a1 1 0 1 1 2 0v1h3V4a1 1 0 1 1 2 0v1h3V4a1 1 0 1 1 2 0v1h1a2 2 0 0 1 2 2v2H3V7a2 2 0 0 1 2-2h1ZM3 19v-8h18v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Zm5-6a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2H8Z"
              clip-rule="evenodd"
            />
          </svg>

          <div class="div-pluses">
            <h1>Гибкий темп обучения</h1>
            <p>
              Учиться можно в любое удобное время. Возможность ускорять и
              замедлять образовательный процесс.
            </p>
          </div>
        </div>
      </div>

      <div>
        <div class="div-tags">
          <Chip>Начало</Chip>
          <Chip>Практика во время учёбы</Chip>
          <Chip>Быстрый рост навыков</Chip>
        </div>

        <div class="div-pluses-main">
          <svg
            class="w-6 h-6 text-gray-800 dark:text-white"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              fill-rule="evenodd"
              d="M6 2a2 2 0 0 0-2 2v15a3 3 0 0 0 3 3h12a1 1 0 1 0 0-2h-2v-2h2a1 1 0 0 0 1-1V4a2 2 0 0 0-2-2h-8v16h5v2H7a1 1 0 1 1 0-2h1V2H6Z"
              clip-rule="evenodd"
            />
          </svg>

          <div class="div-pluses">
            <h1>Подача материала</h1>
            <p>
              Понятная подача материала вместе с закреплением прочитанного в
              виде интерактивных тестов.
            </p>
          </div>
        </div>
      </div>

      <div>
        <div class="div-tags">
          <Chip>Профессия</Chip>
          <Chip>Стандарт</Chip>
          <Chip>Обновление</Chip>
        </div>

        <div class="div-pluses-main">
          <svg
            class="w-6 h-6 text-gray-800 dark:text-white"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="none"
            viewBox="0 0 24 24"
          >
            <path
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M17.651 7.65a7.131 7.131 0 0 0-12.68 3.15M18.001 4v4h-4m-7.652 8.35a7.13 7.13 0 0 0 12.68-3.15M6 20v-4h4"
            />
          </svg>

          <div class="div-pluses">
            <h1>Современные курсы</h1>
            <p>
              Новейшие курсы, затрагивающие последние тренды, обновления и
              стандарты.
            </p>
          </div>
        </div>
      </div>

      <div>
        <div class="div-tags">
          <Chip>Будущее</Chip>
          <Chip>Навыки</Chip>
          <Chip>IT</Chip>
        </div>

        <div class="div-pluses-main">
          <svg
            class="w-6 h-6 text-gray-800 dark:text-white"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="none"
            viewBox="0 0 24 24"
          >
            <path
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 6v13m0-13 4 4m-4-4-4 4"
            />
          </svg>

          <div class="div-pluses">
            <h1>Карьерное будущее</h1>

            <p>
              Курсы помогают построить свою новую карьеру. Освоить разные навыки
              в сфере IT.
            </p>
          </div>
        </div>
      </div>
    </article>
  </section>

  <section class="to-media-section">
    <div>
      <h1c>Школа не только про обучение</h1c>
      <router-link to="/media">Смотри наше медиа</router-link>
    </div>
  </section>

  <section class="help-panel-section">
    <h1c>Рады помочь нашим студентам</h1c>
    <h1c>Готовы принять критику в свою сторону</h1c>

    <div class="help-panel-wrapper">
      <HelpPanel />
    </div>
  </section>

  <Footer />
</template>

<script setup>
import Header from "@/ui/main/Header.vue";
import Footer from "@/ui/main/Footer.vue";
import StarterTab from "@/ui/main/StarterTab.vue";
import router from "@/router/router";
import h1c from "@/ui/small/H1c.vue";
import Chip from "@/ui/small/Chip.vue";
import HelpPanel from "@/components/otherComponents/HelpPanel.vue";
</script>

<style scoped>
.course-features {
  margin-top: 20vh;
}

.tab {
  height: 500px;

  @media (max-width: 770px) {
    height: 400px;
  }

  @media (max-width: 500px) {
    height: 440px;
  }
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
    border: 1px solid black;
    transition: 0.3s ease-in-out;
    background: rgba(59, 59, 59, 0);

    &:hover {
      background: rgba(59, 59, 59, 0.15);
      transform: scale(1.1);
    }

    p {
      font-size: 16px;
      font-weight: bold;
      padding: 0px 15px;
      color: black;
      transition: 0.3s ease-in-out;
    }
  }

  @media (max-width: 770px) {
    flex-flow: wrap;

    button {
      p {
        font-size: 12px;
      }
    }
  }

  @media (max-width: 500px) {
    flex-flow: column wrap;
  }
}

.devschool-base-info {
  max-width: 1600px;
  width: 80%;
  margin-top: 100px;

  @media (max-width: 1100px) {
    width: 95%;
  }

  @media (max-width: 930px) {
    width: 98%;
  }

  @media (max-width: 770px) {
    flex-flow: column;
  }
}

.devschool-pluses-wrapper {
  margin-top: 40px;
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 14px;

  > div {
    background: linear-gradient(90deg, hsl(81, 100%, 87%) 0%, #fff6cc 100%);
    padding: 40px;
    border-radius: 20px;
    border: 1px solid rgb(228, 227, 179);
    box-sizing: border-box;

    .div-tags {
      display: flex;
      gap: 10px;
      width: 100%;
    }

    .div-pluses-main {
      display: flex;
      flex-direction: row;
      column-gap: 20px;
      text-align: left;
      margin-top: 20px;

      h1 {
        font-weight: normal;
      }

      p {
        margin-top: 14px;
        font-size: 20px;
      }
    }

    > div > div {
      flex: 1;
    }
  }

  svg {
    width: 40px;
    height: 40px;
  }

  @media (max-width: 1100px) {
    grid-template-columns: 100%;
  }

  @media (max-width: 600px) {
    margin-top: 24px;

    > div {
      padding: 20px 16px;
      border-radius: 14px;

      .div-tags {
        flex-wrap: wrap;
      }

      .div-pluses-main {
        column-gap: 12px;
        margin-top: 16px;
        align-items: flex-start;

        > div {
          flex: 1 1 auto;
          min-width: 0;
        }

        h1 {
          font-size: 18px;
          line-height: 1.2;
          margin: 0;
        }

        p {
          font-size: 14px;
          line-height: 1.35;
          margin-top: 8px;
        }
      }
    }

    svg {
      width: 28px;
      height: 28px;
      margin-top: 18px;
      flex-shrink: 0;
    }
  }

  @media (max-width: 420px) {
    > div {
      padding: 16px 14px;

      .div-pluses {
        h1 {
          font-size: 16px;
        }

        p {
          font-size: 13px;
        }
      }
    }

    svg {
      width: 24px;
      height: 24px;
    }
  }
}

.help-panel-section {
  margin-top: 70px;
  width: 80%;
  padding: 25px;
  border: 1px solid black;
  border-radius: 20px;
  background-color: #ffffff;
  max-width: 1600px;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='192' height='192' viewBox='0 0 192 192'%3E%3Cpath fill='%23000000' fill-opacity='0.1' d='M192 15v2a11 11 0 0 0-11 11c0 1.94 1.16 4.75 2.53 6.11l2.36 2.36a6.93 6.93 0 0 1 1.22 7.56l-.43.84a8.08 8.08 0 0 1-6.66 4.13H145v35.02a6.1 6.1 0 0 0 3.03 4.87l.84.43c1.58.79 4 .4 5.24-.85l2.36-2.36a12.04 12.04 0 0 1 7.51-3.11 13 13 0 1 1 .02 26 12 12 0 0 1-7.53-3.11l-2.36-2.36a4.93 4.93 0 0 0-5.24-.85l-.84.43a6.1 6.1 0 0 0-3.03 4.87V143h35.02a8.08 8.08 0 0 1 6.66 4.13l.43.84a6.91 6.91 0 0 1-1.22 7.56l-2.36 2.36A10.06 10.06 0 0 0 181 164a11 11 0 0 0 11 11v2a13 13 0 0 1-13-13 12 12 0 0 1 3.11-7.53l2.36-2.36a4.93 4.93 0 0 0 .85-5.24l-.43-.84a6.1 6.1 0 0 0-4.87-3.03H145v35.02a8.08 8.08 0 0 1-4.13 6.66l-.84.43a6.91 6.91 0 0 1-7.56-1.22l-2.36-2.36A10.06 10.06 0 0 0 124 181a11 11 0 0 0-11 11h-2a13 13 0 0 1 13-13c2.47 0 5.79 1.37 7.53 3.11l2.36 2.36a4.94 4.94 0 0 0 5.24.85l.84-.43a6.1 6.1 0 0 0 3.03-4.87V145h-35.02a8.08 8.08 0 0 1-6.66-4.13l-.43-.84a6.91 6.91 0 0 1 1.22-7.56l2.36-2.36A10.06 10.06 0 0 0 107 124a11 11 0 0 0-22 0c0 1.94 1.16 4.75 2.53 6.11l2.36 2.36a6.93 6.93 0 0 1 1.22 7.56l-.43.84a8.08 8.08 0 0 1-6.66 4.13H49v35.02a6.1 6.1 0 0 0 3.03 4.87l.84.43c1.58.79 4 .4 5.24-.85l2.36-2.36a12.04 12.04 0 0 1 7.51-3.11A13 13 0 0 1 81 192h-2a11 11 0 0 0-11-11c-1.94 0-4.75 1.16-6.11 2.53l-2.36 2.36a6.93 6.93 0 0 1-7.56 1.22l-.84-.43a8.08 8.08 0 0 1-4.13-6.66V145H11.98a6.1 6.1 0 0 0-4.87 3.03l-.43.84c-.79 1.58-.4 4 .85 5.24l2.36 2.36a12.04 12.04 0 0 1 3.11 7.51A13 13 0 0 1 0 177v-2a11 11 0 0 0 11-11c0-1.94-1.16-4.75-2.53-6.11l-2.36-2.36a6.93 6.93 0 0 1-1.22-7.56l.43-.84a8.08 8.08 0 0 1 6.66-4.13H47v-35.02a6.1 6.1 0 0 0-3.03-4.87l-.84-.43c-1.59-.8-4-.4-5.24.85l-2.36 2.36A12 12 0 0 1 28 109a13 13 0 1 1 0-26c2.47 0 5.79 1.37 7.53 3.11l2.36 2.36a4.94 4.94 0 0 0 5.24.85l.84-.43A6.1 6.1 0 0 0 47 84.02V49H11.98a8.08 8.08 0 0 1-6.66-4.13l-.43-.84a6.91 6.91 0 0 1 1.22-7.56l2.36-2.36A10.06 10.06 0 0 0 11 28 11 11 0 0 0 0 17v-2a13 13 0 0 1 13 13c0 2.47-1.37 5.79-3.11 7.53l-2.36 2.36a4.94 4.94 0 0 0-.85 5.24l.43.84A6.1 6.1 0 0 0 11.98 47H47V11.98a8.08 8.08 0 0 1 4.13-6.66l.84-.43a6.91 6.91 0 0 1 7.56 1.22l2.36 2.36A10.06 10.06 0 0 0 68 11 11 11 0 0 0 79 0h2a13 13 0 0 1-13 13 12 12 0 0 1-7.53-3.11l-2.36-2.36a4.93 4.93 0 0 0-5.24-.85l-.84.43A6.1 6.1 0 0 0 49 11.98V47h35.02a8.08 8.08 0 0 1 6.66 4.13l.43.84a6.91 6.91 0 0 1-1.22 7.56l-2.36 2.36A10.06 10.06 0 0 0 85 68a11 11 0 0 0 22 0c0-1.94-1.16-4.75-2.53-6.11l-2.36-2.36a6.93 6.93 0 0 1-1.22-7.56l.43-.84a8.08 8.08 0 0 1 6.66-4.13H143V11.98a6.1 6.1 0 0 0-3.03-4.87l-.84-.43c-1.59-.8-4-.4-5.24.85l-2.36 2.36A12 12 0 0 1 124 13a13 13 0 0 1-13-13h2a11 11 0 0 0 11 11c1.94 0 4.75-1.16 6.11-2.53l2.36-2.36a6.93 6.93 0 0 1 7.56-1.22l.84.43a8.08 8.08 0 0 1 4.13 6.66V47h35.02a6.1 6.1 0 0 0 4.87-3.03l.43-.84c.8-1.59.4-4-.85-5.24l-2.36-2.36A12 12 0 0 1 179 28a13 13 0 0 1 13-13zM84.02 143a6.1 6.1 0 0 0 4.87-3.03l.43-.84c.8-1.59.4-4-.85-5.24l-2.36-2.36A12 12 0 0 1 83 124a13 13 0 1 1 26 0c0 2.47-1.37 5.79-3.11 7.53l-2.36 2.36a4.94 4.94 0 0 0-.85 5.24l.43.84a6.1 6.1 0 0 0 4.87 3.03H143v-35.02a8.08 8.08 0 0 1 4.13-6.66l.84-.43a6.91 6.91 0 0 1 7.56 1.22l2.36 2.36A10.06 10.06 0 0 0 164 107a11 11 0 0 0 0-22c-1.94 0-4.75 1.16-6.11 2.53l-2.36 2.36a6.93 6.93 0 0 1-7.56 1.22l-.84-.43a8.08 8.08 0 0 1-4.13-6.66V49h-35.02a6.1 6.1 0 0 0-4.87 3.03l-.43.84c-.79 1.58-.4 4 .85 5.24l2.36 2.36a12.04 12.04 0 0 1 3.11 7.51A13 13 0 1 1 83 68a12 12 0 0 1 3.11-7.53l2.36-2.36a4.93 4.93 0 0 0 .85-5.24l-.43-.84A6.1 6.1 0 0 0 84.02 49H49v35.02a8.08 8.08 0 0 1-4.13 6.66l-.84.43a6.91 6.91 0 0 1-7.56-1.22l-2.36-2.36A10.06 10.06 0 0 0 28 85a11 11 0 0 0 0 22c1.94 0 4.75-1.16 6.11-2.53l2.36-2.36a6.93 6.93 0 0 1 7.56-1.22l.84.43a8.08 8.08 0 0 1 4.13 6.66V143h35.02z'%3E%3C/path%3E%3C/svg%3E");

  .help-panel-wrapper {
    width: 100%;
    margin-top: 50px;
    display: flex;
    justify-content: center;
  }

  @media (max-width: 1100px) {
    width: 95%;
  }

  @media (max-width: 930px) {
    width: 98%;
  }
}

.to-media-section {
  width: 80%;
  margin-top: 70px;
  padding: 20px;
  border-radius: 20px;
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
  max-width: 1600px;
  background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' version='1.1' xmlns:xlink='http://www.w3.org/1999/xlink' xmlns:svgjs='http://svgjs.dev/svgjs' width='1440' height='560' preserveAspectRatio='none' viewBox='0 0 1440 560'%3e%3cg mask='url(%26quot%3b%23SvgjsMask1000%26quot%3b)' fill='none'%3e%3crect width='1440' height='560' x='0' y='0' fill='rgba(30%2c 30%2c 30%2c 1)'%3e%3c/rect%3e%3cpath d='M-94.3 398.88L-94.3 398.88' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-94.3 398.88L-94.97 505.95' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-94.3 398.88L58.42 360.58' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-94.3 398.88L66.16 535.6' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-94.3 398.88L-59.59 693.61' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-94.97 505.95L-94.97 505.95' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-94.97 505.95L66.16 535.6' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-94.97 505.95L-59.59 693.61' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-59.59 693.61L-59.59 693.61' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-59.59 693.61L96.88 641.44' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-59.59 693.61L66.16 535.6' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-59.59 693.61L216.41 647.89' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M58.42 360.58L58.42 360.58' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M58.42 360.58L198.83 349.58' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M58.42 360.58L66.16 535.6' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M66.16 535.6L66.16 535.6' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M66.16 535.6L96.88 641.44' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M66.16 535.6L199.92 495.4' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M66.16 535.6L216.41 647.89' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M96.88 641.44L96.88 641.44' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M96.88 641.44L216.41 647.89' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M198.83 349.58L198.83 349.58' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M198.83 349.58L199.92 495.4' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M199.92 495.4L199.92 495.4' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M216.41 647.89L216.41 647.89' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M216.41 647.89L199.92 495.4' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M216.41 647.89L392.99 671.54' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M363.67 396.74L363.67 396.74' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M363.67 396.74L396.41 511.55' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M363.67 396.74L502.32 366.33' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M363.67 396.74L198.83 349.58' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M363.67 396.74L199.92 495.4' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M396.41 511.55L396.41 511.55' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M396.41 511.55L499.33 548.95' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M396.41 511.55L392.99 671.54' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M396.41 511.55L498.89 653.59' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M392.99 671.54L392.99 671.54' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M392.99 671.54L498.89 653.59' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M392.99 671.54L499.33 548.95' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M502.32 366.33L502.32 366.33' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M502.32 366.33L396.41 511.55' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M502.32 366.33L499.33 548.95' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M502.32 366.33L693.26 359.37' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M502.32 366.33L709.54 261.24' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M499.33 548.95L499.33 548.95' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M499.33 548.95L498.89 653.59' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M499.33 548.95L674.66 534.68' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M499.33 548.95L363.67 396.74' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M498.89 653.59L498.89 653.59' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M498.89 653.59L710.5 653.71' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M498.89 653.59L674.66 534.68' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M498.89 653.59L216.41 647.89' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M709.54 261.24L709.54 261.24' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M709.54 261.24L693.26 359.37' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M709.54 261.24L829.72 349.8' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M709.54 261.24L859.15 194.97' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M709.54 261.24L852.35 108.04' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M693.26 359.37L693.26 359.37' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M693.26 359.37L829.72 349.8' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M693.26 359.37L674.66 534.68' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M693.26 359.37L828.59 489.02' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M674.66 534.68L674.66 534.68' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M674.66 534.68L710.5 653.71' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M674.66 534.68L828.59 489.02' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M674.66 534.68L810.94 687.53' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M710.5 653.71L710.5 653.71' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M710.5 653.71L810.94 687.53' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M710.5 653.71L828.59 489.02' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M852.35 108.04L852.35 108.04' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M852.35 108.04L859.15 194.97' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M852.35 108.04L982.5 39.64' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M852.35 108.04L982.41 221.71' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M859.15 194.97L859.15 194.97' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M859.15 194.97L982.41 221.71' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M859.15 194.97L829.72 349.8' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M859.15 194.97L982.5 39.64' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M859.15 194.97L996.49 383.16' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M829.72 349.8L829.72 349.8' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M829.72 349.8L828.59 489.02' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M829.72 349.8L996.49 383.16' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M828.59 489.02L828.59 489.02' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M828.59 489.02L993.74 533.92' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M810.94 687.53L810.94 687.53' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M810.94 687.53L964.74 673.53' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M947.13 -92.98L947.13 -92.98' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M947.13 -92.98L982.5 39.64' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M947.13 -92.98L1142.81 -88.6' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M947.13 -92.98L852.35 108.04' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M947.13 -92.98L1115.13 100.63' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M947.13 -92.98L859.15 194.97' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M947.13 -92.98L982.41 221.71' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M982.5 39.64L982.5 39.64' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M982.5 39.64L1115.13 100.63' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M982.5 39.64L982.41 221.71' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M982.5 39.64L1142.81 -88.6' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M982.41 221.71L982.41 221.71' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M982.41 221.71L1139.83 192.48' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M996.49 383.16L996.49 383.16' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M996.49 383.16L1119.43 346.41' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M996.49 383.16L993.74 533.92' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M996.49 383.16L982.41 221.71' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M996.49 383.16L1135.21 501.43' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M996.49 383.16L828.59 489.02' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M993.74 533.92L993.74 533.92' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M993.74 533.92L964.74 673.53' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M993.74 533.92L1135.21 501.43' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M964.74 673.53L964.74 673.53' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M964.74 673.53L1126.92 677.56' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M964.74 673.53L828.59 489.02' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1142.81 -88.6L1142.81 -88.6' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1142.81 -88.6L1271.33 -72.89' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1142.81 -88.6L1238.26 70.81' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1115.13 100.63L1115.13 100.63' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1115.13 100.63L1139.83 192.48' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1115.13 100.63L1238.26 70.81' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1115.13 100.63L982.41 221.71' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1115.13 100.63L1280.62 187.58' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1115.13 100.63L1142.81 -88.6' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1139.83 192.48L1139.83 192.48' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1139.83 192.48L1280.62 187.58' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1139.83 192.48L1119.43 346.41' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1139.83 192.48L1238.26 70.81' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1139.83 192.48L1258.95 356.68' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1119.43 346.41L1119.43 346.41' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1119.43 346.41L1258.95 356.68' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1119.43 346.41L1135.21 501.43' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1119.43 346.41L982.41 221.71' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1135.21 501.43L1135.21 501.43' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1135.21 501.43L1271.62 523.49' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1135.21 501.43L1126.92 677.56' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1126.92 677.56L1126.92 677.56' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1126.92 677.56L1310.81 679.95' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1126.92 677.56L993.74 533.92' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1126.92 677.56L1271.62 523.49' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1126.92 677.56L1402.09 668.26' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1271.33 -72.89L1271.33 -72.89' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1271.33 -72.89L1238.26 70.81' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1271.33 -72.89L1389.11 98.67' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1271.33 -72.89L1115.13 100.63' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1238.26 70.81L1238.26 70.81' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1238.26 70.81L1280.62 187.58' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1280.62 187.58L1280.62 187.58' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1280.62 187.58L1389.11 98.67' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1280.62 187.58L1425.51 261.07' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1258.95 356.68L1258.95 356.68' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1258.95 356.68L1271.62 523.49' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1271.62 523.49L1271.62 523.49' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1271.62 523.49L1389.92 505.03' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1271.62 523.49L1310.81 679.95' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1271.62 523.49L1402.09 668.26' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1310.81 679.95L1310.81 679.95' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1310.81 679.95L1402.09 668.26' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1310.81 679.95L1389.92 505.03' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1310.81 679.95L1135.21 501.43' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1389.11 98.67L1389.11 98.67' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1389.11 98.67L1238.26 70.81' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1389.11 98.67L1425.51 261.07' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1389.11 98.67L1590.04 64.02' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1389.11 98.67L1610.37 191.15' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1425.51 261.07L1425.51 261.07' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1425.51 261.07L1442.89 377.56' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1442.89 377.56L1442.89 377.56' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1442.89 377.56L1389.92 505.03' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1442.89 377.56L1610.44 383.41' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1442.89 377.56L1548.81 508.11' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1442.89 377.56L1258.95 356.68' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1442.89 377.56L1271.62 523.49' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1389.92 505.03L1389.92 505.03' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1389.92 505.03L1548.81 508.11' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1389.92 505.03L1402.09 668.26' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1389.92 505.03L1258.95 356.68' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1402.09 668.26L1402.09 668.26' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1402.09 668.26L1564.44 674.98' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1604.8 -104.5L1604.8 -104.5' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1604.8 -104.5L1590.04 64.02' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1604.8 -104.5L1610.37 191.15' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1604.8 -104.5L1389.11 98.67' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1590.04 64.02L1590.04 64.02' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1590.04 64.02L1610.37 191.15' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1590.04 64.02L1425.51 261.07' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1610.37 191.15L1610.37 191.15' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1610.37 191.15L1610.44 383.41' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1610.37 191.15L1425.51 261.07' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1610.37 191.15L1442.89 377.56' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1610.44 383.41L1610.44 383.41' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1610.44 383.41L1548.81 508.11' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1610.44 383.41L1425.51 261.07' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1610.44 383.41L1389.92 505.03' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1548.81 508.11L1548.81 508.11' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1548.81 508.11L1564.44 674.98' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1548.81 508.11L1402.09 668.26' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1564.44 674.98L1564.44 674.98' stroke='rgba(49%2c 33%2c 121%2c 1)' stroke-width='1.5'%3e%3c/path%3e%3ccircle r='5' cx='-94.3' cy='398.88' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='-94.97' cy='505.95' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='-59.59' cy='693.61' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='58.42' cy='360.58' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='66.16' cy='535.6' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='96.88' cy='641.44' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='198.83' cy='349.58' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='199.92' cy='495.4' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='216.41' cy='647.89' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='363.67' cy='396.74' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='396.41' cy='511.55' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='392.99' cy='671.54' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='502.32' cy='366.33' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='499.33' cy='548.95' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='498.89' cy='653.59' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='709.54' cy='261.24' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='693.26' cy='359.37' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='674.66' cy='534.68' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='710.5' cy='653.71' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='852.35' cy='108.04' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='859.15' cy='194.97' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='829.72' cy='349.8' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='828.59' cy='489.02' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='810.94' cy='687.53' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='947.13' cy='-92.98' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='982.5' cy='39.64' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='982.41' cy='221.71' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='996.49' cy='383.16' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='993.74' cy='533.92' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='964.74' cy='673.53' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1142.81' cy='-88.6' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1115.13' cy='100.63' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1139.83' cy='192.48' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1119.43' cy='346.41' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1135.21' cy='501.43' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1126.92' cy='677.56' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1271.33' cy='-72.89' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1238.26' cy='70.81' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1280.62' cy='187.58' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1258.95' cy='356.68' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1271.62' cy='523.49' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1310.81' cy='679.95' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1389.11' cy='98.67' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1425.51' cy='261.07' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1442.89' cy='377.56' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1389.92' cy='505.03' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1402.09' cy='668.26' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1604.8' cy='-104.5' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1590.04' cy='64.02' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1610.37' cy='191.15' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1610.44' cy='383.41' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1548.81' cy='508.11' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3ccircle r='5' cx='1564.44' cy='674.98' fill='rgba(49%2c 33%2c 121%2c 1)'%3e%3c/circle%3e%3cpath d='M-82.45 513.18L-82.45 513.18' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-82.45 513.18L38.78 525.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-82.45 513.18L-79.77 696.09' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M405.46 521.69L405.46 521.69' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M405.46 521.69L539.4 532.72' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M405.46 521.69L519.34 648.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M405.46 521.69L360.87 712.36' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M405.46 521.69L640.3 678.15' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M405.46 521.69L700.08 527.12' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M405.46 521.69L701.35 405.03' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M700.08 527.12L700.08 527.12' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M700.08 527.12L701.35 405.03' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M700.08 527.12L846.23 534.81' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M700.08 527.12L539.4 532.72' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M700.08 527.12L804.71 403.15' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M700.08 527.12L640.3 678.15' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M839.67 706.42L839.67 706.42' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M839.67 706.42L987.3 690.11' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M839.67 706.42L846.23 534.81' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M839.67 706.42L973.68 558' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M839.67 706.42L640.3 678.15' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M839.67 706.42L700.08 527.12' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M839.67 706.42L1121.2 699.46' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M987.3 690.11L987.3 690.11' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M987.3 690.11L973.68 558' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M987.3 690.11L1121.2 699.46' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1416.38 106.66L1416.38 106.66' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1416.38 106.66L1416.66 195.33' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1416.38 106.66L1432.23 -58.85' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1416.38 106.66L1253.32 58.73' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1416.38 106.66L1594.53 61.9' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1416.38 106.66L1280.04 238.25' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1416.38 106.66L1589.13 194.59' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1594.53 61.9L1594.53 61.9' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1594.53 61.9L1589.13 194.59' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1594.53 61.9L1557.74 -97.21' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1594.53 61.9L1432.23 -58.85' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1589.13 194.59L1589.13 194.59' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1589.13 194.59L1580.9 353.57' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1609.16 524.18L1609.16 524.18' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1609.16 524.18L1549.97 672.93' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1609.16 524.18L1580.9 353.57' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1609.16 524.18L1412.87 524.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1609.16 524.18L1408.09 395.5' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-79.77 696.09L-79.77 696.09' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-79.77 696.09L63.73 701.24' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-79.77 696.09L38.78 525.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-79.77 696.09L360.87 712.36' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M-79.77 696.09L405.46 521.69' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M38.78 525.28L38.78 525.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M38.78 525.28L63.73 701.24' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M38.78 525.28L405.46 521.69' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M63.73 701.24L63.73 701.24' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M63.73 701.24L-82.45 513.18' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M63.73 701.24L360.87 712.36' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M63.73 701.24L405.46 521.69' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M63.73 701.24L519.34 648.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M360.87 712.36L360.87 712.36' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M360.87 712.36L519.34 648.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M539.4 532.72L539.4 532.72' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M539.4 532.72L519.34 648.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M539.4 532.72L640.3 678.15' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M539.4 532.72L701.35 405.03' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M519.34 648.28L519.34 648.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M519.34 648.28L640.3 678.15' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M701.35 405.03L701.35 405.03' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M701.35 405.03L804.71 403.15' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M701.35 405.03L846.23 534.81' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M701.35 405.03L965.49 362.5' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M701.35 405.03L640.3 678.15' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M640.3 678.15L640.3 678.15' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M804.71 403.15L804.71 403.15' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M804.71 403.15L846.23 534.81' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M846.23 534.81L846.23 534.81' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M846.23 534.81L973.68 558' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M982.7 262.02L982.7 262.02' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M982.7 262.02L965.49 362.5' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M982.7 262.02L1116 227.25' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M982.7 262.02L1103.24 339.45' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M982.7 262.02L804.71 403.15' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M965.49 362.5L965.49 362.5' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M965.49 362.5L1103.24 339.45' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M965.49 362.5L804.71 403.15' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M973.68 558L973.68 558' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M973.68 558L1096.97 532.84' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1116 227.25L1116 227.25' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1116 227.25L1103.24 339.45' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1103.24 339.45L1103.24 339.45' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1096.97 532.84L1096.97 532.84' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1096.97 532.84L1121.2 699.46' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1096.97 532.84L987.3 690.11' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1121.2 699.46L1121.2 699.46' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1121.2 699.46L1305.1 644.26' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1121.2 699.46L973.68 558' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1121.2 699.46L1295.57 502.6' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1237.94 -108.17L1237.94 -108.17' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1237.94 -108.17L1253.32 58.73' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1237.94 -108.17L1432.23 -58.85' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1237.94 -108.17L1416.38 106.66' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1237.94 -108.17L1557.74 -97.21' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1253.32 58.73L1253.32 58.73' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1253.32 58.73L1280.04 238.25' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1280.04 238.25L1280.04 238.25' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1280.04 238.25L1257.72 373.81' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1280.04 238.25L1416.66 195.33' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1280.04 238.25L1116 227.25' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1280.04 238.25L1408.09 395.5' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1257.72 373.81L1257.72 373.81' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1257.72 373.81L1295.57 502.6' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1257.72 373.81L1408.09 395.5' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1257.72 373.81L1103.24 339.45' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1257.72 373.81L1116 227.25' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1257.72 373.81L1412.87 524.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1295.57 502.6L1295.57 502.6' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1295.57 502.6L1412.87 524.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1295.57 502.6L1305.1 644.26' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1295.57 502.6L1408.09 395.5' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1295.57 502.6L1388.93 646.23' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1295.57 502.6L1096.97 532.84' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1305.1 644.26L1305.1 644.26' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1305.1 644.26L1388.93 646.23' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1305.1 644.26L1412.87 524.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1432.23 -58.85L1432.23 -58.85' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1432.23 -58.85L1557.74 -97.21' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1416.66 195.33L1416.66 195.33' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1416.66 195.33L1589.13 194.59' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1408.09 395.5L1408.09 395.5' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1408.09 395.5L1412.87 524.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1412.87 524.28L1412.87 524.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1412.87 524.28L1388.93 646.23' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1388.93 646.23L1388.93 646.23' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1388.93 646.23L1549.97 672.93' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1557.74 -97.21L1557.74 -97.21' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1557.74 -97.21L1416.38 106.66' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1557.74 -97.21L1589.13 194.59' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1580.9 353.57L1580.9 353.57' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1580.9 353.57L1408.09 395.5' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1580.9 353.57L1416.66 195.33' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1580.9 353.57L1412.87 524.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1580.9 353.57L1594.53 61.9' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1549.97 672.93L1549.97 672.93' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3cpath d='M1549.97 672.93L1412.87 524.28' stroke='hsl(263.20000000000005%2c 83.6%25%2c 52.8%25)' stroke-width='1.5'%3e%3c/path%3e%3ccircle r='25' cx='-82.45' cy='513.18' fill='url(%26quot%3b%23SvgjsRadialGradient1001%26quot%3b)'%3e%3c/circle%3e%3ccircle r='25' cx='405.46' cy='521.69' fill='url(%26quot%3b%23SvgjsRadialGradient1001%26quot%3b)'%3e%3c/circle%3e%3ccircle r='25' cx='700.08' cy='527.12' fill='url(%26quot%3b%23SvgjsRadialGradient1001%26quot%3b)'%3e%3c/circle%3e%3ccircle r='25' cx='839.67' cy='706.42' fill='url(%26quot%3b%23SvgjsRadialGradient1001%26quot%3b)'%3e%3c/circle%3e%3ccircle r='25' cx='987.3' cy='690.11' fill='url(%26quot%3b%23SvgjsRadialGradient1001%26quot%3b)'%3e%3c/circle%3e%3ccircle r='25' cx='1416.38' cy='106.66' fill='url(%26quot%3b%23SvgjsRadialGradient1001%26quot%3b)'%3e%3c/circle%3e%3ccircle r='25' cx='1594.53' cy='61.9' fill='url(%26quot%3b%23SvgjsRadialGradient1001%26quot%3b)'%3e%3c/circle%3e%3ccircle r='25' cx='1589.13' cy='194.59' fill='url(%26quot%3b%23SvgjsRadialGradient1001%26quot%3b)'%3e%3c/circle%3e%3ccircle r='25' cx='1609.16' cy='524.18' fill='url(%26quot%3b%23SvgjsRadialGradient1001%26quot%3b)'%3e%3c/circle%3e%3ccircle r='5' cx='-79.77' cy='696.09' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='38.78' cy='525.28' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='63.73' cy='701.24' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='360.87' cy='712.36' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='539.4' cy='532.72' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='519.34' cy='648.28' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='701.35' cy='405.03' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='640.3' cy='678.15' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='804.71' cy='403.15' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='846.23' cy='534.81' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='982.7' cy='262.02' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='965.49' cy='362.5' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='973.68' cy='558' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1116' cy='227.25' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1103.24' cy='339.45' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1096.97' cy='532.84' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1121.2' cy='699.46' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1237.94' cy='-108.17' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1253.32' cy='58.73' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1280.04' cy='238.25' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1257.72' cy='373.81' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1295.57' cy='502.6' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1305.1' cy='644.26' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1432.23' cy='-58.85' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1416.66' cy='195.33' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1408.09' cy='395.5' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1412.87' cy='524.28' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1388.93' cy='646.23' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1557.74' cy='-97.21' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1580.9' cy='353.57' fill='%23aa88df'%3e%3c/circle%3e%3ccircle r='5' cx='1549.97' cy='672.93' fill='%23aa88df'%3e%3c/circle%3e%3c/g%3e%3cdefs%3e%3cmask id='SvgjsMask1000'%3e%3crect width='1440' height='560' fill='white'%3e%3c/rect%3e%3c/mask%3e%3cradialGradient id='SvgjsRadialGradient1001'%3e%3cstop stop-color='white' offset='0.1'%3e%3c/stop%3e%3cstop stop-color='rgba(84%2c 17%2c 190%2c 1)' offset='0.2'%3e%3c/stop%3e%3cstop stop-color='rgba(84%2c 17%2c 190%2c 0)' offset='1'%3e%3c/stop%3e%3c/radialGradient%3e%3c/defs%3e%3c/svg%3e");

  div {
    height: 400px;
    display: flex;
    flex-flow: column wrap;
    align-items: center;
    justify-content: center;
    gap: 20px;

    h1 {
      color: white;
      border-radius: 20px;
      padding: 20px;
      text-align: center;
    }

    a {
      background: white;
      text-decoration: none;
      color: black;
      padding: 15px;
      border-radius: 14px;
      border: 1px solid white;
      transition: 0.5s ease;

      &:hover {
        background: rgba(207, 207, 207, 0);
        color: white;
        transform: scale(1.1);
      }
    }
  }

  @media (max-width: 1100px) {
    width: 95%;
    margin-top: 60px;
  }

  @media (max-width: 930px) {
    width: 98%;
    margin-top: 40px;
  }
}
</style>
