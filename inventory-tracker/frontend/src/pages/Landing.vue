<template>
  <div class="landing">
    <div class="landing-content">
      <h1 class="landing-title">
        Track Your Stock<br />
        <span class="highlight">Smarter & Faster</span>
      </h1>

      <p class="landing-desc">
        A simple and powerful inventory management system to track your stock,
        monitor low inventory, and manage items effortlessly.
      </p>

      <div class="landing-actions">
        <router-link to="/inventory" class="btn-get-started">
          Get Started →
        </router-link>
        <router-link to="/dashboard" class="btn-dashboard">
          View Dashboard
        </router-link>
      </div>

      <div class="features-grid">
        <div class="feature-card">
          <div class="feature-icon blue">
            <i class="fas fa-boxes-stacked"></i>
          </div>
          <div class="feature-title">Track Inventory</div>
          <div class="feature-desc">Add, edit and delete items with ease</div>
        </div>
        <div class="feature-card">
          <div class="feature-icon red">
            <i class="fas fa-triangle-exclamation"></i>
          </div>
          <div class="feature-title">Low Stock Alerts</div>
          <div class="feature-desc">Get notified when stock runs low</div>
        </div>
        <div class="feature-card">
          <div class="feature-icon green">
            <i class="fas fa-chart-line"></i>
          </div>
          <div class="feature-title">Coverage Dashboard</div>
          <div class="feature-desc">Monitor test coverage and metrics</div>
        </div>
        <div class="feature-card">
          <div class="feature-icon amber">
            <i class="fas fa-file-csv"></i>
          </div>
          <div class="feature-title">Export Reports</div>
          <div class="feature-desc">Download coverage results as CSV</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
</script>

<style scoped>
.landing {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 3rem 2rem;
  overflow: hidden; 
}

.landing-content {
  max-width: 750px;
  text-align: center;
}

.badge-tag {
  display: inline-block;
  background: #958866;
  color: #4b2304;
  font-size: 12px;
  font-weight: 500;
  padding: 5px 14px;
  border-radius: 99px;
  margin-bottom: 1.5rem;
  border: 1px solid #bee3f8;
}

.landing-title {
  font-size: 3rem;
  font-weight: 700;
  color: #1862ac;
  line-height: 1.2;
  margin-bottom: 1.25rem;
}

.highlight {
  color: #050c11;
}

.landing-desc {
  font-size: 1.05rem;
  color: #481a1a;
  line-height: 1.7;
  margin-bottom: 2rem;
  max-width: 560px;
  margin-left: auto;
  margin-right: auto;
}

.landing-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  margin-bottom: 3rem;
}

.btn-get-started {
  background-color: #2c3e50;
  color: white;
  padding: 0.75rem 2rem;
  border-radius: 8px;
  text-decoration: none;
  font-size: 0.95rem;
  font-weight: 500;
  transition: background 0.2s;
}

.btn-get-started:hover {
  background-color: #1a252f;
}

.btn-dashboard {
  background-color: white;
  color: #2c3e50;
  padding: 0.75rem 2rem;
  border-radius: 8px;
  text-decoration: none;
  font-size: 0.95rem;
  font-weight: 500;
  border: 1px solid #ddd;
  transition: all 0.2s;
}

.btn-dashboard:hover {
  background-color: #f5f6fa;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
  text-align: left;
}

.feature-card {
  background: white;
  border-radius: 10px;
  padding: 1.25rem;
  box-shadow: 0 1px 4px rgba(0,0,0,0.07);
}

.feature-icon {
  width: 40px;
  height: 40px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  margin-bottom: 0.75rem;
}

.feature-icon.blue { background: #e8f4fd; color: #2980b9; }
.feature-icon.red { background: #fdf0f0; color: #e74c3c; }
.feature-icon.green { background: #e8f8f0; color: #27ae60; }
.feature-icon.amber { background: #fef9e7; color: #f39c12; }

.feature-title {
  font-size: 0.9rem;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 4px;
}

.feature-desc {
  font-size: 0.8rem;
  color: #888;
  line-height: 1.5;
}
</style>