<template>
  <div class="widget-card">
    <h4>Assignments ({{ tasks.length }})</h4>
    <p class="completed-text">{{ filteredTasks.length }} tasks for selected date</p>

    <div class="task-list">
      <TaskItem
        v-for="task in filteredTasks"
        :key="task.id"
        :task="task"
        :current-date="selectedDateString"
      @toggle-status="toggleTaskStatus"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import TaskItem from './TaskItem.vue';

const props = defineProps({
  selectedDate: Date
});

const tasks = ref([
  { id: 1, title: 'Task 1', isCompleted: true, date: '2025-10-17' },
  { id: 2, title: 'Task 2', isCompleted: true, date: '2025-10-17' },
  { id: 3, title: 'Task 3', isCompleted: false, date: '2025-10-18' },
  { id: 4, title: 'Task 4', isCompleted: false, date: '2025-10-16' },
]);

const selectedDateString = computed(() => {
  if (!props.selectedDate) return '';
  return props.selectedDate.toISOString().split('T')[0];
});

const filteredTasks = computed(() => {
  if (!selectedDateString.value) return [];
  return tasks.value.filter(task => task.date === selectedDateString.value);
});

function toggleTaskStatus(taskId) {
  const task = tasks.value.find(t => t.id === taskId);
  if (task) {
    task.isCompleted = !task.isCompleted;
  }
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
