var app = new Vue({
    el: '#app',
    data: {
        currentDate: (new Date()),
        timer: '',
        customTimestamp: '',
        customDate: '',
        newTimestamp: '',
        newDate: ''
    },
    created: function() {
        this.updateDate()
        this.timer = setInterval(this.updateDate, 1000)
    },
    computed: {
        ts: function() {
            return Math.floor(this.currentDate.getTime() / 1000)
        },
        date: function() {
            return dateFns.format(this.currentDate, 'YYYY-MM-DDTHH:mm:ss[Z]')
        }
    },
    methods: {
        updateDate: function() {
            this.currentDate = new Date()
        },
        getNewDate: function() {
            if (this.customTimestamp === "") {
                return
            }

            var vm = this
            fetch(`/api/v1/ts/${this.customTimestamp}`)
                .then(response => response.json())
                .then(data => vm.newDate = data.date)
                .catch(error => console.log("Error hitting API: " + error))
        },
        getNewTimestamp: function() {
            if (this.customDate === "") {
                return
            }

            var vm = this
            fetch(`/api/v1/date/${this.customDate}`)
                .then(response => response.json())
                .then(data => vm.newTimestamp = data.timestamp)
                .catch(error => console.log("Error hitting API: " + error))
        },
        cancelAutoUpdate: function() { clearInterval(this.timer) }
    },
    beforeDestroy: function() {
        this.cancelAutoUpdate()
    },
    watch: {
        customTimestamp: function(_newTimestamp) {
            this.getNewDate()
        },
        customDate: function(_newDate) {
            this.getNewTimestamp()
        }
    }
})