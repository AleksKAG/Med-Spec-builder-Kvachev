<template>
  <div class="cabinet-selector">
    <h2>Выберите кабинеты</h2>
    <div v-for="dept in groupedCabinets" :key="dept.name" class="department">
      <h3>{{ dept.name }}</h3>
      <label v-for="cab in dept.cabinets" :key="cab.id" class="cabinet-item">
        <input
          type="checkbox"
          :value="cab.id"
          v-model="selectedIds"
        />
        {{ cab.number }} – {{ cab.name }}
      </label>
    </div>
  </div>
</template>

<script>
export default {
  props: ['cabinets'],
  emits: ['update:selected'],
  data() {
    return {
      selectedIds: []
    }
  },
  computed: {
    groupedCabinets() {
      const groups = {}
      this.cabinets.forEach(cab => {
        if (!groups[cab.department]) {
          groups[cab.department] = { name: cab.department, cabinets: [] }
        }
        groups[cab.department].cabinets.push(cab)
      })
      return Object.values(groups)
    }
  },
  watch: {
    selectedIds(newVal) {
      this.$emit('update:selected', newVal)
    }
  }
}
</script>

<style scoped>
.department {
  margin-bottom: 20px;
}
.cabinet-item {
  display: block;
  margin: 4px 0;
  cursor: pointer;
}
</style>
