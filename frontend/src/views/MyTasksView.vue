<template>
  <div class="my-tasks-container">
    <div class="back-container">
      <button class="back-btn" @click="goBack">← Back to Dashboard</button>
    </div>
    <h2>My Tasks</h2>

    <!-- Форма добавления новой задачи -->
    <div class="add-task-form">
      <input v-model="newTask.title" placeholder="Название задачи" />
      <input v-model="newTask.description" placeholder="Описание" />
      <input v-model="newTask.date" type="date" />
      <button @click="createTask">Добавить задачу</button>
    </div>

    <!-- Список всех задач в боксах -->
    <div class="task-grid">
        <div class="task-box" v-for="task in tasks" :key="task.id" :class="{ completed: task.completed }">
            <div class="checkbox" @click="toggleTaskStatus(task.id)">
            <svg v-if="task.completed" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                <path d="M10.97 4.97a.75.75 0 0 1 1.07 1.05l-3.99 4.99a.75.75 0 0 1-1.08.02L4.324 8.384a.75.75 0 1 1 1.06-1.06l2.094 2.093 3.473-4.425a.267.267 0 0 1 .02-.022z"/>
            </svg>
            </div>

            <input v-model="task.title" placeholder="Название" />
            <input v-model="task.description" placeholder="Описание" />
            <input v-model="task.date" type="date" />
            <button @click="updateTask(task.id)">Сохранить</button>
            <button @click="deleteTask(task.id)" class="delete-btn">Удалить</button>
        </div>
</div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

const tasks = ref([]);
const newTask = ref({ title: '', description: '', completed: false, date: new Date().toISOString().split('T')[0] });
const router = useRouter();

function goBack() {
  router.push({ name: 'dashboard' });
}
// Загрузка всех задач
async function fetchTasks() {
  try {
    const response = await axios.get('http://localhost:8080/tasks');
    tasks.value = response.data;
  } catch (error) {
    console.error('Ошибка загрузки задач:', error);
  }
}

onMounted(fetchTasks);

// Создание
async function createTask() {
  if (!newTask.value.title) return;
  try {
    const response = await axios.post('http://localhost:8080/tasks', newTask.value);
    tasks.value.push(response.data);
    newTask.value = { title: '', description: '', completed: false, date: new Date().toISOString().split('T')[0] };
  } catch (error) {
    console.error('Ошибка создания:', error);
  }
}

// Toggle
async function toggleTaskStatus(taskId) {
  const task = tasks.value.find(t => t.id === taskId);
  if (task) {
    try {
      const updated = { ...task, completed: !task.completed };
      await axios.put(`http://localhost:8080/tasks/${taskId}`, updated);
      task.completed = !task.completed;
    } catch (error) {
      console.error('Ошибка toggle:', error);
    }
  }
}

// Редактирование (сохранить)
async function updateTask(taskId) {
  const task = tasks.value.find(t => t.id === taskId);
  if (task) {
    try {
      await axios.put(`http://localhost:8080/tasks/${taskId}`, task);
      console.log('Задача обновлена');
    } catch (error) {
      console.error('Ошибка обновления:', error);
    }
  }
}

// Удаление
async function deleteTask(taskId) {
  try {
    await axios.delete(`http://localhost:8080/tasks/${taskId}`);
    tasks.value = tasks.value.filter(t => t.id !== taskId);
  } catch (error) {
    console.error('Ошибка удаления:', error);
  }
}
</script>

<style scoped>

.back-container {
  display: flex;
  justify-content: flex-start;
  margin-bottom: 10px;
}

.back-btn {
  background-color: #6c63ff;
  color: white;
  border: none;
  padding: 10px 16px;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: background-color 0.2s, transform 0.1s;
}

.back-btn:hover {
  background-color: #5952d4;
  transform: translateY(-1px);
}

h2 {
  font-weight: 600;
  margin-bottom: 20px;
}
.my-tasks-container {
  padding: 20px;
  background-color: #f7f8fc;
  min-height: 100vh;
}

h2 {
  font-weight: 600;
  margin-bottom: 20px;
}

.add-task-form {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-bottom: 30px;
  max-width: 400px;
  background-color: white;
  padding: 16px;
  border-radius: 12px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.05);
}

.add-task-form input,
.add-task-form button {
  width: 100%;
}

.task-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
}

.task-box {
  background-color: #fff;
  border: 1px solid #e3e3e3;
  padding: 16px;
  border-radius: 12px;
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.05);
  display: flex;
  flex-direction: column;
  gap: 8px;
  transition: transform 0.2s, box-shadow 0.2s;
}

.task-box:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 18px rgba(0, 0, 0, 0.1);
}

.task-box.completed {
  background-color: #f0f0f0;
  opacity: 0.8;
}

.checkbox {
  cursor: pointer;
  align-self: flex-end;
}

input {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 6px;
}

button {
  padding: 8px;
  background-color: #6c63ff;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.2s;
}

button:hover {
  background-color: #5952d4;
}

.delete-btn {
  background-color: #ff5a5a;
}

.delete-btn:hover {
  background-color: #e04b4b;
}
</style>