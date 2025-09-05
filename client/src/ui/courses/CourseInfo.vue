<template>
  <section :class="['tab', { 'reverse-layout': reverseLayout }]">
    <img
      class="course-img"
      :src="imageLink"
      alt="course-img"
      @click="router.push({ path: courseLink })"
    />

    <div class="course-text-desc">
      <router-link class="course-text" :to="courseLink">{{
        courseText
      }}</router-link>

      <p class="course-description">{{ courseDescription }}</p>

      <router-link class="course-link" :to="courseLink"
        >Посмотреть курс</router-link
      >
    </div>
  </section>
</template>

<script setup>
import router from "@/router/router";

defineProps({
  imageLink: [String, Object],
  courseLink: String,
  courseText: String,
  courseDescription: String,
  reverseLayout: {
    type: Boolean,
    default: false,
  },
});
</script>

<style scoped>
.tab:first-of-type {
  margin-top: 50px;
}

.tab {
  padding: 60px 0;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 30px;
  max-width: 1200px;
  width: 98%;
  height: auto;
}

.reverse-layout {
  flex-flow: row-reverse;
}

.course-img {
  border-radius: 20px;
  width: 100%;
  height: auto;
  max-width: 600px;
  flex-shrink: 2;
  border: 1px solid rgb(168, 255, 168);
  cursor: pointer;
  transition: transform 0.2s ease;

  &:hover {
    border: 1px solid rgb(109, 109, 109);
    transform: translateY(-3px);
  }
}

.course-text {
  color: rgb(41, 41, 41);
  font-weight: normal;
  font-size: 30px;
  text-align: justify;
  text-decoration: none;
}

.course-text:hover {
  text-decoration: underline;
}

.course-description {
  margin: 20px 0px 0px 0px;
  font-size: 20px;
  color: rgb(73, 73, 73);
  text-align: justify;
}

.course-text-desc {
  display: flex;
  flex-flow: column;
  max-width: 600px;
  width: 100%;
}

.course-link {
  margin: 10px 0px 0px 0px;
  font-size: 18px;
  color: rgb(0, 0, 0);
  text-decoration-line: underline;
}

.course-link:hover {
  color: rgb(93, 93, 93);
}

@media (max-width: 910px) {
  .tab {
    flex-flow: column wrap;
  }
}

@media (max-width: 750px) {
  .tab {
    padding: 30px 20px;
    gap: 15px;
  }

  .course-text {
    font-size: 24px;
    text-align: center;
  }

  .course-description {
    font-size: 18px;
  }

  .course-link {
    font-size: 16px;
    text-align: center;
  }
}

@media (max-width: 400px) {
  .tab {
    padding: 26px 16px;
    gap: 10px;
  }

  .course-text {
    font-size: 20px;
  }

  .course-description {
    font-size: 14px;
  }

  .course-link {
    font-size: 12px;
  }
}
</style>
