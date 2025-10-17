<template>
  <div class="task-item" :class="{ completed: task.isCompleted }">
    <div class="checkbox" @click="emit('toggle-status', task.id)">
      <svg v-if="task.isCompleted" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
        <path d="M10.97 4.97a.75.75 0 0 1 1.07 1.05l-3.99 4.99a.75.75 0 0 1-1.08.02L4.324 8.384a.75.75 0 1 1 1.06-1.06l2.094 2.093 3.473-4.425a.267.267 0 0 1 .02-.022z"/>
      </svg>
    </div>

    <RouterLink :to="taskLink" class="task-link">
      <span>{{ task.title }}</span>
    </RouterLink>

    <button @click="emit('delete-task', task.id)" class="delete-btn">×</button>
  </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  task: Object,
  currentDate: String
});
const emit = defineEmits(['toggle-status', 'delete-task']);

const taskLink = computed(() => `/task/${props.task.id}?from=${props.currentDate}`);
</script>

<style scoped>
.task-item {
  display: flex;
  align-items: center;
  gap: 12px;
  background-color: #fdfdfd;
  padding: 14px;
  border-radius: 8px;
  border: 1px solid #f0f0f0;
}

.checkbox {
  width: 20px;
  height: 20px;
  border-radius: 4px;
  border: 2px solid #e0e0e0;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
  cursor: pointer;
  flex-shrink: 0;
}

.task-item.completed .checkbox {
  background-color: var(--primary-color);
  border-color: var(--primary-color);
  color: white;
}

.task-link {
  text-decoration: none;
  color: var(--text-color);
  flex-grow: 1;
}

.task-item.completed .task-link span {
  text-decoration: line-through;
  color: var(--subtle-text-color);
}

.delete-btn {
  background: #eee;
  border: none;
  color: #888;
  cursor: pointer;
  font-size: 14px;
}
</style>