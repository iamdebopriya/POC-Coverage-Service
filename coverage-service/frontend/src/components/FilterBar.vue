<script setup lang="ts">
defineProps<{ services: string[]; selectedService: string; from: string; to: string }>()
const emit = defineEmits<{
  (e: 'update:selectedService', v: string): void
  (e: 'update:from', v: string): void
  (e: 'update:to', v: string): void
  (e: 'filter'): void
  (e: 'download'): void
}>()
</script>
<template>
  <div class="bar">

    <div class="field">
      <label>
        Service
      </label>

      <select
        :value="selectedService"
        @change="emit('update:selectedService', ($event.target as HTMLSelectElement).value)"
      >
        <option value="">
          All Services
        </option>

        <option
          v-for="s in services"
          :key="s"
          :value="s"
        >
          {{ s }}
        </option>
      </select>
    </div>

    <div class="field">
      <label>
        From
      </label>

      <input
        type="datetime-local"
        :value="from"
        @input="emit('update:from', ($event.target as HTMLInputElement).value)"
      />
    </div>

    <div class="field">
      <label>
        To
      </label>

      <input
        type="datetime-local"
        :value="to"
        @input="emit('update:to', ($event.target as HTMLInputElement).value)"
      />
    </div>

    <div class="actions">

      <button
        class="btn-f"
        @click="emit('filter')"
      >
        Apply Filter
      </button>

      <button
        class="btn-d"
        @click="emit('download')"
      >
        Export CSV
      </button>

    </div>

  </div>
</template>
<style scoped>
.bar {
  display: flex;
  flex-wrap: wrap;
  align-items: end;
  gap: 0.9rem;

  padding: 1rem;

  border-radius: 18px;

  background: #fafafa;
  border: 1px solid #e5e5e5;
}

/* FIELD */

.field {
  display: flex;
  flex-direction: column;
  gap: 0.38rem;
  min-width: 180px;
  flex: 1;
}

label {
  font-size: 0.68rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.08em;
  color: #737373;
}

/* INPUTS */

input,
select {
  height: 42px;

  padding: 0 0.9rem;

  border-radius: 12px;

  border: 1px solid #d6d6d6;

  background: white;

  color: #18181b;

  font-size: 0.88rem;

  outline: none;

  transition:
    border-color 0.18s ease,
    background 0.18s ease;
}

input:focus,
select:focus {
  border-color: #a3a3a3;
  background: #fcfcfc;
}

/* BUTTONS */

.actions {
  display: flex;
  gap: 0.65rem;
}

/* FILTER */

.btn-f,
.btn-d {
  height: 42px;

  padding: 0 1.1rem;

  border-radius: 12px;

  font-size: 0.84rem;
  font-weight: 600;

  cursor: pointer;

  transition: 0.18s ease;
}

/* PRIMARY */

.btn-f {
  background: #111111;
  color: white;
  border: 1px solid #111111;
}

.btn-f:hover {
  opacity: 0.92;
}

/* SECONDARY */

.btn-d {
  background: #f3f3f3;
  color: #3f3f46;
  border: 1px solid #dcdcdc;
}

.btn-d:hover {
  background: #ebebeb;
}

/* MOBILE */

@media (max-width: 768px) {

  .bar {
    padding: 0.9rem;
  }

  .field {
    min-width: 100%;
  }

  .actions {
    width: 100%;
  }

  .btn-f,
  .btn-d {
    flex: 1;
  }
}
</style>
