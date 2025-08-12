// Простой модуль для хранения токена в памяти (переживает навигацию,
// но не переживает перезагрузку страницы — это и нужно для безопасности).
let accessToken = null;

export default {
  getAccessToken() {
    return accessToken;
  },
  setAccessToken(token) {
    accessToken = token;
  },
  clear() {
    accessToken = null;
  },
};
