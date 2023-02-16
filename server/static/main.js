"use strict";

const connectedElemented = document.querySelector("#connected");
const connectedStream = new EventSource("/api/get/quantity");

connectedStream.addEventListener('updateDevicesCount', (event) => {
  connectedElemented.textContent = event.data + " devices connected.";
});