<template>
  <div class="email" @click="$emit('showBody', email._source)" :class="{ 'email-active': emailActive === email._source }">
    <div class="row m-0">
      <div class="col-2" style="align-self: center;">
        <div class="icon">
          <label>{{ email._source.From[0].toUpperCase() }}</label>
        </div>
      </div>
      <div class="col " style="align-self: center;">
        <h2>{{ email._source.From }}</h2>
        <p>{{ email._source.Subject }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    email: Object,
    emailActive: Object
  },
};
</script>

<style scoped src="@/assets/css/email.css">

</style>
