function getIndex (id, list) {
    for (let i = 0; i < list.length; i++) {
        if (id === list[i].id) {
            return i
        }
    }
    return -1
}

const app = Vue.createApp({
    data() {
        return {
            messages: []
        }
    },
    template: '<messages-list :messages="messages"/>',
})

app.component('messages-list', {
        props: ['messages'],
        data: function (){
            return {
                message: null
            }
        },
        template:
        /*html*/
            `
            <div style="position: relative; width: 300px">
                <message-form :messages="messages" :message="message"/>
                <message-row v-for = "message in messages" 
                                    :key = "message.id" 
                                    :message = "message" 
                                    :editMessage = "editMessage"/>
            </div>
        `,
        methods: {
            async getMessages() {
                const {data} = await axios.get("/message/")
                data.forEach(message => this.messages.push(message))
            },

            editMessage (message) {
                this.message = message

            }
        },
        created: function () {
            this.getMessages();
        },
    }
)

app.component('message-row', {
    props: ['message', 'editMessage'],
    template:
    /*html*/
        `<div> 
            <i>{{message.id}}</i>. {{message.text}}
            <span style="position: absolute; right: 0">
                <input type="button" value = "Edit" @click="edit"/>
                <!--input type="button" value="Delete" @click="delete"/-->
            </span>
        </div>`,
    methods: {
        edit: function () {this.editMessage(this.message)},
        delete: function () {}
    }

})

app.component('message-form', {
    props: ['messages', 'message'],
    watch: {
        message:function (n, o) {
            this.text = n.text
            this.id = n.id
        }
    },
    data: function () {
        return {
            text: '',
            id : ''
        }
    },
    template: /*html*/
        `
        <div>
            <input type="text" placeholder="enter your message" v-model = "text"/>
            <input type="button" value="publish" @click = "save"/>
        </div>
    `,
    methods: {
        save: function () {

            if (this.id) {
                var message = {id: this.id, text: this.text};
                axios.put("/message/"+ this.id, message).then(response => {
                    updatedMessage = response.data
                    let index = getIndex(updatedMessage.id, this.messages);
                    this.messages.splice(index, 1, updatedMessage)
                })
            }
            else {
                let message = {text: this.text};
                axios.post("/message/", message).then(response => {
                    this.messages.push(response.data)

                })
            }
            this.id = ''
            this.text = ''


        }
    }
})

app.mount('#app')