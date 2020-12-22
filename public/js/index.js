var app = new Vue({
    el: '#app',
    data: {
        currentDate: (new Date()),
        timer: ''
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
            // 2006-01-02T15:04:05Z
            return dateFns.format(this.currentDate, 'YYYY-MM-DDTHH:mm:ss[Z]')
        }
    },
    methods: {
        updateDate: function() {
            this.currentDate = new Date()
        },
        cancelAutoUpdate: function() { clearInterval(this.timer) }
    },
    beforeDestroy: function() {
        this.cancelAutoUpdate()
    }
})