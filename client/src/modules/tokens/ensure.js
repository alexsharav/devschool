import token from "@/modules/tokens/token";

export async function ensureAccessToken() {
  if (token.getAccessToken()) return true; // уже есть в памяти
  try {
    const r = await fetch("http://localhost:8080/refresh", {
      method: "POST",
      credentials: "include", // чтобы ушла HttpOnly-кука refresh_token
    });
    if (!r.ok) return false;
    const body = await r.json();
    if (body?.access_token) {
      token.setAccessToken(body.access_token);
      return true;
    }
    return false;
  } catch {
    return false;
  }
}
