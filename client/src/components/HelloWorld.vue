<template>
  <v-container class="my-10">
    <v-sheet class="d-flex justify-center">
      <v-sheet width="100%" class="px-7">
        <v-card
          color="#a1f1ff"
          class="pa-3 mb-4 rounded-lg"
          v-for="item in datas"
          :key="item.title"
        >
          <h3 class="my-3 text-center">{{ item.title }}</h3>
          <GChart type="PieChart" :data="item.data" :options="chartOptions"
        /></v-card>
      </v-sheet>
    </v-sheet>
  </v-container>
</template>

<script>
import { GChart } from "vue-google-charts";
export default {
  name: "HelloWorld",

  components: {
    GChart,
  },

  data() {
    return {
      datas: [],
      // Array will be automatically processed with visualization.arrayToDataTable function
      chartData: {
        title: "たけのこきのこ戦争",
        data: [
          ["Task", "Hours per Day"],
          ["きのこ派", 11],
          ["たけのこ派", 10],
        ],
      },
      options: {},
      chartOptions: {
        is3D: true,
        backgroundColor: "#a1f1ff",
      },
    };
  },

  mounted() {
    // this.VueAxios.get("/api/user/me", {
    //   withCredentials: true,
    // });
    for (let a = 0; a < 10; a++) {
      this.datas.push(this.chartData);
    }
  },
};
</script>
