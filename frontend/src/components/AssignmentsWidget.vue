<template>
  <div class="widget-card">
    <h4>Assignments for {{ selectedDateString }}</h4>
    <p class="completed-text">
      {{ completedCount }}/{{ filteredTasks.length }} completed
    </p>

    <div class="task-list">
      <div
        v-for="task in filteredTasks"
        :key="task.id"
        class="task-item"
        :class="{ completed: task.completed }"
      >
        <label>
          <input
            type="checkbox"
            :checked="task.completed"
            @change="toggleTaskStatus(task.id)"
          />
          <span class="task-title">{{ task.title }}</span>
        </label>
        <p class="task-desc">{{ task.description }}</p>
      </div>

      <p v-if="filteredTasks.length === 0" class="empty">
        No tasks for this date.
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'

const props = defineProps({
  selectedDate: Date
})

const tasks = ref([])

// Форматируем дату в строку YYYY-MM-DD
const selectedDateString = computed(() => {
  if (!props.selectedDate) return ''
  const year = props.selectedDate.getFullYear()
  const month = String(props.selectedDate.getMonth() + 1).padStart(2, '0')
  const day = String(props.selectedDate.getDate()).padStart(2, '0')
  return `${year}-${month}-${day}`
})

// Фильтрация по текущей дате
const filteredTasks = computed(() => {
  if (!selectedDateString.value) return []
  return tasks.value.filter(task => task.date === selectedDateString.value)
})

// Счётчик выполненных
const completedCount = computed(() => filteredTasks.value.filter(t => t.completed).length)

// Загрузка всех задач (с бэкенда)
onMounted(async () => {
  try {
    const response = await axios.get('http://localhost:8080/tasks')
    tasks.value = response.data
  } catch (error) {
    console.error('Ошибка загрузки задач:', error)
  }
})

// Переключение статуса выполнения
async function toggleTaskStatus(taskId) {
  const task = tasks.value.find(t => t.id === taskId)
  if (task) {
    try {
      const updatedTask = { ...task, completed: !task.completed }
      await axios.put(`http://localhost:8080/tasks/${taskId}`, updatedTask)
      task.completed = !task.completed
    } catch (error) {
      console.error('Ошибка обновления статуса:', error)
    }
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

.task-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.task-item {
  background-color: #fff;
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 10px 12px;
  transition: background-color 0.2s;
}

.task-item.completed {
  background-color: #e8f5e9;
  text-decoration: line-through;
  color: #777;
}

.task-title {
  font-weight: 600;
  margin-left: 8px;
}

.task-desc {
  margin: 4px 0 0 24px;
  font-size: 0.9rem;
  color: #555;
}

.empty {
  color: #aaa;
  font-style: italic;
  text-align: center;
  margin-top: 12px;
}
</style>