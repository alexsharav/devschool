<template>
  <!--сделать работу с js(отправка сообщения)-->
  <section class="help-panel">
    <div class="contact-information">
      <h1 class="contact-information-text">Информация и контакты</h1>

      <div class="our-information-office info-item">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 576 512"
          fill="currentColor"
        >
          <path
            d="M288 0c-79.5 0-144 64.5-144 144 0 112 144 256 144 256s144-144 144-256c0-79.5-64.5-144-144-144zm0 192c-26.5 0-48-21.5-48-48s21.5-48  66-24 21.5 48 48-21.5 48-48 48zM560 464h-64v-48c0-8.8-7.2-16-16-16h-48v-64c0-8.8-7.2-16-16-16h-64v-48c0-8.8-7.2-16-16-16h-32v144c0 8.8 7.2 16 16 16h160v48c0 8.8 7.2 16 16 16h80c8.8 0 16-7.2 16-16v-80c0-8.8-7.2-16-16-16h-16v-64c0-8.8-7.2-16-16-16h-48v-48c0-8.8-7.2-16-16-16h-32v192c0 8.8 7.2 16 16 16h96v64c0 8.8-7.2 16-16 16z"
          />
        </svg>

        <div class="info-text">
          <h1>Наш офис</h1>

          <p>667000, г. Кызыл, ул. Кочетова, д.53, помещение 3.</p>
        </div>
      </div>

      <div class="our-information-phone info-item">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 512 512"
          fill="currentColor"
        >
          <path
            d="M511.1 380.8c1.1 18.1-4.7 35.8-16.1 49.1l-56.1 64.1c-13.6 15.5-34.5 21.4-53.6 15.2-59.4-19.5-121.2-54.7-185.1-114.4S60.7 232.3 41.2 172.9c-6.2-19.1-.3-40 15.2-53.6L120.6 63.2c13.3-11.3 31-17.2 49.1-16.1 13.3.8 25.2 7.6 32.2 18.3l47.2 72.6c9.3 14.3 7.4 33-4.4 45.3l-31.4 32.8c27.7 49.6 69.5 91.3 119.1 119l32.8-31.4c12.3-11.8 31-13.7 45.3-4.4l72.6 47.2c10.7 7 17.5 18.9 18.3 32.3z"
          />
        </svg>

        <div class="info-text">
          <h1>Номер телефона</h1>

          <p>8-923-267-48-88</p>
        </div>
      </div>

      <div class="our-information-education info-item">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 640 512"
          fill="currentColor"
        >
          <path
            d="M320 32L0 176l320 144 192-86.4v80.6c0 10.7-6.7 20.2-16.6 23.7L320 480 80 368V288H32v96c0 18.4 11 35.1 27.9 42.4l256 112c5.6 2.4 11.8 3.6 18.1 3.6s12.5-1.2 18.1-3.6l256-112c16.9-7.4 27.9-24 27.9-42.4V233.6l32-14.4V176L320 32z"
          />
        </svg>

        <div class="info-text">
          <h1>Образование</h1>

          <p>
            Профессиональное обучение<br />
            Адаптированные образовательные программы не реализуются<br />
            Форма обучения – очная, заочная<br />
            Язык образования – русский.<br />
          </p>
        </div>
      </div>
    </div>

    <div class="send-email">
      <h1 class="send-email-text">Напиши нам письмо</h1>

      <form class="email-form" @submit.prevent="helpPanelMessage">
        <div class="name-fields">
          <input type="text" v-model="userName" placeholder="Имя" />
          <input type="text" v-model="userSecondName" placeholder="Фамилия" />
        </div>

        <input type="email" v-model="userEmail" placeholder="Email*" required />

        <textarea
          class="resize-textarea"
          placeholder="Сообщение"
          rows="6"
          v-model="userMessage"
        ></textarea>

        <button type="submit" class="submit-button">Отправить</button>
      </form>
    </div>
  </section>
</template>

<script setup>
import { ref } from "vue";

const userName = ref("");
const userSecondName = ref("");
const userEmail = ref("");
const userMessage = ref("");

async function helpPanelMessage() {}
</script>

<style scoped>
.help-panel {
  border: 1px solid rgba(167, 167, 167, 0.705);
  border-radius: 18px;
  background: rgb(250, 250, 250);
  width: 98%;
  max-width: 1200px;
  height: auto;
  display: flex;
  flex-flow: row wrap;
}

.contact-information {
  padding: 30px 0px 0px 0px;
  border-radius: 18px;
  background: white;
  display: flex;
  flex-flow: column wrap;
  align-items: center;
  width: 35%;
  border-right: 1px solid rgba(167, 167, 167, 0.705);
}

.contact-information-text,
.send-email-text {
  font-weight: normal;
  font-size: 24px;
  color: rgb(59, 59, 59);
}

.send-email {
  padding: 30px 0px 0px 0px;
  display: flex;
  flex-flow: column wrap;
  align-items: center;
  width: 65%;
}

.info-item {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  padding: 20px 0px 0px 0px;
  width: 90%;

  svg {
    width: 30px;
    height: 30px;
    flex-shrink: 0;
  }
}

.info-text {
  display: flex;
  flex-direction: column;
  gap: 4px;

  h1 {
    font-size: 20px;
    margin: 0;
    font-weight: normal;
    color: rgb(77, 77, 77);
  }

  p {
    font-size: 16px;
    margin: 0;
    line-height: 1.4;
    color: gray;
  }
}

.our-information-education {
  margin: 0px 0px 30px 0px;
}

.email-form {
  width: 95%;
  display: flex;
  flex-flow: column;
  gap: 15px;
  margin: 20px 0px;

  label {
    font-size: 14px;
    font-weight: 500;
    color: #333;
  }

  input,
  textarea {
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    font-size: 14px;
  }
}

.name-fields {
  display: flex;
  flex-flow: wrap;
  gap: 10px;

  input {
    flex: 1;
  }
}

.resize-textarea {
  resize: none;
}

.submit-button {
  background-color: #3b3b3b;
  color: white;
  padding: 12px;
  font-size: 16px;
  border: none;
  border-radius: 8px;
  cursor: pointer;

  &:hover {
    background-color: #525252;
  }
}

@media (max-width: 630px) {
  .help-panel {
    flex-flow: column-reverse;
  }

  .contact-information {
    border-top: 1px solid rgba(167, 167, 167, 0.705);
    border-right: none;
  }

  .contact-information,
  .send-email {
    width: 100%;
  }
}

@media (max-width: 930px) {
  .contact-information-text,
  .send-email-text {
    font-size: 20px;
  }

  .info-item {
    img {
      width: 25px;
      height: 25px;
    }

    h1 {
      font-size: 17px;
    }

    p {
      font-size: 14px;
    }
  }
}

@media (max-width: 700px) {
  .contact-information-text,
  .send-email-text {
    font-size: 18px;
  }

  .info-item {
    img {
      width: 25px;
      height: 25px;
    }

    h1 {
      font-size: 15px;
    }

    p {
      font-size: 12px;
    }
  }
}
</style>
