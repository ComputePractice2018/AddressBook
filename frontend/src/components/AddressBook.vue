<template>
  <div class="addressbook">
    <h1>{{ title }}</h1>

    <h3 v-if="error">Ошибка: {{error}}</h3>

    <table>
      <tr>
        <th>Имя</th>
        <th>Фамилия</th>
        <th>Телефон</th>
        <th>E-mail</th>
        <th>Github</th>
      </tr>
      <tr v-for="contact in contact_list" v-bind:key="contact.phone">
        <td>{{contact.name}}</td>
        <td>{{contact.surname}}</td>
        <td>{{contact.phone}}</td>
        <td>{{contact.email}}</td>
        <td>{{contact.github}}</td>
        <td><button v-on:click="edit_contact(contact)">Редактировать контакт</button></td>
        <td><button v-on:click="remove_contact(contact)">Удалить контакт</button></td>
      </tr>
    </table>

    <h3 v-if="edit_index == -1">Новый контакт</h3>
    <form>
      <p>Имя: <input type="text" v-model="new_contact.name"></p>
      <p>Фамилия: <input type="text" v-model="new_contact.surname"></p>
      <p>Телефон: <input type="text" v-model="new_contact.phone"></p>
      <p>Email: <input type="text" v-model="new_contact.email"></p>
      <p>Gihub: <input type="text" v-model="new_contact.github"></p>
      <button v-if="edit_index == -1" v-on:click="add_new_contact">Добавить контакт</button>
      <button v-if="edit_index > -1" v-on:click="end_of_edition">Закончить редактирование</button>
    </form>
  </div>
</template>

<script>
const axios = require('axios')

export default {
  name: 'AddressBook',
  props: {
    title: String
  },
  data: function () {
    return {
      edit_index: -1,
      error: '',
      contact_list: [],
      new_contact:
        {
          'name': '',
          'surname': '',
          'phone': '',
          'email': '',
          'github': ''
        }
    }
  },
  mounted: function () {
    this.get_contacts()
  },
  methods: {
    get_contacts: function () {
      this.error = ''
      const url = '/api/addressbook/contacts'
      axios.get(url).then(response => {
        this.contact_list = response.data
      }).catch(response => {
        this.error = response.response.data
      })
    },
    add_new_contact: function () {
      this.error = ''
      const url = '/api/addressbook/contacts'
      axios.post(url, this.new_contact).then(response => {
        this.contact_list.push(this.new_contact)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    remove_contact: function (item) {
      this.error = ''
      const url = '/api/addressbook/contacts/' + this.contact_list.indexOf(item)
      axios.delete(url).then(response => {
        this.contact_list.splice(this.contact_list.indexOf(item), 1)
      }).catch(response => {
        this.error = response.response.data
      })
    },
    edit_contact: function (item) {
      this.edit_index = this.contact_list.indexOf(item)
      this.new_contact = this.contact_list[this.edit_index]
    },
    end_of_edition: function () {
      this.error = ''
      const url = '/api/addressbook/contacts/' + this.edit_index
      axios.put(url, this.new_contact).then(response => {
        this.edit_index = -1
        this.new_contact = {
          'name': '',
          'surname': '',
          'phone': '',
          'email': '',
          'github': ''
        }
      }).catch(response => {
        this.error = response.response.data
      })
    }
  }
}
</script>

<style scoped>

</style>
