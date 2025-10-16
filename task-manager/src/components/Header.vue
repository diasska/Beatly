<template>
  <header class="header">
    <h1>Hi, {{ name }}</h1>
    <p>Let's finish your task today!</p>
  </header>
</template>
<script setup>
defineProps({ name: { type: String, default: 'User' } });
</script>
<style scoped>
h1 { font-size: 28px; font-weight: 700; }
p { color: var(--subtle-text-color); margin-top: 4px; }
</style>
