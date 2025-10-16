<template>
  <div class="dashboard-layout">
    <Sidebar />
    <main class="main-content">
      <Header name="User" />
      <TodayTaskCard />
      <AssignmentsWidget />
    </main>
    <aside class="right-sidebar">
      <ProfileWidget userName="User" />
      <NotesWidget />
    </aside>
  </div>
</template>

<script setup>
import Sidebar from '@/components/Sidebar.vue';
import Header from '@/components/Header.vue';
import TodayTaskCard from '@/components/TodayTaskCard.vue';
import AssignmentsWidget from '@/components/AssignmentsWidget.vue';
import ProfileWidget from '@/components/ProfileWidget.vue';
import NotesWidget from '@/components/NotesWidget.vue';
</script>

<style scoped>
.dashboard-layout {
  display: grid;
  grid-template-columns: 240px 1fr 350px;
  gap: 24px;
  padding: 24px;
  height: 100vh;
}

.main-content, .right-sidebar {
  display: flex;
  flex-direction: column;
  gap: 24px;
}
</style>
