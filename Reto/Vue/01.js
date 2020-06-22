const app = new Vue({
    el: '#app',
    data: {
        titles: 'Hola mundo con Vue',
        fruts: [
            {name: 'Manzana', state: false},
            {name: 'Pera', state: true},
            {name: 'Mango', state: false}
        ],
        newFrut: ''
    },
    methods: {
        addFrut() {
            this.fruts.push({
                name: this.newFrut
            })
        }
    }
})