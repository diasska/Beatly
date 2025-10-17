<template>
  <nav class="sidebar">
    <div class="logo-container">
      <img src="@/assets/logo.png" alt="Our name logo" class="logo-img" />
      <h2>Our name</h2>
    </div>
    <ul class="nav-links">
      <li class="nav-link">
        <RouterLink to="/">Dashboard</RouterLink>
      </li>
      <li class="nav-link">
        <RouterLink to="/my-tasks">My tasks</RouterLink>
      </li>
    </ul>
  </nav>
</template>

<style scoped>
.sidebar {
  background-color: var(--card-background-color);
  padding: 24px;
  border-radius: 16px;
  box-shadow: 0 8px 24px var(--shadow-color);
  display: flex;
  flex-direction: column;
}

.logo-container {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 40px;
}

.logo-img {
  width: 36px;
  height: 36px;
}

.nav-links {
  list-style: none;
  padding: 0;
  margin: 0;
}

.nav-link {
  padding: 12px 16px;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
}

.nav-link.active {
  background-color: var(--accent-color);
  color: var(--primary-color);
}

.nav-link a {
  text-decoration: none;
  color: inherit;
}
.nav-link.active a {
  font-weight: bold;
}
</style>