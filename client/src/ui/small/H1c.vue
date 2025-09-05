<template>
  <h1><slot /></h1>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 50px;
  text-align: left;
  line-height: 0.95;
}

@media (max-width: 1350px) {
  h1 {
    font-size: 44px;
  }
}

@media (max-width: 830px) {
  h1 {
    font-size: 40px;
  }
}

@media (max-width: 770px) {
  h1 {
    font-size: 36px;
  }
}

@media (max-width: 550px) {
  h1 {
    font-size: 30px;
  }
}
</style>
