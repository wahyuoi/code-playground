import Ember from 'ember';

export default Ember.Route.extend({
  beforeModel() {
    this.replaceWith('rentals');
  },

  model() {
    return this.get('store').findAll('rental');
  }
});

