<template>
  <div class="widget-card notes-card">
    <div class="notes-header">
      <h4>Notes</h4>
    </div>

    <form @submit.prevent="addNote" class="add-note-form">
      <input
        v-model="newNoteContent"
        type="text"
        placeholder="Add a new note..."
      />
      <button type="submit">+</button>
    </form>

    <ul class="notes-list">
      <li
        v-for="note in notes"
        :key="note.id"
        class="note-item"
      >
        <input
          v-if="editingNoteId === note.id"
          v-model="note.content"
          @blur="finishEditing(note)"
          @keyup.enter="finishEditing(note)"
          class="edit-input"
          v-focus
        />
        <span v-else @dblclick="startEditing(note)">
          {{ note.content }}
        </span>

        <button @click="deleteNote(note.id)" class="delete-note-btn">×</button>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

const notes = ref([]);
const newNoteContent = ref('');
const editingNoteId = ref(null);

const API_URL = 'http://localhost:8080/notes';

onMounted(async () => {
  await fetchNotes();
});

// Функция для получения заметок
async function fetchNotes() {
  try {
    const response = await axios.get(API_URL);
    notes.value = response.data;
  } catch (err) {
    console.error('Ошибка получения заметок:', err);
  }
}

// Добавление заметки
async function addNote() {
  if (newNoteContent.value.trim() === '') return;

  try {
    await axios.post(API_URL, {
      content: newNoteContent.value,
      
    });
    // Обновляем список после добавления
    await fetchNotes();
    newNoteContent.value = '';
  } catch (err) {
    console.error('Ошибка при добавлении заметки:', err);
  }
}

// Удаление заметки
async function deleteNote(id) {
  try {
    await axios.delete(`${API_URL}/${id}`);
    // Обновляем список после удаления
    await fetchNotes();
  } catch (err) {
    console.error('Ошибка при удалении заметки:', err);
  }
}

// Начало редактирования
function startEditing(note) {
  editingNoteId.value = note.id;
}

// Завершение редактирования
async function finishEditing(note) {
  try {
    await axios.put(`${API_URL}/${note.id}`, {
      content: note.content,
    });
    editingNoteId.value = null;
    // Обновляем список после редактирования
    await fetchNotes();
  } catch (err) {
    console.error('Ошибка при обновлении заметки:', err);
  }
}
</script>
<style scoped>
.notes-card {
  background-color: var(--accent-color);
  padding: 24px;
  border-radius: 16px;
  display: flex;
  flex-direction: column;
}

.notes-header {
  margin-bottom: 16px;
}

/* Форма добавления */
.add-note-form {
  display: flex;
  gap: 8px;
  margin-bottom: 16px;
}
.add-note-form input {
  flex-grow: 1;
  border: 1px solid #ddd;
  border-radius: 6px;
  padding: 8px 12px;
  font-family: 'Poppins', sans-serif;
}
.add-note-form button {
  border: none;
  background-color: var(--primary-color);
  color: white;
  font-weight: bold;
  border-radius: 6px;
  width: 36px;
  height: 36px;
  cursor: pointer;
}

/* Список заметок */
.notes-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.note-item {
  background-color: white;
  padding: 12px;
  border-radius: 8px;
  font-size: 14px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 8px;
}
.note-item span {
  flex-grow: 1;
  cursor: pointer;
}

.edit-input {
  flex-grow: 1;
  border: none;
  padding: 0;
  font-size: 14px;
  font-family: 'Poppins', sans-serif;
  outline: none;
}

.delete-note-btn {
  border: none;
  background: #eee;
  color: #888;
  border-radius: 50%;
  cursor: pointer;
  width: 22px;
  height: 22px;
  font-size: 14px;
  line-height: 22px;
  text-align: center;
  flex-shrink: 0;
}
</style>