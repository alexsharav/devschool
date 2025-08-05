function hasCookie(name) {
  const cookies = document.cookie.split("; ");
  return cookies.some(cookie => cookie.startsWith(name + "="));
}

export { hasCookie };
