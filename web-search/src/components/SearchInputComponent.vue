<template>
  <div class="search-bar">
    <input type="text" placeholder="Enter a word to search in emails" class="input" @keyup.enter="$emit('searchByTerm', $event.target.value)" />
    <i class="search icon"></i>
  </div>
</template>

<script>
export default {
  name: "SearchInputComponent",
};
</script>

<style scoped src="@/assets/css/search-bar.css">

</style>
