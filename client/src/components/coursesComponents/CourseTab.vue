<template>
  <section class="course-tab">
    <div class="center-div">
      <h1 class="course-tab-main-text">{{ courseMainText }}</h1>
      <p class="course-tab-desc">{{ courseSecondaryText }}</p>
    </div>
  </section>
</template>

<script setup>
const props = defineProps({
  courseMainText: String,
  courseSecondaryText: String,
});
</script>

<style scoped>
.course-tab {
  align-items: center;
  display: flex;
  flex-flow: column wrap;
  width: 98%;
  margin-top: 14px;
  border-radius: 30px;
  border: 1px solid rgb(176, 252, 255);
  background: #91ffcf;
  background: linear-gradient(
    90deg,
    rgba(145, 255, 207, 1) 0%,
    rgba(170, 252, 151, 1) 100%
  );
  padding-bottom: 100px;
}

.course-tab-main-text {
  margin-top: 140px;
  font-size: 45px;
  font-weight: bold;
  color: rgb(63, 63, 63);
}

.course-tab-desc {
  padding: 10px 0px;
  margin-top: 20px;
  font-size: 25px;
  color: rgb(71, 71, 71);
}

.center-div {
  display: flex;
  flex-flow: column;
  justify-items: left;
  width: 60%;
}

@media (max-width: 750px) {
  .course-tab-desc {
    font-size: 17px;
    margin-top: 10px;
    padding: 10px 0%;
  }

  .course-tab-main-text {
    font-size: 30px;
    margin-top: 130px;
  }

  .center-div {
    width: 70%;
  }

  .course-tab {
    padding-bottom: 40px;
  }
}

@media (max-width: 400px) {
  .course-tab-desc {
    font-size: 16px;
  }

  .course-tab-main-text {
    font-size: 24px;
    margin-top: 100px;
    padding: 10px 0px;
  }

  .center-div {
    width: 90%;
  }

  .course-tab {
    padding-bottom: 30px;
  }
}
</style>
