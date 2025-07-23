<template>
  <div class="rain-container">
    <canvas ref="canvas" class="rain-canvas"></canvas>
  </div>
</template>

<script>
import { ref, onMounted, onUnmounted } from "vue";

export default {
  setup() {
    const canvas = ref(null);
    let animationFrame;
    let raindrops = [];
    let ctx;

    class Raindrop {
      constructor(width, height) {
        this.width = width;
        this.height = height;
        this.reset();
      }

      reset() {
        this.x = Math.random() * this.width;
        this.y = Math.random() * -this.height;
        this.z = Math.random() * 0.5 + 0.5;
        this.length = Math.floor(this.z * 20) + 10;
        this.speed = this.z * 10 + 5;
        this.thickness = this.z * 1.5;
      }

      fall() {
        this.y += this.speed;
        if (this.y > this.height) {
          this.reset();
        }
      }

      draw() {
        ctx.beginPath();
        ctx.lineWidth = this.thickness;
        ctx.strokeStyle = `rgba(255, 255, 255, ${this.z * 0.8})`;
        ctx.moveTo(this.x, this.y);
        ctx.lineTo(this.x, this.y + this.length);
        ctx.stroke();
      }
    }

    const initRain = () => {
      if (!canvas.value) return;

      ctx = canvas.value.getContext("2d");
      resizeCanvas();

      raindrops = [];
      const raindropCount = Math.floor(
        (canvas.value.width * canvas.value.height) / 10000,
      );

      for (let i = 0; i < raindropCount; i++) {
        raindrops.push(new Raindrop(canvas.value.width, canvas.value.height));
      }
    };

    const resizeCanvas = () => {
      if (!canvas.value) return;
      canvas.value.width = canvas.value.clientWidth;
      canvas.value.height = canvas.value.clientHeight;
    };

    const animate = () => {
      if (!canvas.value || !ctx) return;

      ctx.fillStyle = "rgba(0, 0, 0, 0.15)";
      ctx.fillRect(0, 0, canvas.value.width, canvas.value.height);

      raindrops.forEach((drop) => {
        drop.fall();
        drop.draw();
      });

      animationFrame = requestAnimationFrame(animate);
    };

    const handleResize = () => {
      resizeCanvas();
      initRain();
    };

    onMounted(() => {
      initRain();
      animate();
      window.addEventListener("resize", handleResize);
    });

    onUnmounted(() => {
      window.removeEventListener("resize", handleResize);
      cancelAnimationFrame(animationFrame);
    });

    return { canvas };
  },
};
</script>

<style scoped>
.rain-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: #000;
  overflow: hidden;
  margin: 0;
  padding: 0;
  cursor: pointer;
}

.rain-canvas {
  display: block;
  width: 100%;
  height: 100%;
}
</style>
