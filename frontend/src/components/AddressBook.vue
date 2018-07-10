<template>
  <div class="addressbook">
    <b-jumbotron bg-variant="info" text-variant="white" border-variant="dark">
      <template slot="header">
        {{ title}}
      </template>
    </b-jumbotron>

    <h3 v-if="error">Ошибка: {{error}}</h3>

    <table class="table">
      <thead>
      <tr>
        <th>Имя</th>
        <th>Фамилия</th>
        <th>Телефон</th>
        <th>E-mail</th>
        <th>Github</th>
        <th></th>
        <th></th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="contact in contact_list" v-bind:key="contact.id">
        <td>{{contact.name}}</td>
        <td>{{contact.surname}}</td>
        <td>{{contact.phone}}</td>
        <td>{{contact.email}}</td>
        <td>{{contact.github}}</td>
        <td><b-button variant="warning" v-on:click="edit_contact(contact)">Редактировать контакт</b-button></td>
        <td><b-button variant="danger" v-on:click="remove_contact(contact)">Удалить контакт</b-button></td>
      </tr>
      </tbody>
    </table>

    <b-card>
    <b-form>
      <h3 v-if="edit_index == -1">Новый контакт</h3>
      <b-form-group label="Имя:">
        <b-form-input v-model="new_contact.name" required>
        </b-form-input>
      </b-form-group>
      <b-form-group label="Фамилия:">
        <b-form-input v-model="new_contact.surname" required>
        </b-form-input>
      </b-form-group>
      <b-form-group label="Телефон:">
        <b-form-input v-model="new_contact.phone" required>
        </b-form-input>
      </b-form-group>
      <b-form-group label="Email:">
        <b-form-input v-model="new_contact.email" type="email" required>
        </b-form-input>
      </b-form-group>
      <b-form-group label="Github:">
        <b-form-input v-model="new_contact.github" required>
        </b-form-input>
      </b-form-group>
      <b-button variant="success" v-if="edit_index == -1" v-on:click="add_new_contact">Добавить контакт</b-button>
      <b-button variant="success" v-if="edit_index > -1" v-on:click="end_of_edition">Закончить редактирование</b-button>
    </b-form>
    </b-card>
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
          'id': 0,
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
      const url = '/api/addressbook/contacts/' + item.id
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
      const url = '/api/addressbook/contacts/' + this.new_contact.id
      axios.put(url, this.new_contact).then(response => {
        this.edit_index = -1
        this.new_contact = {
          'id': 0,
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
