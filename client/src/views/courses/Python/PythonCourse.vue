<template>
  <CourseLessonList
    :courseTitle="course.title"
    :courseDesc="course.cDescription"
    :chapters="course.chapters"
  />
</template>

<script setup>
import CourseLessonList from "@/components/coursesComponents/CourseLessonlist.vue";
import allLessons from "@/router/courseLessons.js";

const course = allLessons["Python"];
</script>
