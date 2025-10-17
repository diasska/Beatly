<template>
  <div class="task-detail-container">
    <div v-if="task">
      <h1>{{ task.title }}</h1>
      <p>
        <strong>Status:</strong>
        <span :class="task.isCompleted ? 'status-completed' : 'status-progress'">
          {{ task.isCompleted ? 'Completed' : 'In Progress' }}
        </span>
      </p>

      <button @click="toggleStatus" class="status-button">
        {{ task.isCompleted ? 'Mark as In Progress' : 'Mark as Completed' }}
      </button>

    </div>
    <p v-else>
      Loading task... or task not found.
    </p>

    <RouterLink :to="backLink" class="back-link">← Back to Dashboard</RouterLink>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import axios from 'axios';

const route = useRoute();
const task = ref(null);


onMounted(async () => {
  const taskId = parseInt(route.params.id);
  try {
    const response = await axios.get(`http://localhost:8080/tasks/${taskId}`);
    task.value = response.data;
  } catch (error) {
    console.error('Ошибка загрузки задачи:', error);
  }
});

async function toggleStatus() {
  if (task.value) {
    try {
      const updatedTask = { ...task.value, completed: !task.value.completed };
      await axios.put(`http://localhost:8080/tasks/${task.value.id}`, updatedTask);
      task.value.completed = !task.value.completed;
    } catch (error) {
      console.error('Ошибка обновления статуса:', error);
    }
  }
}

const backLink = computed(() => {
  const fromDate = route.query.from;
  if (fromDate) {
    return `/?date=${fromDate}`;
  }
  return '/';
});
</script>

<style scoped>
.task-detail-container {
  padding: 40px;
  max-width: 600px;
  margin: 50px auto;
  background-color: white;
  border-radius: 16px;
  box-shadow: 0 8px 24px rgba(0,0,0,0.1);
}
h1 { margin-bottom: 20px; }
.status-completed { color: green; font-weight: bold; }
.status-progress { color: orange; font-weight: bold; }
.status-button {
  margin: 20px 0;
  padding: 10px 20px;
  border: none;
  background-color: var(--primary-color);
  color: white;
  border-radius: 8px;
  cursor: pointer;
}
.back-link {
  display: inline-block;
  margin-top: 20px;
}
</style>