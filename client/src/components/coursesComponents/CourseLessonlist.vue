<template>
  <Header />
  <CourseTab :courseMainText="courseTitle" :courseSecondaryText="courseDesc" />

  <div class="lessons-list">
    <div v-for="(chapter, chapterIndex) in chapters" class="flex-div">
      <h1 class="chapter-header">
        {{ chapterIndex + 1 }}. {{ chapter.chapterTitle }}
      </h1>

      <router-link
        v-for="(lesson, lessonIndex) in chapter.lessons"
        class="lesson-routers"
        :to="lesson.path"
      >
        {{ lessonIndex + 1 }}. {{ lesson.title }}
      </router-link>
    </div>
  </div>

  <Footer />
</template>

<script setup>
import Footer from "../mainComponents/Footer.vue";
import Header from "../mainComponents/Header.vue";
import CourseTab from "./CourseTab.vue";

defineProps({
  courseTitle: String,
  courseDesc: String,
  chapters: Array,
});
</script>

<style scoped>
.lessons-list {
  width: 50%;
  margin-top: 40px;
}

.flex-div {
  display: flex;
  flex-flow: column;
}

.chapter-header {
  color: rgb(85, 85, 85);
  font-weight: bold;
  padding-bottom: 20px;
  margin-top: 20px;
  font-size: 30px;
}

.lesson-routers {
  font-size: 23px;
  color: #4b4b4b;
  text-decoration: none;
  margin-left: 40px;
  margin-top: 20px;
  padding-bottom: 20px;
  cursor: pointer;
}

.lesson-routers:hover {
  text-decoration-line: underline;
}

.chapter-header,
.lesson-routers {
  border-bottom: 1px solid rgb(219, 219, 219);
}

@media (max-width: 750px) {
  .lessons-list {
    width: 80%;
    margin-top: 40px;
  }

  .chapter-header {
    font-size: 22px;
  }

  .lesson-routers {
    font-size: 17px;
    margin-top: 14px;
    padding-bottom: 17px;
  }
}

@media (max-width: 400px) {
  .lessons-list {
    width: 94%;
    margin-top: 40px;
  }

  .chapter-header {
    font-size: 19px;
  }

  .lesson-routers {
    font-size: 15px;
    margin-top: 12px;
    padding-bottom: 15px;
  }
}
</style>
