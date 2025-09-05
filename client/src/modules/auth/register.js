import router from "@/router/router";

function validatePassword(pwd) {
  const minLength = 8;
  const maxLength = 100;
  const hasUpperCase = /[A-Z]/.test(pwd);
  const hasNumber = /\d/.test(pwd);

  if (pwd.length < minLength) return "Пароль должен быть не короче 8 символов.";
  if (pwd.length > maxLength) return "Пароль должен быть длинее 100 символов.";
  if (!hasUpperCase)
    return "Пароль должен содержать хотя бы одну заглавную букву.";
  if (!hasNumber) return "Пароль должен содержать хотя бы одну цифру.";
  return "";
}

async function register(username, email, password, confirmPassword, error, usernameError, emailError, passwordError, captchaRef) {
  error.value = "";
  usernameError.value = "";
  emailError.value = "";
  passwordError.value = "";
  
  
  if (!username.value.trim()) {
    usernameError.value = "Поле 'Имя' обязательно.";
  } else if (username.value.length < 1 || username.value.length > 50) {
    usernameError.value = "Имя должно быть от 1 до 50 символов.";
  }

  if (!email.value.trim()) {
    emailError.value = "Поле 'Почта' обязательно.";
  } else if (email.value.length < 1 || username.value.length > 100) {
    emailError.value = "Эмейл должен быть от 1 до 100 символов.";
  }

  const pwdError = validatePassword(password.value);
  if (pwdError) {
    passwordError.value = pwdError;
  }


  if (password.value !== confirmPassword.value) {
    error.value = "Пароли не совпадают.";
  }

  if (!captchaRef.value.trim()) {
    error.value = "Пожалуйста, пройдите проверку reCAPTCHA.";
  }
  
  if (passwordError.value || emailError.value || usernameError.value || error.value) {
    return;
  }

  console.log(captchaRef.value)
  try {
    const response = await fetch("http://localhost:8080/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        username: username.value,
        email: email.value,
        password: password.value,
        token: captchaRef.value,
      }),
    });

    if (!response.ok) {
      const text = await response.text();
      throw new Error(text);
    }

    router.push("/login");
  } catch (err) {
    error.value = err.message;
  }
}

export { register };
