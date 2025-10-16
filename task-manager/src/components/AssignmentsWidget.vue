<template>
  <div class="widget-card">
    <h4>Assignments ({{ tasks.length }})</h4>
    <p class="completed-text">{{ completedTasksCount }} / {{ tasks.length }} completed</p>
    <div class="task-list">
      <TaskItem v-for="task in tasks" :key="task.id" :task="task" @toggle-status="toggleTaskStatus" />
    </div>
  </div>
</template>
<script setup>
import { ref, computed } from 'vue';
import TaskItem from './TaskItem.vue';
const tasks = ref([
  { id: 1, title: 'Task 1', isCompleted: true }, { id: 2, title: 'Task 2', isCompleted: true },
  { id: 3, title: 'Task 3', isCompleted: false }, { id: 4, title: 'Task 4', isCompleted: false },
]);
const completedTasksCount = computed(() => tasks.value.filter(t => t.isCompleted).length);
function toggleTaskStatus(taskId) {
  const task = tasks.value.find(t => t.id === taskId);
  if (task) task.isCompleted = !task.isCompleted;
}
</script>
<style scoped>
.widget-card {
  background-color: var(--card-background-color);
  padding: 24px;
  border-radius: 16px;
  box-shadow: 0 8px 24px var(--shadow-color);
  flex-grow: 1;
}
</style>
