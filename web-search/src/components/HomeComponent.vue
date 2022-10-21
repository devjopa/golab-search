<template>
  <Header />
  <MDBRow class="m-0">
    <MDBCol md="5" style="background-color: #0000000a;">
      <SearchInputComponent @searchByTerm="executeSearchByTerm" />
      <FinderComponent ref="childComponentRef" @showBody="showBody" />
    </MDBCol>
    <MDBCol md="7">
      <BodyComponent :email="email" />
    </MDBCol>
  </MDBRow>

</template>

<script>

import Header from "./HeaderComponent.vue";
import FinderComponent from "./FinderComponent.vue";
import SearchInputComponent from "./SearchInputComponent.vue"
import BodyComponent from "./BodyComponent.vue"
import { MDBRow, MDBCol } from 'mdb-vue-ui-kit';


export default {
  name: "HomeComponent",
  components: {
    Header,
    FinderComponent,
    SearchInputComponent,
    MDBCol,
    MDBRow,
    BodyComponent
  },
  data: () => {
    return {
      email: ""
    };
  },
  methods: {
    async executeSearchByTerm(term) {
      this.$refs.childComponentRef.executeSearchByTerm(term);
      this.email = ""
    },
    showBody(data) {
      this.email = data;
    }
  },
};
</script>
