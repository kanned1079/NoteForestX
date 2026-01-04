<script setup lang="ts">

</script>

<template>
  <div class="load-card">
    <div class="card-name-price">
      <div class="load-card-nm">UIverse Logistics, Inc.</div>
      <div class="load-card-price">$900</div>
    </div>
    <div class="card-time">
      <div class="load-card-tm">3h</div>
      <div class="load-card-mi">$3.60/mi</div>
    </div>
    <div class="card-info">
      <div class="card-date-wrapper">
        <div class="month">Jun</div>
        <div class="date">07</div>
        <div class="day">Wed</div>
      </div>
      <div class="card-city-unit-wrapper">
        <div class="city-with-mi">
          <span class="round-icon"></span>
          <div class="city-nm">Marlboro, NJ</div>
          <span class="mi">16 mi</span>
        </div>
        <div class="city-with-mi">
          <span class="square-icon"></span>
          <div class="city-nm">Mansfield, PA</div>
          <span class="mi">157 mi</span>
        </div>
      </div>
      <div class="rate-check-box">
        <div class="rc">RATE CHECK</div>
        <div class="rate">$950</div>
        <div class="rc">$3.80/mi</div>
      </div>
    </div>
    <div class="units-wrapper">
      <div class="pills-wrapper">
        <span class="unit-pill">TL</span>
        <span class="unit-pill">F</span>
        <span class="unit-pill">48 ft</span>
        <span class="unit-pill">46000 lbs</span>
      </div>
      <div class="mi">197mi</div>
    </div>
  </div>

</template>

<style scoped>
.load-card {
  --primary-color: #484848;
  --background: rgb(255, 255, 255);
  --hover-bg: rgb(255, 249, 241);
  --box-shadow: rgb(218, 226, 234);
  --border: rgb(255, 157, 31);
  --price-color: rgb(0, 181, 41);
  --unit-pills-bg: rgb(228, 247, 232);
  --unit-pills-color: rgb(72, 72, 72);
  --bold: 600;
  --semi-bold: 500;
  font-family: "poppins", sans-serif;
  display: grid;
  grid-row-gap: 4px;
  text-decoration: none;
  padding: 8px 16px;
  border-radius: 8px;
  box-shadow: var(--box-shadow) 0px 0px 10px 0px;
  background: var(--background);
  margin-bottom: 10px;
  color: var(--primary-color);
  transition: 0.4s ease-in-out;
  border-bottom: 6px solid var(--unit-pills-bg);
  cursor: pointer;
  margin: 8px;
}
.load-card:hover {
  box-shadow: var(--box-shadow) 0px 2px 4px 0px;
  background: var(--hover-bg);
  border-bottom: 6px solid var(--border);
  transition: 0.2s ease-in-out;
}
.card-name-price,
.card-time,
.card-info,
.units-wrapper {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.card-name-price {
  font-size: 20px;
}
.load-card-price {
  color: var(--price-color);
  font-weight: var(--bold);
}
.card-time {
  font-size: 12px;
  opacity: 0.5;
}
.card-info {
  column-gap: 12px;
  margin: 8px 0px 8px 19px;
  text-align: center;
}
.card-date-wrapper {
  display: flex;
  flex-direction: column;
}
.card-info .month,
.card-info .day {
  font-size: 12px;
  opacity: 0.5;
}
.card-info .date {
  font-size: 28px;
  font-weight: var(--semi-bold);
  transition: 0.4s ease-in-out;
}
.city-with-mi {
  display: flex;
  align-items: center;
}
.city-nm {
  font-size: 14px;
}
.mi {
  font-size: 12px;
  opacity: 0.5;
  flex-shrink: 0;
}
.rc {
  font-size: 9px;
  opacity: 0.5;
}
.card-city-unit-wrapper {
  display: flex;
  flex-direction: column;
  row-gap: 30px;
  position: relative;
  overflow: hidden;
}
.card-city-unit-wrapper::before {
  content: "";
  position: absolute;
  height: 71%;
  top: 15%;
  left: 4px;
  border-left: 2px dotted gray;
  opacity: 0.5;
  z-index: 0;
}
.city-with-mi {
  display: flex;
  align-items: center;
  column-gap: 8px;
}
.round-icon {
  display: block;
  height: 9px;
  width: 9px;
  border-radius: 10px;
  border: 2px solid gray;
  background-color: #ffffff;
  z-index: 1;
}
.square-icon {
  display: block;
  height: 9px;
  width: 9px;
  border: 2px solid var(--border);
  background-color: #ffffff;
  z-index: 1;
}
.rate-check-box {
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
  row-gap: 4px;
  border: 1px dashed rgb(145, 145, 145);
  border-radius: 2px;
  padding: 4px;
  transition: 0.2s ease-in-out;
}
.load-card:hover .rate-check-box {
  border: 1px dashed var(--background);
  color: var(--background);
  background-color: var(--border);
}
.load-card:hover .rate-check-box .rc,
.load-card:hover .rate-check-box .mi {
  opacity: 1;
}
.rate-check-box .rate {
  font-size: 12px;
  height: 15px;
  line-height: 16px;
  opacity: 1;
  transition: opacity 2s ease 0s;
}
.pills-wrapper {
  display: flex;
  align-items: center;
  column-gap: 4px;
  color: rgb(72, 72, 72);
  font-size: 12px;
}
.unit-pill {
  background-color: var(--unit-pills-bg);
  border-radius: 10px;
  padding: 4px 12px;
  transition: 0.4s ease-in-out;
}
.unit-pill:hover {
  background-color: var(--price-color);
  color: var(--background);
}

</style>