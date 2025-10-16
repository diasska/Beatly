<template>
  <div class="widget-card profile-card">
    <div class="user-info">
      <div class="avatar">
        <!-- Иконка пользователя -->
      </div>
      <h3>{{ userName }}</h3>
    </div>

    <div class="calendar-wrapper">
      <v-calendar
        is-expanded
        transparent
        borderless
        locale="ru-RU"
        :attributes="attrs"
      >
        <template #header="{ page }">
          <div class="calendar-header">
            <h4 class="month-title">{{ page.title }}</h4>
            <button class="add-task-btn">+ Add Task</button>
          </div>
        </template>
      </v-calendar>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

defineProps({ userName: String });

const attrs = ref([
  {
    key: 'today',
    highlight: {
      color: '#6c5ce7',
      fillMode: 'solid',
    },
    dates: new Date(),
  },
]);
</script>

<style>
.profile-card {
  background-color: var(--accent-color);
  padding: 24px;
  border-radius: 16px;
}
.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 20px;
}
.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: white;
}


.calendar-wrapper {
  background-color: var(--card-background-color);
  border-radius: 12px;
  padding: 16px;
}

.calendar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  padding: 0 5px 10px 5px;
}

.month-title {
  color: var(--text-color);
  font-weight: 600;
  text-transform: capitalize;
}

.add-task-btn {
  background: none;
  border: 1px solid #eee;
  border-radius: 6px;
  padding: 4px 8px;
  font-size: 12px;
  cursor: pointer;
}

.profile-card .vc-container {
  border: none !important;
  font-family: 'Poppins', sans-serif;
  padding: 0;
}


.profile-card .vc-weekday {
  color: var(--subtle-text-color) !important;
  font-size: 0.8rem;
  font-weight: 500;
}

.profile-card {
  color: var(--text-color) !important;
  font-weight: 500 !important;
  border-radius: 8px;
}
.profile-card .vc-day-content:hover {
  background-color: var(--accent-color) !important;
}

.profile-card .vc-highlight {
  background-color: var(--primary-color) !important;
  color: #fff !important;
}
</style>
