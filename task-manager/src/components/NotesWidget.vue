<template>
  <div class="widget-card notes-card">
    <h4>Notes</h4>
    <div class="note-item">
      <span>u can do this !</span>
      <button class="delete-note-btn">x</button>
    </div>
  </div>
</template>
<style scoped>
.widget-card { background-color: var(--accent-color); padding: 24px; border-radius: 16px; }
.note-item {
  background-color: white;
  padding: 12px;
  border-radius: 8px;
  font-size: 14px; display: flex;
  justify-content: space-between;
  align-items: center; margin-top: 16px; }
.delete-note-btn { border: none; background: #eee; color: #888; border-radius: 4px; cursor: pointer; width: 20px; height: 20px; }
</style>
