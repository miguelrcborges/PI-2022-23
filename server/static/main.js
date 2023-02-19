"use strict";

const connectedElemented = document.querySelector("#connected");
const monitorButton = document.querySelector("#monitor-button");
const dataElement = document.querySelector("#data");
const connectedStream = new EventSource("/api/get/quantity");
let dataStream;
let debug

connectedStream.addEventListener('updateDevicesCount', (event) => {
  connectedElemented.textContent = event.data + " devices connected currently.";
});

const processDataStream = (data) => {
  let elements = [];
  for (const ip in data) {
    const row = document.createElement('tr');

    const ipElement = document.createElement('td');
    ipElement.textContent = ip;
    row.appendChild(ipElement);

    const locationElement = document.createElement('td');
    const locationX = document.createElement('p');
    const locationY = document.createElement('p');
    locationX.textContent = "x: " + data[ip].Position.X;
    locationY.textContent = "y: " + data[ip].Position.Y;
    locationElement.appendChild(locationX);
    locationElement.appendChild(locationY);
    row.appendChild(locationElement);

    const angleElement = document.createElement('td');
    angleElement.textContent = data[ip].Angle + 'º';
    row.appendChild(angleElement);

    const destinyElement = document.createElement('td');
    const destinyX = document.createElement('p');
    const destinyY = document.createElement('p');
    destinyX.textContent = "x: " + data[ip].Target.X;
    destinyY.textContent = "y: " + data[ip].Target.Y;
    destinyElement.appendChild(destinyX);
    destinyElement.appendChild(destinyY);
    row.appendChild(destinyElement);

    elements.push(row);
  }
  dataElement.replaceChildren(...elements);
}

const startDataStream = () => {
  monitorButton.textContent = "Stop monitoring"
  dataStream = new EventSource("/api/get/data");
  dataStream.addEventListener('updateDevicesData', (e) => processDataStream(JSON.parse(e.data)));
  monitorButton.addEventListener('click', closeDataStream, {once: true});
}

const closeDataStream = () => {
  monitorButton.textContent = "Start monitoring"
  dataStream.close();
  monitorButton.addEventListener('click', startDataStream, {once: true});
}

monitorButton.addEventListener('click', startDataStream, {once: true});