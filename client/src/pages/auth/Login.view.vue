<template>
  <div class="page-wrapper">
    <div class="tab-background">
      <svg viewBox="0 0 1440 768" preserveAspectRatio="none">
        <defs>
          <radialGradient id="leftCircle" cx="50%" cy="50%" r="50%">
            <stop offset="0%" stop-color="#87FFB9" stop-opacity="0.7" />
            <stop offset="100%" stop-color="#87FFB9" stop-opacity="0" />
          </radialGradient>

          <radialGradient id="rightCircle" cx="50%" cy="50%" r="50%">
            <stop offset="0%" stop-color="#F8FFB0" stop-opacity="0.7" />
            <stop offset="100%" stop-color="#F8FFB0" stop-opacity="0" />
          </radialGradient>
        </defs>

        <circle cx="200" cy="1000" r="700" fill="url(#leftCircle)" />
        <circle cx="1400" cy="-200" r="700" fill="url(#rightCircle)" />
      </svg>
    </div>

    <div class="homelink-wrapper">
      <router-link class="homelink" to="/">
        <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path
            d="M15 6L9 12L15 18"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>

        <p>Вернуться на главную</p>
      </router-link>
    </div>

    <div class="form-container">
      <form class="formlogin" @submit.prevent="login">
        <h1>Аутентификация</h1>

        <div class="input-group">
          <input id="email" v-model="email" type="email" placeholder=" " />
          <label for="email">Почта</label>
        </div>
        <p v-if="emailError" class="error-message">{{ emailError }}</p>

        <div class="input-group password-wrapper">
          <input
            :type="showPassword ? 'text' : 'password'"
            id="password"
            v-model="password"
            placeholder=" "
          />
          <label for="password">Пароль</label>
          <span class="eye-icon" @click="togglePassword">
            <svg
              v-if="showPassword"
              xmlns="http://www.w3.org/2000/svg"
              height="20"
              width="20"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                d="M1 1l22 22"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
              <path
                d="M17.94 17.94A10.94 10.94 0 0 1 12 19c-5 0-9.27-3.11-11-7 
                   0-1.14.31-2.2.84-3.13M6.26 6.26A10.94 10.94 0 0 1 12 5
                   c5 0 9.27 3.11 11 7a10.94 10.94 0 0 1-1.58 2.44"
              />
            </svg>
            <svg
              v-else
              xmlns="http://www.w3.org/2000/svg"
              height="20"
              width="20"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
            >
              <path d="M1 12s4-7 11-7 11 7 11 7-4 7-11 7-11-7-11-7z" />
              <circle cx="12" cy="12" r="3" />
            </svg>
          </span>
        </div>
        <p v-if="passwordError" class="error-message">{{ passwordError }}</p>

        <div
          id="captcha"
          class="captcha"
          data-sitekey="6LelWpgrAAAAADeeNQEUvmJZ2yssQooXXzGtHcP2"
        ></div>

        <button type="submit">Войти</button>

        <div class="divider">или</div>

        <router-link class="login-link" to="/registration">
          Зарегистрировать новый аккаунт
        </router-link>

        <p v-if="error" style="color: red; text-align: center">{{ error }}</p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { login as performLogin } from "@/modules/auth/login";

const siteKey = import.meta.env.VITE_RECAPTCHA_SITE_KEY;
const captchaToken = ref("");
const email = ref("");
const password = ref("");
const showPassword = ref(false);
const error = ref("");
const emailError = ref("");
const passwordError = ref("");

function togglePassword() {
  showPassword.value = !showPassword.value;
}

async function login() {
  await performLogin(
    email,
    password,
    error,
    emailError,
    passwordError,
    captchaToken
  );
}

onMounted(() => {
  if (window.grecaptcha) {
    window.grecaptcha.ready(() => {
      window.grecaptcha.render("captcha", {
        sitekey: siteKey,
        callback: (token) => {
          captchaToken.value = token;
        },
        "error-callback": () => {
          error.value = "Ошибка загрузки капчи.";
        },
        "expired-callback": () => {
          captchaToken.value = "";
        },
      });
    });
  }
});
</script>

<style scoped>
.captcha {
  margin-top: 15px;
  align-self: center;
}

.tab-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;

  svg {
    width: 100%;
    height: 100%;
  }
}

.page-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 20px;
  min-height: 100vh;
  width: 100%;
}

.form-container {
  z-index: 1;
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 95%;
}

.homelink-wrapper {
  width: 95%;
  display: flex;
  align-items: left;
}

.homelink {
  z-index: 1;
  margin-top: 15px;
  color: black;
  display: flex;
  flex-flow: row;
  text-decoration: none;

  p {
    font-size: 19px;
  }

  svg {
    height: 24px;
    width: 24px;
  }

  &:hover {
    color: rgb(158, 158, 158);
    text-decoration-line: underline;
  }
}

.formlogin {
  display: flex;
  flex-direction: column;
  justify-items: center;
  max-width: 400px;
  width: 100%;
  padding: 40px;
  margin: 20px 0;
  background: white;
  border-radius: 16px;
  box-shadow: 3px 3px 24px rgb(199, 199, 199);

  h1 {
    text-align: center;
    color: rgb(41, 41, 41);
  }
}

.input-group {
  position: relative;
  margin-top: 24px;

  input {
    width: 100%;
    padding: 14px;
    font-size: 16px;
    background: linear-gradient(
      90deg,
      rgb(223, 255, 228) 25%,
      rgb(251, 255, 223) 75%
    );
    border: none;
    border-radius: 14px;
    outline: none;
    box-sizing: border-box;
  }

  label {
    position: absolute;
    top: 12px;
    left: 12px;
    padding: 0 4px;
    color: gray;
    font-size: 16px;
    transition: 0.2s ease;
    pointer-events: none;
  }
}

.input-group input:focus + label,
.input-group input:not(:placeholder-shown) + label {
  top: -18px;
  left: 8px;
  font-size: 12px;
}

.password-wrapper input {
  padding-right: 40px;
}

.eye-icon {
  position: absolute;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  cursor: pointer;
  color: gray;
}

button[type="submit"] {
  background-color: rgb(59, 59, 59);
  border-radius: 18px;
  color: white;
  padding: 10px;
  margin-top: 15px;
  border: none;
  font-size: 16px;
  font-weight: bold;
  width: 100%;
  transition: 0.4s ease;
}

button[type="submit"]:hover {
  background-color: rgb(126, 193, 255);
}

.divider {
  text-align: center;
  margin: 24px 0 16px;
  color: gray;
  position: relative;
  font-size: 14px;
}

.divider::before,
.divider::after {
  content: "";
  position: absolute;
  top: 50%;
  width: 40%;
  height: 1px;
  background-color: #ccc;
}

.divider::before {
  left: 0;
}

.divider::after {
  right: 0;
}

.login-link {
  display: block;
  text-align: center;
  margin-bottom: 10px;
  font-size: 14px;
  color: #333;
  text-decoration: none;
}

.login-link:hover {
  text-decoration: underline;
  color: rgb(100, 100, 100);
}

.error-message {
  color: red;
  font-size: 13px;
  margin-top: 4px;
  margin-left: 4px;
}
</style>
