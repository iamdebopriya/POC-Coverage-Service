<script setup lang="ts">
defineProps<{
  label: string
  value: string | number
  sub?: string
}>()
</script>

<template>
  <div class="stat-card">

    <div class="lbl">
      {{ label }}
    </div>

    <div class="val">
      {{ value }}
    </div>

    <div
      v-if="sub"
      class="sub"
    >
      {{ sub }}
    </div>

  </div>
</template>
<style scoped>
.stat-card {
  background: #fafafa;

  border: 1px solid #e5e5e5;

  border-radius: 18px;

  padding: 1rem 1.1rem;

  display: flex;
  flex-direction: column;
  gap: 0.45rem;

  transition:
    transform 0.18s ease,
    background 0.18s ease,
    border-color 0.18s ease;
}

.stat-card:hover {
  background: #f7f7f7;
  border-color: #d4d4d4;
  transform: translateY(-1px);
}

/* LABEL */

.lbl {
  font-size: 0.68rem;
  font-weight: 700;

  text-transform: uppercase;
  letter-spacing: 0.08em;

  color: #737373;
}

/* VALUE */

.val {
  font-size: 2rem;
  font-weight: 700;

  line-height: 1;

  color: #111111;

  letter-spacing: -0.04em;
}

/* SUBTEXT */

.sub {
  font-size: 0.78rem;
  color: #8a8a8a;
  margin-top: 0.1rem;
}

/* MOBILE */

@media (max-width: 768px) {

  .stat-card {
    padding: 0.9rem 1rem;
  }

  .val {
    font-size: 1.7rem;
  }
}
</style>
