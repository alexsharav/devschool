<template>
  <div class="tab-background">
    <svg viewBox="0 0 1440 768" preserveAspectRatio="none">
      <defs>
        <radialGradient id="leftCircle" cx="50%" cy="50%" r="50%">
          <stop offset="0%" stop-color="#87FFB9" stop-opacity="0.7" />
          <stop offset="100%" stop-color="#87FFB9" stop-opacity="0" />
        </radialGradient>

        <radialGradient id="rightCircle" cx="50%" cy="50%" r="50%">
          <stop offset="0%" stop-color="#F8FFB0" stop-opacity="0.7" />
          <stop offset="100%" stop-color="#F8FFB0" stop-opacity="0" />
        </radialGradient>
      </defs>

      <circle cx="200" cy="1000" r="700" fill="url(#leftCircle)" />
      <circle cx="1400" cy="-200" r="700" fill="url(#rightCircle)" />
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
</template>

<style scoped>
.tab-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;

  svg {
    width: 100%;
    height: 100%;
  }
}

.homelink-wrapper {
  display: flex;
  min-height: 100vh;
  align-items: center;
}

.homelink {
  color: black;
  display: flex;
  flex-flow: row;
  text-decoration: none;
  box-shadow: 3px 3px 24px rgb(185, 185, 185);
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
    transform: scale(1.1);
    text-decoration-line: underline;
  }
}
</style>
