<template>
  <p><slot /></p>
</template>

<style scoped>
p {
  border: 1px solid black;
  border-radius: 20px;
  padding: 10px;
}

@media (max-width: 600px) {
  p {
    font-size: 12px;
    padding: 6px 10px;
    line-height: 1.2;
    word-break: break-word;
    -webkit-hyphens: auto;
    hyphens: auto;
  }
}

@media (max-width: 420px) {
  p {
    flex: 1 1 auto;
    text-align: center;
  }
}
</style>
