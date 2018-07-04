<template>
  <div class="home">
    <AddressBook title="Адресная книга"/>
  </div>
</template>

<script>
// @ is an alias to /src
import AddressBook from '@/components/AddressBook.vue'

export default {
  name: 'home',
  components: {
    AddressBook
  }
}
</script>
