"use strict";

const connectedElemented = document.querySelector("#connected");
const monitorButton = document.querySelector("#monitor-button");
const connectedStream = new EventSource("/api/get/quantity");
let dataStream;

connectedStream.addEventListener('updateDevicesCount', (event) => {
  connectedElemented.textContent = event.data + " devices connected currently.";
});

const startDataStream = () => {
  monitorButton.textContent = "Stop monitoring"
  dataStream = new EventSource("/api/get/data");
  dataStream.addEventListener('updateDevicesData', (e) => console.log(e.data));
  monitorButton.addEventListener('click', closeDataStream, {once: true});
}

const closeDataStream = () => {
  monitorButton.textContent = "Start monitoring"
  dataStream.close();
  monitorButton.addEventListener('click', startDataStream, {once: true});
}

monitorButton.addEventListener('click', startDataStream, {once: true});