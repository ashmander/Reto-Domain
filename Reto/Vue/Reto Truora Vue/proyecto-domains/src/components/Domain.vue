<template>
  <div>
      <div class="card">
          <div class="card-header">
              <h2>Rutas Consultadas</h2>
          </div>
          <div class="card-body">
              <b-table v-if="items != undefined" hover striped :items="items" :fields="fields"> </b-table>
          </div>
      </div>
  </div>
</template>

<script>
import allEndpoints from '../services/AllEndpoints'
export default {
    name: 'Domain',
    data() {
        return {
            fields: ['host'],
            items: []
        }
    },
    methods: {
        async getAllEndpoints() {
            var response = await allEndpoints.getEndpoints()
            this.items = response.data
            console.log(items)
        }
    },
    created() {
        this.getAllEndpoints()
    }
}
</script>

<style>

</style>