<template>
  <div class="task-item" :class="{ completed: task.isCompleted }" @click="emit('toggle-status', task.id)">
    <div class="checkbox">
      <svg v-if="task.isCompleted" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
        <path d="M10.97 4.97a.75.75 0 0 1 1.07 1.05l-3.99 4.99a.75.75 0 0 1-1.08.02L4.324 8.384a.75.75 0 1 1 1.06-1.06l2.094 2.093 3.473-4.425a.267.267 0 0 1 .02-.022z"/>
      </svg>
    </div>
    <span>{{ task.title }}</span>
  </div>
</template>
<script setup>
defineProps({ task: Object });
const emit = defineEmits(['toggle-status']);
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
  cursor: pointer;
  transition: all 0.2s ease;
}

.task-item:hover {
  border-color: var(--primary-color);
  background-color: var(--accent-color);
}

.task-item.completed span {
  text-decoration: line-through;
  color: var(--subtle-text-color);
}

.checkbox {
  width: 20px;
  height: 20px;
  border-radius: 4px;
  border: 2px solid #e0e0e0; /* <-- Серая рамка для пустого чекбокса */
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}

/* Стили для отмеченного чекбокса */
.task-item.completed .checkbox {
  background-color: var(--primary-color);
  border-color: var(--primary-color);
  color: white; /* Цвет для галочки */
}
</style>
