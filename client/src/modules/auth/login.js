import router from "@/router";

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
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: email.value,
        password: password.value,
      }),
      credentials: "include",
    });

    if (response.ok) {
      router.push("/profile");
    } else {
      error.value = "Ошибка авторизации";
    }
  } catch (err) {
    error.value = err.message;
  }
}

export { login }