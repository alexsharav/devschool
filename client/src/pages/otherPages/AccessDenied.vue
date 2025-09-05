<template>
  <div class="tab-background">
    <div class="access-denied-wrapper">
      <p>Запрос отказан: неверные права пользователя</p>

      <div class="homelink-wrapper">
        <router-link class="homelink" to="/">
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
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
    </div>
  </div>
</template>

<style scoped>
.tab-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-image: url("data:image/svg+xml,%3Csvg width='600' height='600' viewBox='0 0 64 64' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M8 16c4.418 0 8-3.582 8-8s-3.582-8-8-8-8 3.582-8 8 3.582 8 8 8zm0-2c3.314 0 6-2.686 6-6s-2.686-6-6-6-6 2.686-6 6 2.686 6 6 6zm33.414-6l5.95-5.95L45.95.636 40 6.586 34.05.636 32.636 2.05 38.586 8l-5.95 5.95 1.414 1.414L40 9.414l5.95 5.95 1.414-1.414L41.414 8zM40 48c4.418 0 8-3.582 8-8s-3.582-8-8-8-8 3.582-8 8 3.582 8 8 8zm0-2c3.314 0 6-2.686 6-6s-2.686-6-6-6-6 2.686-6 6 2.686 6 6 6zM9.414 40l5.95-5.95-1.414-1.414L8 38.586l-5.95-5.95L.636 34.05 6.586 40l-5.95 5.95 1.414 1.414L8 41.414l5.95 5.95 1.414-1.414L9.414 40z' fill='%23b3b3b3' fill-opacity='0.4' fill-rule='evenodd'/%3E%3C/svg%3E");
}

.access-denied-wrapper {
  display: flex;
  flex-flow: column wrap;
  align-items: center;
  justify-content: center;

  > p {
    background: rgb(255, 255, 255);
    padding: 24px;
    border-radius: 18px;
    box-shadow: 0px 0px 26px rgb(189, 189, 189);
    font-size: 23px;
  }
}

.homelink {
  color: black;
  display: flex;
  flex-flow: row;
  text-decoration: none;
  border-radius: 18px;
  padding: 24px 24px 24px 14px;
  transition: transform 0.5s ease;

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
</style>
