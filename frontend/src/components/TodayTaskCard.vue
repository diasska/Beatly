<template>
  <div class="card">
    <div class="text-content">
      <h2>Today Task</h2>
      <p>Check your daily tasks</p>
      <button class="check-all-btn" @click="goToMyTasks">Check all →</button>
    </div>
    <img src="@/assets/illus.png" alt="Person working on tasks" class="hero-image" />
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router' 

const router = useRouter()


function goToMyTasks() {
  router.push({ name: 'my-tasks' }) 
}
</script>

<style scoped>
.card {
  border-radius: 16px;
  padding: 32px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  overflow: hidden;
}

.text-content p {
  color: var(--subtle-text-color);
  margin: 8px 0 24px;
}

.check-all-btn {
  background-color: var(--primary-color);
  color: white;
  border: none;
  padding: 12px 24px;
  border-radius: 8px;
  font-weight: 500;
  cursor: pointer;
}

.hero-image {
  max-width: 200px;
  height: auto;
}
</style>