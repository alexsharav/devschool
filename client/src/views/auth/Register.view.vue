<template>
  <div class="page-wrapper">
    <div class="tab-background">
      <svg viewBox="0 0 1440 768" preserveAspectRatio="none">
        <defs>
          <radialGradient id="leftCircle" cx="50%" cy="50%" r="50%">
            <stop offset="0%" stop-color="#a3ffea" stop-opacity="0.8" />
            <stop offset="100%" stop-color="#a3ffea" stop-opacity="0" />
          </radialGradient>

          <radialGradient id="rightCircle" cx="50%" cy="50%" r="50%">
            <stop offset="0%" stop-color="#B1FFCA" stop-opacity="0.5" />
            <stop offset="100%" stop-color="#B1FFCA" stop-opacity="0" />
          </radialGradient>
        </defs>

        <circle cx="200" cy="1000" r="1200" fill="url(#leftCircle)" />
        <circle cx="1400" cy="-200" r="1200" fill="url(#rightCircle)" />
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
      <form class="formregister" @submit.prevent="register">
        <h1>Регистрация</h1>

        <div class="input-group">
          <input id="username" v-model="username" placeholder=" " />
          <label for="username">Имя</label>
        </div>
        <p v-if="usernameError" class="error-message">{{ usernameError }}</p>

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

        <div class="input-group password-wrapper">
          <input
            :type="showPassword ? 'text' : 'password'"
            id="confirmPassword"
            v-model="confirmPassword"
            placeholder=" "
          />
          <label for="confirmPassword">Повторите пароль</label>
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

        <button type="submit">Зарегистрироваться</button>

        <div class="divider">или</div>

        <router-link class="login-link" to="/login">
          Войти в имеющийся аккаунт
        </router-link>

        <p v-if="error" style="color: red">{{ error }}</p>
        <p v-if="success" style="color: green">Регистрация успешна!</p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const username = ref("");
const email = ref("");
const password = ref("");
const confirmPassword = ref("");
const showPassword = ref(false);
const success = ref(false);
const error = ref("");
const usernameError = ref("");
const emailError = ref("");

function togglePassword() {
  showPassword.value = !showPassword.value;
}

function validatePassword(pwd) {
  const minLength = 8;
  const hasUpperCase = /[A-Z]/.test(pwd);
  const hasNumber = /\d/.test(pwd);

  if (pwd.length < minLength) return "Пароль должен быть не короче 8 символов.";
  if (!hasUpperCase)
    return "Пароль должен содержать хотя бы одну заглавную букву.";
  if (!hasNumber) return "Пароль должен содержать хотя бы одну цифру.";
  return "";
}

async function register() {
  usernameError.value = "";
  emailError.value = "";

  if (!username.value.trim()) {
    usernameError.value = "Поле 'Имя' обязательно.";
  } else if (username.value.length == 1 || username.value.length > 32) {
    usernameError.value = "Поле должно содержать от 1 до 32 символов";
  }

  if (!email.value.trim()) {
    emailError.value = "Поле 'Почта' обязательно.";
  } else if (email.value.length == 1 || email.value.length > 32) {
    emailError.value = "Поле должно содержать от 10 до 256 символов";
  }

  if (usernameError.value || emailError.value) {
    success.value = false;
    return;
  }

  const pwdError = validatePassword(password.value);
  if (pwdError) {
    passwordError.value = pwdError;
    success.value = false;
    return;
  }

  if (password.value !== confirmPassword.value) {
    error.value = "Пароли не совпадают.";
    success.value = false;
    return;
  }

  try {
    const response = await fetch("http://localhost:8080/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        username: username.value,
        email: email.value,
        password: password.value,
      }),
    });

    if (!response.ok) {
      const text = await response.text();
      throw new Error(text);
    }

    success.value = true;
    error.value = "";
  } catch (err) {
    success.value = false;
    error.value = err.message;
  }
}
</script>

<style scoped>
.tab-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: -1;

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

.formregister {
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
    padding: 12px;
    font-size: 16px;
    background: linear-gradient(
      90deg,
      rgba(209, 246, 255, 1) 0%,
      rgba(227, 255, 240, 1) 51%
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
  margin-top: 30px;
  border: none;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
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
