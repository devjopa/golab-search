<template>
  <div class="wrapper" v-if="email.Body !== ''">
    <div class="row m-0 email" >
      <div class="col-1" style="align-self: center;">
        <div class="icon">
          <label>{{ email.From[0] }}</label>
        </div>
      </div>
      <div class="col" style="align-self: center;">
        <p class="from"> From: {{ email.From }}</p>
        <p class="from"> To: {{ email.To }}</p>
        <p class="to"> Subject: {{ email.Subject }}</p>
      </div>

    </div>

    <div class="body">
      <p> {{ email.Body }}</p>
    </div>
  </div>
  <div class="wrapper" v-if="email.Body === ''">
    <div style="margin-top:15rem">
      <img src="../assets/instructions.png" class="img-size"/>
    </div>
    <div class="label-information">
      <label>Please select email row to show body information...</label>
    </div>
  </div>
</template>
<script>
export default {
  name: "BodyComponent",
  props: {
    email: Object,
  },
};
</script>

<style scoped src="@/assets/css/body.css">

</style>