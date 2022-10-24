<template>
  <div class="d-flex align-items-start mb-3" >
    <Suspense>
      <div class="col">
        <ListComponent :searchTerm="searchTerm" :emails="emails" @loadMoreData="loadMoreData" @showBody="showBody" />
      </div>
    </Suspense>
  </div>
</template>

<script>
import ListComponent from "./ListComponent.vue";
import { searchEmailsByTerm } from "../../services/index";

export default {
  name: "FinderComponent",
  data: () => {
    return {
      searchTerm: "forecast",
      from: 0,
      maxResults: 10,
      emails: [],
      totalRows: 100
    };
  },
  components: {
    ListComponent
  },
  methods: {
    async loadData() {
      const dataSearchApi = await searchEmailsByTerm({
        term: this.searchTerm,
        from: this.from,
        maxResults: this.maxResults,
      });
      this.totalRows = dataSearchApi.Hits.Total.Value;
      return dataSearchApi.Hits.Hits;
    },
    async executeSearchByTerm(term) {
      this.searchTerm = term;
      this.from = 0;
      this.emails = await this.loadData();

    },

    async loadMoreData() {
      this.from = this.from + this.maxResults + 1;

      if(this.from <= this.totalRows)
      {
        let newEmails = await this.loadData();
        this.emails.push(...newEmails);
      }
    },

    showBody(data) {
      this.$emit("showBody", data)
    }
  },
};
</script>
