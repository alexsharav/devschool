import { checktoken } from "./tokenchecker";

async function requireGuest(to, from, next) {
  const isAuthenticated = await checktoken();

  if (isAuthenticated) {
    // Если авторизован, редиректим на главную
    next("/accessdenied");
  } else {
    // Если не авторизован, разрешаем переход
    next();
  }
}

// Функция для защиты маршрутов, требующих авторизации
async function requireAuth(to, from, next) {
  const isAuthenticated = await checktoken();

  if (isAuthenticated) {
    next();
  } else {
    next("/accessdenied");
  }
}

export { requireAuth, requireGuest };
