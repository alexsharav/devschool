<template>
  <Header />
  
  <CourseInfo v-for="course in courseLessons" :imageLink="course.iLink" :courseLink="'/courses' + course.cLink"
    :courseText="course.title" :courseDescription="course.cDescription" :reverseLayout="course.rLayout" />

  <Footer />
</template>

<script setup>
import Header from "@/components/mainComponents/Header.vue";
import Footer from "@/components/mainComponents/Footer.vue";
import CourseInfo from "@/components/coursesComponents/CourseInfo.vue";
import courseLessons from "@/router/courseLessons";
</script>

<style scoped>
</style>