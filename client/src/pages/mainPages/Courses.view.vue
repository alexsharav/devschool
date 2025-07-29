<template>
  <Header />

  <CourseInfo
    v-for="course in courseLessons"
    :imageLink="course.iLink"
    :courseLink="'/courses' + course.cLink"
    :courseText="course.title"
    :courseDescription="course.cDescription"
    :reverseLayout="course.rLayout"
  />

  <Footer />
</template>

<script setup>
import Header from "@/ui/main/Header.vue";
import Footer from "@/ui/main/Footer.vue";
import CourseInfo from "@/ui/courses/CourseInfo.vue";
import courseLessons from "@/router/courseLessons";
</script>

<style scoped></style>
