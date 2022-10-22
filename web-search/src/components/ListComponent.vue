<template>
  <div class="scrolling-component" id="infinite-list" style="height: calc(100vh - 12.1rem); overflow: auto;">
    <Email v-for="email in emails" :emailActive="emailActive" :email="email" :key="email._id" @showBody="showBody" />
  </div>
</template>

<script>
import Email from "./EmailComponent.vue";
import { ref, onMounted, onUnmounted } from "vue";

export default {
  name: "ListComponent",
  components: {
    Email,
  },
  props: {
    emails: Array,
  },
  data: () => {
    return {
      emailActive: {From: '', To:'', Body:'', Date:''}
    };
  },
  emits: ["loadMoreData", "showBody"],
  async setup(props, context) {
    const scrollComponent = ref(null);

    onMounted(() => {
      const listElm = document.querySelector('#infinite-list');
      listElm.addEventListener('scroll', () => {
        if (listElm.scrollTop + listElm.clientHeight >= listElm.scrollHeight) {
          context.emit("loadMoreData");
        }
      });
    });

    onUnmounted(() => {
      const listElm = document.querySelector('#infinite-list');
      listElm.removeEventListener("scroll", handleScroll);
    });

    const handleScroll = () => {
      let element = scrollComponent.value;
      if (element.getBoundingClientRect().bottom < window.innerHeight) {
        context.emit("loadMoreData");
      }
    };

    return {
      scrollComponent,
    };
  },
  methods: {
    showBody(data) {
      this.emailActive = data;
      this.$emit("showBody", data)
    }
  }
};
</script>
