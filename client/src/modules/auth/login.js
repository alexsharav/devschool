import router from "@/router/router";

async function login(email, password, error, emailError, passwordError, captchaRef) {
  error.value = "";
  emailError.value = "";
  passwordError.value = "";

  if (!email.value.trim()) {
    emailError.value = "Поле почта не заполнено";
  }
  if (!password.value.trim()) {
    passwordError.value = "Пароль не введен";
  }
  if (!captchaRef.value.trim()) {
    error.value = "Пожалуйста, пройдите проверку reCAPTCHA.";
  }
  if (passwordError.value || emailError.value || error.value) {
    return;
  }

  try {
    const response = await fetch("http://localhost:8080/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        email: email.value,
        password: password.value,
        token: captchaRef.value,
      }),
      credentials: "include", // нужно, чтобы пришла HttpOnly кука с refresh_token
    });

    if (!response.ok) {
      error.value = "Ошибка авторизации";
      token.clear();
      return;
    }

    router.push("/profile");
  } catch (err) {
    error.value = "Сетевая ошибка";
    token.clear();
  }
}

export { login };
