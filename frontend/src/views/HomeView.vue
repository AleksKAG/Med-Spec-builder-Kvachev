<template>
  <div class="home">
    <CabinetSelector
      v-if="cabinets.length"
      :cabinets="cabinets"
      v-model:selected="selectedCabinets"
    />
    <button @click="generateSpec" :disabled="selectedCabinets.length === 0">
      Сформировать спецификацию
    </button>
    <button @click="exportExcel" :disabled="selectedCabinets.length === 0">
      Скачать Excel
    </button>

    <div v-if="spec.length" class="spec-result">
      <h3>Спецификация оборудования</h3>
      <table>
        <thead>
          <tr>
            <th>Наименование</th>
            <th>Производитель</th>
            <th>Модель</th>
            <th>Артикул</th>
            <th>Кол-во</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, idx) in spec" :key="idx">
            <td>{{ item.item_name }}</td>
            <td>{{ item.manufacturer }}</td>
            <td>{{ item.model }}</td>
            <td>{{ item.article }}</td>
            <td>{{ item.total_qty }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import api from '@/api/api'
import CabinetSelector from '@/components/CabinetSelector.vue'

export default {
  components: { CabinetSelector },
  data() {
    return {
      cabinets: [],
      selectedCabinets: [],
      spec: []
    }
  },
  async mounted() {
    const res = await api.get('/cabinets')
    this.cabinets = res.data
  },
  methods: {
    async generateSpec() {
      const res = await api.post('/spec', { cabinet_ids: this.selectedCabinets })
      this.spec = res.data
    },
    async exportExcel() {
      const blob = await api.post('/export/xlsx', { cabinet_ids: this.selectedCabinets }, {
        responseType: 'blob'
      }).then(r => r.data)
      const url = window.URL.createObjectURL(blob)
      const a = document.createElement('a')
      a.href = url
      a.download = 'specification.xlsx'
      a.click()
    }
  }
}
</script>

<style scoped>
button {
  margin: 10px;
  padding: 8px 16px;
}
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}
th, td {
  border: 1px solid #ccc;
  padding: 8px;
  text-align: left;
}
</style>
