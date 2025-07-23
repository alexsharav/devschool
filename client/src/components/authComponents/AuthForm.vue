<!-- <template>
  <section class="container-page">
    <div class="auth-container">
      <button class="remover-button">
        <img class="x-image" :src="xPNG" @click="router.push({ path: '/' })" />
      </button>

      <form name="auth-form">
        <h1 class="caption-form">{{ captionForm }}</h1>

        <div class="item-form">
          <input
            type="email"
            id="auth"
            class="input-form"
            v-model="userData.email"
            placeholder="Почта"
            required
          />
        </div>

        <div class="item-form">
          <input
            type="password"
            id="password"
            class="input-form"
            v-model="userData.password"
            placeholder="Пароль"
            required
          />
        </div>

        <div class="buttons-form">
          <button class="auth-button" @click="authFunc">
            <span style="color: white">{{ buttonForm }}</span>
          </button>

          <br />

          <router-link class="reg-link" :to="authAnotherLink">
            <span class="no-acc">{{ linkQuestion1 }} </span>
            <span class="reg-question">{{ linkQuestion2 }}</span>
          </router-link>
        </div>
      </form>
    </div>
  </section>
</template>

<script setup>
import router from "@/router";
import { ref } from "vue";
import xPNG from "@/views/bgImages/x.png";
const props = defineProps({
  captionForm: String,
  buttonForm: String,
  linkQuestion1: String,
  linkQuestion2: String,
  authAnotherLink: String,
  authFunc: Function,
});

const userData = ref({
  email: "",
  password: "",
});
</script>

<style scoped>
.container-page {
  z-index: 2;
  min-height: 100vh;
  width: 100vw;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 10px;
  box-sizing: border-box;
}

.auth-container {
  background: rgb(255, 255, 255);
  border: 2px solid rgba(199, 198, 198, 0.336);
  width: 100%;
  max-width: 400px;
  display: flex;
  flex-flow: column wrap;
  align-items: center;
  border-radius: 18px;
  overflow: auto;
  padding: 20px;
  box-sizing: border-box;
  min-height: auto;
}

.auth-form {
  display: flex;
  flex-flow: column;
  width: 100%;
}

.remover-button {
  border: none;
  cursor: pointer;
  display: block;
  margin-left: auto;
  margin-top: 5px;
  margin-right: 5px;
  background: none;
}

.x-image {
  width: 15px;
  height: 15px;
}

.caption-form {
  text-align: center;
  margin: 10px 0px;
  font-weight: bolder;
  font-size: clamp(24px, 5vw, 29px);
  color: rgb(65, 65, 65);
}

.item-form {
  width: 100%;
  padding: 20px 0;
  display: flex;
  justify-content: center;
}

.input-form {
  border: none;
  border-bottom: 2px solid gray;
  height: 45px;
  width: 90%;
  max-width: 300px;
  padding: 0 10px;
  box-sizing: border-box;
}

.input-form::placeholder {
  color: rgb(102, 102, 102);
  font-weight: bolder;
  font-size: 14px;
}

.buttons-form {
  display: flex;
  flex-flow: column;
  align-items: center;
  width: 100%;
  padding: 45px 0;
}

.auth-button {
  border-radius: 18px;
  background-color: rgb(37, 37, 37);
  border: none;
  cursor: pointer;
  height: 40px;
  width: 90%;
  max-width: 300px;
  transition: background-color 0.5s ease;
  color: white;
  font-weight: bold;
}

.auth-button:hover {
  background-color: rgb(87, 87, 87);
}

.reg-link {
  padding: 17px 0;
  padding-bottom: 0px;
  color: rgb(0, 0, 0);
  font-size: 16px;
  cursor: pointer;
  text-decoration: none;
  text-align: center;
}

.no-acc {
  color: rgb(65, 65, 65);
}

.reg-question {
  text-decoration-line: underline;
}

@media (max-width: 350px) {
  .input-form,
  .auth-button {
    width: 95%;
  }
}
</style> -->
