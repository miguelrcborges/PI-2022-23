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

		const patientNumber = document.createElement('td');
		patientNumber.textContent = data[ip].UserName;
		row.appendChild(patientNumber);

		const patientName = document.createElement('td');
		patientName.textContent = data[ip].UserNumber;
		row.appendChild(patientName);

		const order = document.createElement('td');
		order.textContent = data[ip].Order;
		row.appendChild(order);

		const options = document.createElement('td');
		if (data[ip].PatientNumber === 0) {
			const assign = document.createElement('button');
			assign.textContent = "Assign to a user";
			options.appendChild(assign);
		}
		
		row.appendChild(options);

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
