import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import AddressBook from '@/components/AddressBook.vue'

describe('AddressBook.vue', () => {
  it('renders props.title when passed', () => {
    const title = 'Название'
    const wrapper = shallowMount(AddressBook, {
      propsData: { title }
    })
    expect(wrapper.text()).to.include(title)
  })
  it('renders titles', () => {
    const wrapper = shallowMount(AddressBook, {})
    expect(wrapper.text()).to.include('Имя')
    expect(wrapper.text()).to.include('Фамилия')
    expect(wrapper.text()).to.include('Телефон')
    expect(wrapper.text()).to.include('Email')
    expect(wrapper.text()).to.include('Github')
  })
})
