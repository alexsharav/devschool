<template>
  <section class="tab">
    <div class="tab-background">
      <svg viewBox="0 0 1440 768" preserveAspectRatio="none">
        <defs>
          <radialGradient id="leftCircle" cx="50%" cy="50%" r="50%">
            <stop offset="0%" stop-color="#00f767" stop-opacity="0.8" />
            <stop offset="100%" stop-color="#00f767" stop-opacity="0" />
          </radialGradient>

          <radialGradient id="rightCircle" cx="50%" cy="50%" r="50%">
            <stop offset="0%" stop-color="#0CFAE2" stop-opacity="0.8" />
            <stop offset="100%" stop-color="#0CFAE2" stop-opacity="0" />
          </radialGradient>
        </defs>

        <circle cx="200" cy="-200" r="950" fill="url(#leftCircle)" />
        <circle cx="1400" cy="-200" r="950" fill="url(#rightCircle)" />
      </svg>
    </div>

    <h1 class="dev-tab-main-text">{{ mainText }}</h1>
    <p class="dev-tab-desc">{{ secondaryText }}</p>
    <slot name="buttons"></slot>
  </section>
</template>

<script setup>
defineProps({
  mainText: String,
  secondaryText: String,
});
</script>

<style scoped>
.tab {
  position: relative;
  align-items: center;
  display: flex;
  flex-flow: column wrap;
  width: 100%;
  overflow: hidden;
}

.tab-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
  pointer-events: none;

  svg {
    width: 100%;
    height: 100%;
  }
}

.dev-tab-main-text {
  margin-top: 200px;
  font-size: 50px;
  font-weight: bold;
  color: rgb(41, 41, 41);
}

.dev-tab-desc {
  padding: 10px 0px;
  margin-top: 20px;
  font-size: 20px;
  color: rgb(90, 90, 90);
  text-align: center;
}

@media (max-width: 750px) {
  .dev-tab-desc {
    font-size: 17px;
    margin-top: 10px;
    padding: 10px 5%;
  }

  .dev-tab-main-text {
    font-size: 30px;
    margin-top: 160px;
  }
}

@media (max-width: 400px) {
  .dev-tab-desc {
    font-size: 14px;
  }

  .dev-tab-main-text {
    font-size: 22px;
    margin-top: 120px;
    padding: 10px 0px;
  }
}
</style>
