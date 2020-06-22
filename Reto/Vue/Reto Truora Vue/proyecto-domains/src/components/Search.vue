<template>
  <div>
      <b-card>
          <template v-slot:header>
              <b-form inline class="justify-content-center">
                <b-input  class="col-sm-6 mr-3" placeholder="Escribe el nombre del dominio" v-model="domainToSearch"></b-input>
                <b-button variant="primary" @click="searchDomain">Buscar</b-button>
              </b-form>
          </template>
          <b-card-body>
              <b-list-group>
                  <b-list-group-item>{{domain.host}}</b-list-group-item>
                  <b-list-group-item>{{domain.servers_change}}</b-list-group-item>
                  <b-list-group-item>{{domain.ssl_grade}}</b-list-group-item>
                  <b-list-group-item>{{domain.previous_ssl_grade}}</b-list-group-item>
                  <b-list-group-item>{{domain.logo}}</b-list-group-item>
                  <b-list-group-item>{{domain.title}}</b-list-group-item>
                  <b-list-group-item>{{domain.is_down}}</b-list-group-item>
              </b-list-group>
              <b-table striped hover :items="items"></b-table>
          </b-card-body>
      </b-card>
  </div>
</template>

<script>
import hostInfo from '../services/HostInfo'
export default {
    name: 'Search',
    data() {
        return {
            domainToSearch: '',
            domain: {
                host: "",
                servers_change: false,
                ssl_grade: "",
                previous_ssl_grade: "",
                logo: "",
                title: "",
                is_down: false,
                endpoints: [
                    //{ipAddress: String, grade: String, country: String, owner: String}
                ]
            },
            items: []  
        }
    },
    methods: {
        async searchDomain() {
            var response = await hostInfo.getInfo(this.domainToSearch)
            console.log(response)
            this.domain = response.data
            this.items = this.domain.endpoints
        }
    }
}
</script>
