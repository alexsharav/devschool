<template>
  <Header />

  <CourseTab
    :courseMainText="course.title"
    :courseSecondaryText="course.cDescription"
  />

  <CourseLessonList
    :courseTitle="course.title"
    :courseDesc="course.cDescription"
    :chapters="course.chapters"
  />

  <Footer />
</template>

<script setup>
import Header from "@/ui/main/Header.vue";
import Footer from "@/ui/main/Footer.vue";
import CourseLessonList from "@/ui/courses/CourseLessonlist.vue";
import allLessons from "@/router/courseLessons.js";
import CourseTab from "@/ui/courses/CourseTab.vue";

const course = allLessons["DataAnalysis"];
</script>
